
package services

import (
	log "github.com/sirupsen/logrus"
	"datacollector-edge/api"
	"datacollector-edge/api/dataformats"
	"datacollector-edge/api/validation"
	"datacollector-edge/stages/lib/dataparser"
	"datacollector-edge/stages/stagelibrary"
	"io"
)

type DataParserServiceImpl struct {
	stageContext     api.StageContext
	DataFormat       string                            `ConfigDef:"type=STRING,required=true"`
	DataFormatConfig dataparser.DataParserFormatConfig `ConfigDefBean:"dataFormatConfig"`
}

func init() {
	stagelibrary.SetServiceCreator(dataformats.DataFormatParserServiceName, func() api.Service {
		return &DataParserServiceImpl{}
	})
}

func (d *DataParserServiceImpl) Init(stageContext api.StageContext) []validation.Issue {
	d.stageContext = stageContext
	issues := make([]validation.Issue, 0)
	log.Debug("DataParserService Init method")
	d.DataFormatConfig.Init(d.DataFormat, stageContext, issues)
	return issues
}

func (d *DataParserServiceImpl) GetParser(messageId string, reader io.Reader) (dataformats.RecordReader, error) {
	recordReaderFactory := d.DataFormatConfig.RecordReaderFactory
	return recordReaderFactory.CreateReader(d.stageContext, reader, messageId)
}

func (b *DataParserServiceImpl) Destroy() error {
	return nil
}

func GetDataFormatParserServiceName() string {
	return dataformats.DataFormatParserServiceName
}
