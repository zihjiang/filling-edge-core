
package websocket

import (
	"bytes"
	"errors"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"datacollector-edge/api"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/common"
	"datacollector-edge/stages/lib/datagenerator"
	"datacollector-edge/stages/stagelibrary"
	"net/http"
)

const (
	LIBRARY    = "streamsets-datacollector-basic-lib"
	STAGE_NAME = "com_streamsets_pipeline_stage_destination_websocket_WebSocketDTarget"
)

type WebSocketClientDestination struct {
	*common.BaseStage
	Conf WebSocketTargetConfig `ConfigDefBean:"conf"`
}

type WebSocketTargetConfig struct {
	ResourceUrl               string                                  `ConfigDef:"type=STRING,required=true"`
	Headers                   map[string]string                       `ConfigDef:"type=MAP,required=true"`
	DataFormat                string                                  `ConfigDef:"type=STRING,required=true"`
	DataGeneratorFormatConfig datagenerator.DataGeneratorFormatConfig `ConfigDefBean:"dataGeneratorFormatConfig"`
}

func init() {
	stagelibrary.SetCreator(LIBRARY, STAGE_NAME, func() api.Stage {
		return &WebSocketClientDestination{BaseStage: &common.BaseStage{}}
	})
}

func (w *WebSocketClientDestination) Init(stageContext api.StageContext) []validation.Issue {
	issues := w.BaseStage.Init(stageContext)
	log.Debug("WebSocketClientDestination Init method")
	return w.Conf.DataGeneratorFormatConfig.Init(w.Conf.DataFormat, stageContext, issues)
}

func (w *WebSocketClientDestination) Write(batch api.Batch) error {
	log.WithField("url", w.Conf.ResourceUrl).Debug("WebSocketClientDestination write method")
	recordWriterFactory := w.Conf.DataGeneratorFormatConfig.RecordWriterFactory
	if recordWriterFactory == nil {
		return errors.New("recordWriterFactory is null")
	}

	var requestHeader = http.Header{}
	if w.Conf.Headers != nil {
		for key, value := range w.Conf.Headers {
			requestHeader.Set(key, value)
		}
	}

	c, _, err := websocket.DefaultDialer.Dial(w.Conf.ResourceUrl, requestHeader)
	if err != nil {
		return err
	}

	for _, record := range batch.GetRecords() {
		recordBuffer := bytes.NewBuffer([]byte{})
		recordWriter, err := recordWriterFactory.CreateWriter(w.GetStageContext(), recordBuffer)
		if err != nil {
			return err
		}
		err = recordWriter.WriteRecord(record)
		if err != nil {
			return err
		}
		recordWriter.Flush()
		recordWriter.Close()

		err = c.WriteMessage(websocket.TextMessage, recordBuffer.Bytes())
		if err != nil {
			log.WithError(err).Error("Websocket write error")
			w.GetStageContext().ToError(err, record)
		}
	}

	defer c.Close()
	return nil
}
