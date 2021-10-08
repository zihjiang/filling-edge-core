
package common

type ServiceDefinition struct {
	Name                 string
	Version              string
	ConfigDefinitionsMap map[string]*ConfigDefinition
}
