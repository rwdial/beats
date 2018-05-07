package main

import (
	"os"

	"github.com/rwdial/beats/libbeat/beat"
	"github.com/rwdial/beats/packetbeat/beater"

	// import support protocol modules
	_ "github.com/rwdial/beats/packetbeat/protos/amqp"
	_ "github.com/rwdial/beats/packetbeat/protos/dns"
	_ "github.com/rwdial/beats/packetbeat/protos/http"
	_ "github.com/rwdial/beats/packetbeat/protos/memcache"
	_ "github.com/rwdial/beats/packetbeat/protos/mongodb"
	_ "github.com/rwdial/beats/packetbeat/protos/mysql"
	_ "github.com/rwdial/beats/packetbeat/protos/pgsql"
	_ "github.com/rwdial/beats/packetbeat/protos/redis"
	_ "github.com/rwdial/beats/packetbeat/protos/thrift"
)

var Name = "packetbeat"

// Setups and Runs Packetbeat
func main() {
	if err := beat.Run(Name, "", beater.New()); err != nil {
		os.Exit(1)
	}
}
