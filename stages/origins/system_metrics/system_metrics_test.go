
package system_metrics

import (
	"testing"

	"datacollector-edge/api"
	"datacollector-edge/api/fieldtype"
	"datacollector-edge/container/common"
	"datacollector-edge/container/creation"
	"datacollector-edge/container/execution/runner"
)

func getStageContext(
	configuration []common.Config,
	parameters map[string]interface{},
) *common.StageContextImpl {
	stageConfig := common.StageConfiguration{}
	stageConfig.Library = Library
	stageConfig.StageName = StageName
	stageConfig.Configuration = configuration
	errorSink := common.NewErrorSink()
	return &common.StageContextImpl{
		StageConfig: &stageConfig,
		Parameters:  parameters,
		ErrorSink:   errorSink,
	}
}

func TestOrigin_Init(t *testing.T) {
	configuration := []common.Config{
		{
			Name:  "conf.delay",
			Value: float64(2000),
		},
	}

	stageContext := getStageContext(configuration, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	if stageInstance.(*Origin).Conf.Delay != float64(2000) {
		t.Error("Failed to inject config value for delay")
	}
}

func TestOrigin_Produce_Fetch_HostInfo(t *testing.T) {
	configuration := []common.Config{
		{
			Name:  "conf.fetchHostInfo",
			Value: true,
		},
	}

	stageContext := getStageContext(configuration, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	issues := stageInstance.Init(stageContext)
	if len(issues) > 0 {
		t.Error(issues)
	}

	batchMaker := runner.NewBatchMakerImpl(runner.StagePipe{}, false)
	_, err = stageInstance.(api.Origin).Produce(&defaultOffset, 1, batchMaker)
	if err != nil {
		t.Error("Err :", err)
		return
	}

	records := batchMaker.GetStageOutput()
	if len(records) != 1 {
		t.Error("Expected 1 records but got - ", len(records))
		return
	}

	rootField, _ := records[0].Get()
	mapFieldValue := rootField.Value.(map[string]*api.Field)
	if mapFieldValue["timestamp"] == nil {
		t.Error("Failed to inject timestamp value")
	}
	if mapFieldValue["hostInfo"] == nil || mapFieldValue["hostInfo"].Type != fieldtype.MAP {
		t.Error("Failed to fetch Host Informatinon")
	}

	stageInstance.Destroy()
}

func TestOrigin_Produce_Fetch_CPU(t *testing.T) {
	configuration := []common.Config{
		{
			Name:  "conf.fetchCpuStats",
			Value: true,
		},
	}

	stageContext := getStageContext(configuration, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	issues := stageInstance.Init(stageContext)
	if len(issues) > 0 {
		t.Error(issues)
	}

	batchMaker := runner.NewBatchMakerImpl(runner.StagePipe{}, false)
	_, err = stageInstance.(api.Origin).Produce(&defaultOffset, 1, batchMaker)
	if err != nil {
		t.Error("Err :", err)
		return
	}

	records := batchMaker.GetStageOutput()
	if len(records) != 1 {
		t.Error("Expected 1 records but got - ", len(records))
		return
	}

	rootField, _ := records[0].Get()
	mapFieldValue := rootField.Value.(map[string]*api.Field)
	if mapFieldValue["timestamp"] == nil {
		t.Error("Failed to inject timestamp value")
	}
	if mapFieldValue["cpu"] == nil || mapFieldValue["cpu"].Type != fieldtype.MAP {
		t.Error("Failed to fetch CPU System metrics")
	}

	stageInstance.Destroy()
}

func TestOrigin_Produce_Fetch_Memory(t *testing.T) {
	configuration := []common.Config{
		{
			Name:  "conf.fetchMemStats",
			Value: true,
		},
	}

	stageContext := getStageContext(configuration, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	issues := stageInstance.Init(stageContext)
	if len(issues) > 0 {
		t.Error(issues)
	}

	batchMaker := runner.NewBatchMakerImpl(runner.StagePipe{}, false)
	_, err = stageInstance.(api.Origin).Produce(&defaultOffset, 1, batchMaker)
	if err != nil {
		t.Error("Err :", err)
		return
	}

	records := batchMaker.GetStageOutput()
	if len(records) != 1 {
		t.Error("Expected 1 records but got - ", len(records))
		return
	}

	rootField, _ := records[0].Get()
	mapFieldValue := rootField.Value.(map[string]*api.Field)
	if mapFieldValue["timestamp"] == nil {
		t.Error("Failed to inject timestamp value")
	}
	if mapFieldValue["memory"] == nil || mapFieldValue["memory"].Type != fieldtype.MAP {
		t.Error("Failed to fetch memory System metrics")
	}

	stageInstance.Destroy()
}

func TestOrigin_Produce_Fetch_Disk(t *testing.T) {
	configuration := []common.Config{
		{
			Name:  "conf.fetchDiskStats",
			Value: true,
		},
	}

	stageContext := getStageContext(configuration, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	issues := stageInstance.Init(stageContext)
	if len(issues) > 0 {
		t.Error(issues)
	}

	batchMaker := runner.NewBatchMakerImpl(runner.StagePipe{}, false)
	_, err = stageInstance.(api.Origin).Produce(&defaultOffset, 1, batchMaker)
	if err != nil {
		t.Error("Err :", err)
		return
	}

	records := batchMaker.GetStageOutput()
	if len(records) != 1 {
		t.Error("Expected 1 records but got - ", len(records))
		return
	}

	rootField, _ := records[0].Get()
	mapFieldValue := rootField.Value.(map[string]*api.Field)
	if mapFieldValue["timestamp"] == nil {
		t.Error("Failed to inject timestamp value")
	}
	if mapFieldValue["disk"] == nil || mapFieldValue["disk"].Type != fieldtype.MAP {
		t.Error("Failed to fetch disk System metrics")
	}

	stageInstance.Destroy()
}

func TestOrigin_Produce_Fetch_Network(t *testing.T) {
	configuration := []common.Config{
		{
			Name:  "conf.fetchNetStats",
			Value: true,
		},
	}

	stageContext := getStageContext(configuration, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	issues := stageInstance.Init(stageContext)
	if len(issues) > 0 {
		t.Error(issues)
	}

	batchMaker := runner.NewBatchMakerImpl(runner.StagePipe{}, false)
	_, err = stageInstance.(api.Origin).Produce(&defaultOffset, 1, batchMaker)
	if err != nil {
		t.Error("Err :", err)
		return
	}

	records := batchMaker.GetStageOutput()
	if len(records) != 1 {
		t.Error("Expected 1 records but got - ", len(records))
		return
	}

	rootField, _ := records[0].Get()
	mapFieldValue := rootField.Value.(map[string]*api.Field)
	if mapFieldValue["timestamp"] == nil {
		t.Error("Failed to inject timestamp value")
	}
	if mapFieldValue["network"] == nil || mapFieldValue["network"].Type != fieldtype.MAP {
		t.Error("Failed to fetch network System metrics")
	}

	stageInstance.Destroy()
}

func TestOrigin_Produce_Fetch_ProcessStats(t *testing.T) {
	configuration := []common.Config{
		{
			Name:  "conf.fetchProcessStats",
			Value: true,
		},
		{
			Name:  "conf.processConf.processRegexStr",
			Value: "Test.*",
		},
		{
			Name:  "conf.processConf.userRegexStr",
			Value: ".*",
		},
	}

	stageContext := getStageContext(configuration, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		t.Error(err)
		return
	}

	stageInstance := stageBean.Stage
	if stageInstance == nil {
		t.Error("Failed to create stage instance")
	}

	issues := stageInstance.Init(stageContext)
	if len(issues) > 0 {
		t.Error(issues)
	}

	batchMaker := runner.NewBatchMakerImpl(runner.StagePipe{}, false)
	_, err = stageInstance.(api.Origin).Produce(&defaultOffset, 1, batchMaker)
	if err != nil {
		t.Error("Err :", err)
		return
	}

	records := batchMaker.GetStageOutput()
	if len(records) != 1 {
		t.Error("Expected 1 records but got - ", len(records))
		return
	}

	rootField, _ := records[0].Get()
	mapFieldValue := rootField.Value.(map[string]*api.Field)
	if mapFieldValue["timestamp"] == nil {
		t.Error("Failed to inject timestamp value")
	}
	if mapFieldValue["process"] == nil || mapFieldValue["process"].Type != fieldtype.LIST {
		t.Error("Failed to fetch process metrics")
	}

	stageInstance.Destroy()
}

// Run Benchmark Tests
//    go test -run=^$ -bench=. -memprofile=mem0.out -cpuprofile=cpu0.out
// Profile CPU
// 	  go tool pprof bench.test cpu0.out
// Profile Memory
// 	  go tool pprof --alloc_space bench.test mem0.out
func BenchmarkOrigin_Produce(b *testing.B) {
	configuration := []common.Config{
		{
			Name:  "conf.fetchHostInfo",
			Value: true,
		},
		{
			Name:  "conf.fetchCpuStats",
			Value: true,
		},
		{
			Name:  "conf.fetchMemStats",
			Value: true,
		},
		{
			Name:  "conf.fetchDiskStats",
			Value: true,
		},
		{
			Name:  "conf.fetchDiskStats",
			Value: true,
		},
	}

	stageContext := getStageContext(configuration, nil)
	stageBean, err := creation.NewStageBean(stageContext.StageConfig, stageContext.Parameters, nil)
	if err != nil {
		panic(err)
	}

	stageInstance := stageBean.Stage
	if stageInstance == nil {
		panic("Failed to create stage instance")
	}

	issues := stageInstance.Init(stageContext)
	if len(issues) > 0 {
		panic(issues)
	}

	batchMaker := runner.NewBatchMakerImpl(runner.StagePipe{}, false)
	_, err = stageInstance.(api.Origin).Produce(&defaultOffset, 1, batchMaker)
	if err != nil {
		panic(err)
	}

	stageInstance.Destroy()
}
