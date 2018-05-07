package main

import (
	"os"

	"github.com/rwdial/beats/libbeat/beat"
	"github.com/rwdial/beats/libbeat/mock"
)

func main() {
	if err := beat.Run(mock.Name, mock.Version, mock.New()); err != nil {
		os.Exit(1)
	}
}
