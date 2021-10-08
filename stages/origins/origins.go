
package origins

import (
	_ "datacollector-edge/stages/origins/dev_data_generator"
	_ "datacollector-edge/stages/origins/dev_random"
	_ "datacollector-edge/stages/origins/dev_rawdata"
	_ "datacollector-edge/stages/origins/filetail"
	_ "datacollector-edge/stages/origins/grpc_client"
	_ "datacollector-edge/stages/origins/httpclient"
	_ "datacollector-edge/stages/origins/httpserver"
	_ "datacollector-edge/stages/origins/mqtt"
	_ "datacollector-edge/stages/origins/sensor_reader"
	_ "datacollector-edge/stages/origins/spooler"
	_ "datacollector-edge/stages/origins/system_metrics"
	_ "datacollector-edge/stages/origins/websocketclient"
	_ "datacollector-edge/stages/origins/windows"
)
