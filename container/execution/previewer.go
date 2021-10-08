
package execution

type Previewer interface {
	GetId() string
	ValidateConfigs(timeoutMillis int64) error
	Start(
		batches int,
		batchSize int,
		skipTargets bool,
		stopStage string,
		stagesOverride []StageOutputJson,
		timeoutMillis int64,
		testOrigin bool,
	) error
	Stop() error
	GetStatus() string
	GetOutput() PreviewOutput
}
