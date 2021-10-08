
package processors

import (
	_ "datacollector-edge/stages/processors/delay"
	_ "datacollector-edge/stages/processors/expression"
	_ "datacollector-edge/stages/processors/fieldremover"
	_ "datacollector-edge/stages/processors/http"
	_ "datacollector-edge/stages/processors/identity"
	_ "datacollector-edge/stages/processors/javascript"
	_ "datacollector-edge/stages/processors/random_error"
	_ "datacollector-edge/stages/processors/selector"
)
