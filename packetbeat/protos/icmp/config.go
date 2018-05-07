package icmp

import "github.com/rwdial/beats/packetbeat/protos"

type icmpConfig struct {
	SendRequest        bool `config:"send_request"`
	SendResponse       bool `config:"send_response"`
	TransactionTimeout int  `config:"transaction_timeout"`
}

var (
	defaultConfig = icmpConfig{
		TransactionTimeout: protos.DefaultTransactionTimeout,
	}
)
