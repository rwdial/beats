package main

import (
	"os"

	"github.com/rwdial/beats/libbeat/beat"
	"github.com/rwdial/beats/topbeat/beater"
)

var Name = "topbeat"

func main() {
	if err := beat.Run(Name, "", beater.New()); err != nil {
		os.Exit(1)
	}
}
