
package http

import (
	"bytes"
	"compress/gzip"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"datacollector-edge/api"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/common"
	"datacollector-edge/stages/lib/datagenerator"
	"datacollector-edge/stages/lib/httpcommon"
	"datacollector-edge/stages/stagelibrary"
)

const (
	Library   = "streamsets-datacollector-basic-lib"
	StageName = "com_streamsets_pipeline_stage_destination_http_HttpClientDTarget"
)

type HttpClientDestination struct {
	*common.BaseStage
	*httpcommon.HttpCommon
	Conf HttpClientTargetConfig `ConfigDefBean:"conf"`
}

type HttpClientTargetConfig struct {
	ResourceUrl               string                                  `ConfigDef:"type=STRING,required=true"`
	Headers                   map[string]string                       `ConfigDef:"type=MAP,required=true"`
	SingleRequestPerBatch     bool                                    `ConfigDef:"type=BOOLEAN,required=true"`
	Client                    httpcommon.ClientConfigBean             `ConfigDefBean:"client"`
	DataFormat                string                                  `ConfigDef:"type=STRING,required=true"`
	DataGeneratorFormatConfig datagenerator.DataGeneratorFormatConfig `ConfigDefBean:"dataGeneratorFormatConfig"`
}

func init() {
	stagelibrary.SetCreator(Library, StageName, func() api.Stage {
		return &HttpClientDestination{BaseStage: &common.BaseStage{}, HttpCommon: &httpcommon.HttpCommon{}}
	})
}

func (h *HttpClientDestination) Init(stageContext api.StageContext) []validation.Issue {
	issues := h.BaseStage.Init(stageContext)
	if err := h.InitializeClient(h.Conf.Client); err != nil {
		issues = append(issues, stageContext.CreateConfigIssue(err.Error()))
		return issues
	}
	return h.Conf.DataGeneratorFormatConfig.Init(h.Conf.DataFormat, stageContext, issues)
}

func (h *HttpClientDestination) Write(batch api.Batch) error {
	if h.Conf.SingleRequestPerBatch && len(batch.GetRecords()) > 0 {
		return h.writeSingleRequestPerBatch(batch)
	} else {
		return h.writeSingleRequestPerRecord(batch)
	}
}

func (h *HttpClientDestination) writeSingleRequestPerBatch(batch api.Batch) error {
	var err error
	recordWriterFactory := h.Conf.DataGeneratorFormatConfig.RecordWriterFactory
	batchBuffer := bytes.NewBuffer([]byte{})
	recordWriter, err := recordWriterFactory.CreateWriter(h.GetStageContext(), batchBuffer)
	if err != nil {
		log.Error(err.Error())
		h.GetStageContext().ReportError(err)
		return nil
	}
	for _, record := range batch.GetRecords() {
		err = recordWriter.WriteRecord(record)
		if err != nil {
			log.Error(err.Error())
			h.GetStageContext().ToError(err, record)
		}
	}
	_ = recordWriter.Flush()
	_ = recordWriter.Close()
	err = h.sendToSDC(batchBuffer.Bytes())

	if err != nil {
		log.Error(err.Error())
		for _, record := range batch.GetRecords() {
			h.GetStageContext().ToError(err, record)
		}
	}

	return nil
}

func (h *HttpClientDestination) writeSingleRequestPerRecord(batch api.Batch) error {
	recordWriterFactory := h.Conf.DataGeneratorFormatConfig.RecordWriterFactory
	for _, record := range batch.GetRecords() {
		recordBuffer := bytes.NewBuffer([]byte{})
		recordWriter, err := recordWriterFactory.CreateWriter(h.GetStageContext(), recordBuffer)
		if err != nil {
			log.Error(err.Error())
			h.GetStageContext().ReportError(err)
			continue
		}
		err = recordWriter.WriteRecord(record)
		if err != nil {
			log.Error(err.Error())
			h.GetStageContext().ReportError(err)
			continue
		}
		_ = recordWriter.Flush()
		_ = recordWriter.Close()
		err = h.sendToSDC(recordBuffer.Bytes())
		if err != nil {
			log.Error(err.Error())
			h.GetStageContext().ToError(err, record)
		}
	}
	return nil
}

func (h *HttpClientDestination) sendToSDC(jsonValue []byte) error {
	var buf bytes.Buffer

	if h.Conf.Client.HttpCompression == "GZIP" {
		gz := gzip.NewWriter(&buf)
		if _, err := gz.Write(jsonValue); err != nil {
			return err
		}
		_ = gz.Close()
	} else {
		buf = *bytes.NewBuffer(jsonValue)
	}

	req, err := http.NewRequest("POST", h.Conf.ResourceUrl, &buf)
	if h.Conf.Headers != nil {
		for key, value := range h.Conf.Headers {
			req.Header.Set(key, value)
		}
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if h.Conf.Client.HttpCompression == "GZIP" {
		req.Header.Set("Content-Encoding", "gzip")
	}

	resp, err := h.RoundTrip(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.WithField("status", resp.Status).Debug("Response status")
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	return nil
}
