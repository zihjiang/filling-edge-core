
package el

import (
	"errors"
	"fmt"
	"github.com/madhukard/govaluate"
	"os"
)

type SdcEL struct {
}

func (j *SdcEL) GetHostName(args ...interface{}) (interface{}, error) {
	if len(args) != 0 {
		return "", errors.New(
			fmt.Sprintf("The function 'sdc:hostname' requires 0 arguments but was passed %d", len(args)),
		)
	}

	return os.Hostname()
}

func (j *SdcEL) GetELFunctionDefinitions() map[string]govaluate.ExpressionFunction {
	functions := map[string]govaluate.ExpressionFunction{
		"sdc:hostname": j.GetHostName,
	}
	return functions
}
