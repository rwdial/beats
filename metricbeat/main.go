package main

import (
	"os"

	"github.com/rwdial/beats/metricbeat/beater"
	_ "github.com/rwdial/beats/metricbeat/include"

	"github.com/rwdial/beats/libbeat/beat"
)

var Name = "metricbeat"

func main() {
	if err := beat.Run(Name, "", beater.New()); err != nil {
		os.Exit(1)
	}
}
