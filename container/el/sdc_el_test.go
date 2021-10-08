
package el

import (
	"os"
	"testing"
)

func TestSdcEL(test *testing.T) {
	hostName, _ := os.Hostname()
	evaluationTests := []EvaluationTest{
		{
			Name:       "Test sdc:hostname()",
			Expression: "${sdc:hostname()}",
			Expected:   hostName,
		},
		{
			Name:       "Test function sdc:hostname() - Error 1",
			Expression: "${sdc:hostname('invalid param')}",
			Expected:   "The function 'sdc:hostname' requires 0 arguments but was passed 1",
			ErrorCase:  true,
		},
	}
	RunEvaluationTests(evaluationTests, []Definitions{&SdcEL{}}, test)
}
