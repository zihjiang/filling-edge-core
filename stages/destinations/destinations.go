
package destinations

import (
	_ "datacollector-edge/stages/destinations/azure/eventhubs"
	_ "datacollector-edge/stages/destinations/coap"
	_ "datacollector-edge/stages/destinations/firehose"
	_ "datacollector-edge/stages/destinations/http"
	_ "datacollector-edge/stages/destinations/influxdb"
	_ "datacollector-edge/stages/destinations/kafka"
	_ "datacollector-edge/stages/destinations/kinesis"
	_ "datacollector-edge/stages/destinations/mqtt"
	_ "datacollector-edge/stages/destinations/s3"
	_ "datacollector-edge/stages/destinations/toerror"
	_ "datacollector-edge/stages/destinations/toevent"
	_ "datacollector-edge/stages/destinations/trash"
	_ "datacollector-edge/stages/destinations/websocket"
)
