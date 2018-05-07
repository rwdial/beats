package mongodb

import (
	"github.com/rwdial/beats/packetbeat/config"
	"github.com/rwdial/beats/packetbeat/protos"
)

type mongodbConfig struct {
	config.ProtocolCommon `config:",inline"`
	MaxDocLength          int `config:"max_doc_length"`
	MaxDocs               int `config:"max_docs"`
}

var (
	defaultConfig = mongodbConfig{
		ProtocolCommon: config.ProtocolCommon{
			TransactionTimeout: protos.DefaultTransactionTimeout,
		},
		MaxDocLength: 5000,
		MaxDocs:      10,
	}
)
