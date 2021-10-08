
package datagenerator

import (
	"datacollector-edge/api"
	"datacollector-edge/api/validation"
	"datacollector-edge/container/recordio"
	"datacollector-edge/container/recordio/binaryrecord"
	"datacollector-edge/container/recordio/jsonrecord"
	"datacollector-edge/container/recordio/sdcrecord"
	"datacollector-edge/container/recordio/textrecord"
)

type DataGeneratorFormatConfig struct {
	/* Charset Related -- Shown last */
	Charset string `ConfigDef:"type=STRING,required=true"`

	/** For DELIMITED Content **/
	CsvFileFormat            string `ConfigDef:"type=STRING,required=true"`
	CsvHeader                string `ConfigDef:"type=STRING,required=true"`
	CsvReplaceNewLines       bool   `ConfigDef:"type=BOOLEAN,required=true"`
	CsvReplaceNewLinesString string `ConfigDef:"type=STRING,required=true"`
	CsvCustomDelimiter       string `ConfigDef:"type=STRING,required=true"`
	CsvCustomEscape          string `ConfigDef:"type=STRING,required=true"`
	CsvCustomQuote           string `ConfigDef:"type=STRING,required=true"`

	/** For JSON **/
	JsonMode string `ConfigDef:"type=STRING,required=true"`

	/** For TEXT Content **/
	TextFieldPath          string `ConfigDef:"type=STRING,required=true"`
	TextRecordSeparator    string `ConfigDef:"type=STRING,required=true"`
	TextFieldMissingAction string `ConfigDef:"type=STRING,required=true"`
	TextEmptyLineIfNull    bool   `ConfigDef:"type=BOOLEAN,required=true"`

	/** For AVRO Content **/
	AvroSchemaSource                  string   `ConfigDef:"type=STRING,required=true"`
	AvroSchema                        string   `ConfigDef:"type=STRING,required=true"`
	RegisterSchema                    bool     `ConfigDef:"type=BOOLEAN,required=true"`
	SchemaRegistryUrlsForRegistration []string `ConfigDef:"type=LIST,required=true"`
	SchemaRegistryUrls                []string `ConfigDef:"type=LIST,required=true"`
	SchemaLookupMode                  string   `ConfigDef:"type=STRING,required=true"`
	SubjectToRegister                 string   `ConfigDef:"type=STRING,required=true"`
	SchemaId                          float64  `ConfigDef:"type=NUMBER,required=true"`
	IncludeSchema                     bool     `ConfigDef:"type=BOOLEAN,required=true"`
	AvroCompression                   string   `ConfigDef:"type=STRING,required=true"`

	/** For Binary Content **/
	BinaryFieldPath string `ConfigDef:"type=STRING,required=true"`

	/** For Protobuf Content **/
	ProtoDescriptorFile string `ConfigDef:"type=STRING,required=true"`
	MessageType         string `ConfigDef:"type=STRING,required=true"`

	/** For Whole File Content **/
	FileNameEL                 string `ConfigDef:"type=STRING,required=true,evaluation=EXPLICIT"`
	WholeFileExistsAction      string `ConfigDef:"type=STRING,required=true"`
	IncludeChecksumInTheEvents bool   `ConfigDef:"type=BOOLEAN,required=true"`
	ChecksumAlgorithm          string `ConfigDef:"type=STRING,required=true"`

	/** For XML Content **/
	XmlPrettyPrint    bool   `ConfigDef:"type=BOOLEAN,required=true"`
	XmlValidateSchema bool   `ConfigDef:"type=BOOLEAN,required=true"`
	XmlSchema         string `ConfigDef:"type=STRING,required=true"`
	IsDelimited       bool   `ConfigDef:"type=BOOLEAN,required=true"`

	RecordWriterFactory recordio.RecordWriterFactory
}

func (d *DataGeneratorFormatConfig) Init(
	dataFormat string,
	stageContext api.StageContext,
	issues []validation.Issue,
) []validation.Issue {
	switch dataFormat {
	case "TEXT":
		d.RecordWriterFactory = &textrecord.TextWriterFactoryImpl{TextFieldPath: d.TextFieldPath}
	case "JSON":
		d.RecordWriterFactory = &jsonrecord.JsonWriterFactoryImpl{Mode: d.JsonMode}
	case "BINARY":
		d.RecordWriterFactory = &binaryrecord.BinaryWriterFactoryImpl{BinaryFieldPath: d.BinaryFieldPath}
	case "WHOLE_FILE":
		// Supported format
	case "SDC_JSON":
		d.RecordWriterFactory = &sdcrecord.SDCRecordWriterFactoryImpl{}
	default:
		issues = append(issues, stageContext.CreateConfigIssue("Unsupported Data Format - "+dataFormat))
	}
	return issues
}
