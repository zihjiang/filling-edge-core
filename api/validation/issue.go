
package validation

type Issue struct {
	InstanceName   string            `json:"instanceName"`
	ConfigGroup    string            `json:"configGroup"`
	ConfigName     string            `json:"configName"`
	Message        string            `json:"message"`
	Level          string            `json:"level"`
	Count          int               `json:"count"`
	AdditionalInfo map[string]string `json:"additionalInfo"`
}
