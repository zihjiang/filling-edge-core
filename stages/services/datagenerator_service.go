
package services

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"datacollector-edge/api"
	"datacollector-edge/api/dataformats"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/el"
	"datacollector-edge/stages/lib/datagenerator"
	"datacollector-edge/stages/stagelibrary"
	"io"
)

type DataGeneratorServiceImpl struct {
	stageContext              api.StageContext
	DataFormat                string                                  `ConfigDef:"type=STRING,required=true"`
	DataGeneratorFormatConfig datagenerator.DataGeneratorFormatConfig `ConfigDefBean:"dataGeneratorFormatConfig"`
}

func init() {
	stagelibrary.SetServiceCreator(dataformats.DataFormatGeneratorServiceName, func() api.Service {
		return &DataGeneratorServiceImpl{}
	})
}

func (d *DataGeneratorServiceImpl) Init(stageContext api.StageContext) []validation.Issue {
	d.stageContext = stageContext
	issues := make([]validation.Issue, 0)
	log.Debug("DataGeneratorServiceImpl Init method")
	d.DataGeneratorFormatConfig.Init(d.DataFormat, stageContext, issues)
	return issues
}

func (d *DataGeneratorServiceImpl) GetGenerator(writer io.Writer) (dataformats.RecordWriter, error) {
	recordWriterFactory := d.DataGeneratorFormatConfig.RecordWriterFactory
	return recordWriterFactory.CreateWriter(d.stageContext, writer)
}

func (d *DataGeneratorServiceImpl) Destroy() error {
	return nil
}

func (d *DataGeneratorServiceImpl) IsWholeFileFormat() bool {
	return d.DataFormat == "WHOLE_FILE"
}

func (d *DataGeneratorServiceImpl) GetWholeFileName(record api.Record) (string, error) {
	recordContext := context.WithValue(context.Background(), el.RecordContextVar, record)
	result, err := d.stageContext.Evaluate(d.DataGeneratorFormatConfig.FileNameEL, "fileNameEl", recordContext)
	if err != nil {
		return "", err
	}
	return cast.ToString(result), nil
}

func (d *DataGeneratorServiceImpl) GetWholeFileExistsAction() string {
	return d.DataGeneratorFormatConfig.WholeFileExistsAction
}

func (d *DataGeneratorServiceImpl) GetIncludeChecksumInTheEvents() bool {
	return d.DataGeneratorFormatConfig.IncludeChecksumInTheEvents
}

func (d *DataGeneratorServiceImpl) GetChecksumAlgorithm() string {
	return d.DataGeneratorFormatConfig.ChecksumAlgorithm
}
