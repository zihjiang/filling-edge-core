
package common

const (
	ConfigDefTagName      = "ConfigDef"
	ConfigDefBeanTagName  = "ConfigDefBean"
	ListBeanModelTagName  = "ListBeanModel"
	PredicateModelTagName = "PredicateModel"
	EvaluationExplicit    = "EXPLICIT"
	EvaluationImplicit    = "IMPLICIT"
)

type StageDefinition struct {
	Name                 string
	Library              string
	Version              string
	ConfigDefinitionsMap map[string]*ConfigDefinition
}

type ConfigDefinition struct {
	Name       string
	Type       string
	Required   bool
	FieldName  string
	Evaluation string
	Model      ModelDefinition
}

type ModelDefinition struct {
	ConfigDefinitionsMap map[string]*ConfigDefinition
}
