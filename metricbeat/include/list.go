/*

This file is included in the main file to load all metricsets.

In case only a subset of metricsets should be included, they can be specified manually in the main.go file.

*/
package include

// Make sure all active plugins are loaded
// TODO: create a script to automatically generate this list
import (
	"github.com/rwdial/beats/libbeat/logp"
	"github.com/rwdial/beats/metricbeat/helper"

	// List of all metrics to make sure they are registered
	// Every new metric must be added here
	_ "github.com/rwdial/beats/metricbeat/module/apache/status"
	_ "github.com/rwdial/beats/metricbeat/module/mysql/status"

	// Redis module and metrics
	_ "github.com/rwdial/beats/metricbeat/module/redis"
	_ "github.com/rwdial/beats/metricbeat/module/redis/info"
)

func ListAll() {
	logp.Debug("beat", "Registered Modules and Metrics")
	for module := range helper.Registry.Modulers {
		for metricset := range helper.Registry.MetricSeters[module] {
			logp.Debug("metricbeat", "Registred: Module: %v, MetricSet: %v", module, metricset)
		}
	}
}
