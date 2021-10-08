

package creation

import (
	"datacollector-edge/api"
	"datacollector-edge/container/common"
)

type ServiceBean struct {
	Config  *common.ServiceConfiguration
	Service api.Service
}
