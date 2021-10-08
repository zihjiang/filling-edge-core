//// +build kafka


package kafka

import (
	"github.com/Shopify/sarama"
	"datacollector-edge/container/common"
	"datacollector-edge/container/creation"
	"testing"
	"time"
)

func getStageContext(
	stageConfigurationList []common.Config,
	parameters map[string]interface{},
) *common.StageContextImpl {
	stageConfig := common.StageConfiguration{}
	stageConfig.Library = APACHE_KAFKA_1_1_LIBRARY
	stageConfig.StageName = StageName
	stageConfig.Configuration = stageConfigurationList
	errorSink := common.NewErrorSink()
	return &common.StageContextImpl{
		StageConfig: &stageConfig,
		Parameters:  parameters,
		ErrorSink:   errorSink,
	}
}

func TestKafkaDestination_Init(t *testing.T) {
	stageContext := getStageContext(getTestConfig(), nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	if stageInstance.(*KafkaDestination).Conf.MetadataBrokerList != "10.10.14.43:9092" {
		t.Error("Failed to inject config value for MetadataBrokerList")
	}

	if stageInstance.(*KafkaDestination).Conf.Topic != "edgetest" {
		t.Error("Failed to inject config value for topic")
	}
}

func TestKafkaDestination_mapJVMConfigsToSaramaConfig(t *testing.T) {
	stageContext := getStageContext(getTestConfig(), nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	kafkaDestInstance := stageInstance.(*KafkaDestination)

	if kafkaDestInstance.Conf.MetadataBrokerList != "10.10.14.43:9092" {
		t.Error("Failed to inject config value for MetadataBrokerList")
	}

	if kafkaDestInstance.Conf.Topic != "edgetest" {
		t.Error("Failed to inject config value for topic")
	}

	err = kafkaDestInstance.mapJVMConfigsToSaramaConfig()
	if err != nil {
		t.Error(err)
	}

	if kafkaDestInstance.kafkaClientConf.Net.DialTimeout != 40000*time.Millisecond {
		t.Errorf(
			"Failed to inject sarama config, expected: %f, got: %f",
			40000*time.Millisecond,
			kafkaDestInstance.kafkaClientConf.Net.DialTimeout,
		)
	}

	if !kafkaDestInstance.kafkaClientConf.Net.TLS.Enable {
		t.Error("Failed to inject sarama config TLS enable")
	}

	if !kafkaDestInstance.kafkaClientConf.Net.SASL.Enable {
		t.Error("Failed to inject sarama config SASL enable")
	}

	if kafkaDestInstance.kafkaClientConf.Net.SASL.User != "sampleUser" {
		t.Error("Failed to inject sarama config SASL user")
	}

	if kafkaDestInstance.kafkaClientConf.Net.SASL.Password != "samplePassword" {
		t.Error("Failed to inject sarama config SASL password")
	}

	if kafkaDestInstance.kafkaClientConf.Producer.Retry.Max != 5 {
		t.Error("Failed to inject sarama config Retry max")
	}

	if kafkaDestInstance.kafkaClientConf.Producer.MaxMessageBytes != 10000 {
		t.Error("Failed to inject sarama config Max Message Bytes")
	}

	if kafkaDestInstance.kafkaClientConf.Producer.RequiredAcks != sarama.WaitForLocal {
		t.Error("Failed to inject sarama config RequiredAcks")
	}

	if kafkaDestInstance.kafkaClientConf.Producer.Timeout != 25000*time.Millisecond {
		t.Error("Failed to inject sarama config Timeout")
	}

	if kafkaDestInstance.kafkaClientConf.Producer.Compression != sarama.CompressionSnappy {
		t.Error("Failed to inject sarama config Compression")
	}

	if kafkaDestInstance.kafkaClientConf.Producer.Flush.Frequency != 12500*time.Millisecond {
		t.Error("Failed to inject sarama config Flush Frequency")
	}

	if kafkaDestInstance.kafkaClientConf.Producer.Flush.MaxMessages != 1500 {
		t.Error("Failed to inject sarama config Flush MaxMessages")
	}

	if kafkaDestInstance.kafkaClientConf.Producer.Retry.Max != 5 {
		t.Error("Failed to inject sarama config Retry max")
	}

	if kafkaDestInstance.kafkaClientConf.Producer.Retry.Backoff != 50000*time.Millisecond {
		t.Error("Failed to inject sarama config Retry Backoff time")
	}

}

func getTestConfig() []common.Config {
	kafkaProducerConfigs := make([]interface{}, 12)
	kafkaProducerConfigs[0] = map[string]interface{}{
		"key":   "socket.timeout.ms",
		"value": "40000",
	}
	kafkaProducerConfigs[1] = map[string]interface{}{
		"key":   "ssl.endpoint.identification.algorithm",
		"value": "https",
	}
	kafkaProducerConfigs[2] = map[string]interface{}{
		"key":   "security.protocol",
		"value": "SASL_SSL",
	}
	kafkaProducerConfigs[3] = map[string]interface{}{
		"key":   "sasl.jaas.config",
		"value": `org.apache.kafka.common.security.plain.PlainLoginModule required username="sampleUser" password="samplePassword";`,
	}

	kafkaProducerConfigs[4] = map[string]interface{}{
		"key":   "message.max.bytes",
		"value": "10000",
	}

	kafkaProducerConfigs[5] = map[string]interface{}{
		"key":   "request.required.acks",
		"value": "1",
	}

	kafkaProducerConfigs[6] = map[string]interface{}{
		"key":   "request.timeout.ms",
		"value": "25000",
	}

	kafkaProducerConfigs[7] = map[string]interface{}{
		"key":   "compression.type",
		"value": "snappy",
	}

	kafkaProducerConfigs[8] = map[string]interface{}{
		"key":   "queue.buffering.max.ms",
		"value": "12500",
	}
	kafkaProducerConfigs[9] = map[string]interface{}{
		"key":   "queue.buffering.max.messages",
		"value": "1500",
	}

	kafkaProducerConfigs[10] = map[string]interface{}{
		"key":   "message.send.max.retries",
		"value": "5",
	}

	kafkaProducerConfigs[11] = map[string]interface{}{
		"key":   "retry.backoff.ms",
		"value": "50000",
	}

	configuration := []common.Config{
		{
			Name:  "conf.metadataBrokerList",
			Value: "192.168.1.70:9094",
		},
		{
			Name:  "conf.topic",
			Value: "edgetest",
		},
		{
			Name:  "conf.partitionStrategy",
			Value: "ROUND_ROBIN",
		},
		{
			Name:  "conf.dataFormat",
			Value: "JSON",
		},
		{
			Name:  "conf.kafkaProducerConfigs",
			Value: kafkaProducerConfigs,
		},
	}

	return configuration
}
