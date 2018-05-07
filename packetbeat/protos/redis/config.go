package redis

import (
	"github.com/rwdial/beats/packetbeat/config"
	"github.com/rwdial/beats/packetbeat/protos"
)

type redisConfig struct {
	config.ProtocolCommon `config:",inline"`
}

var (
	defaultConfig = redisConfig{
		ProtocolCommon: config.ProtocolCommon{
			TransactionTimeout: protos.DefaultTransactionTimeout,
		},
	}
)
