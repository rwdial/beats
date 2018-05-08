package elasticsearch

import "github.com/rwdial/beats/libbeat/outputs"

type elasticsearchConfig struct {
	Protocol         string             `config:"protocol"`
	Path             string             `config:"path"`
	Params           map[string]string  `config:"parameters"`
	Username         string             `config:"username"`
	Password         string             `config:"password"`
	ProxyURL         string             `config:"proxy_url"`
	Index            string             `config:"index"`
	IncludeDatestamp bool               `config:"include_datestamp"`
	IDColumn         string             `config:"id_column"`
	DeleteColumn     string             `config:"delete_column"`
	DoDeleteValue    string             `config:"do_delete_value"`
	LoadBalance      bool               `config:"loadbalance"`
	TLS              *outputs.TLSConfig `config:"tls"`
	MaxRetries       int                `config:"max_retries"`
	Timeout          int                `config:"timeout"`
	SaveTopology     bool               `config:"save_topology"`
	Template         Template           `config:"template"`
}

type Template struct {
	Name      string `config:"name"`
	Path      string `config:"path"`
	Overwrite bool   `config:"overwrite"`
}

const (
	defaultBulkSize = 50
)

var (
	defaultConfig = elasticsearchConfig{
		Protocol:    "",
		Path:        "",
		ProxyURL:    "",
		Params:      nil,
		Username:    "",
		Password:    "",
		Timeout:     90,
		MaxRetries:  3,
		TLS:         nil,
		LoadBalance: true,
	}
)
