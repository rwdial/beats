package mysql

import (
	"github.com/rwdial/beats/packetbeat/config"
	"github.com/rwdial/beats/packetbeat/protos"
)

type mysqlConfig struct {
	config.ProtocolCommon `config:",inline"`
	MaxRowLength          int `config:"max_row_length"`
	MaxRows               int `config:"max_rows"`
}

var (
	defaultConfig = mysqlConfig{
		ProtocolCommon: config.ProtocolCommon{
			TransactionTimeout: protos.DefaultTransactionTimeout,
		},
		MaxRowLength: 1024,
		MaxRows:      10,
	}
)
