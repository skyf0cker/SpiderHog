package conf

import (
	"SpiderHog/utils"
	"net/http"
	"sync"
)

var config *Configure
var once sync.Once

const (
	configPath string = "./conf/settings.json"
)

type RequesterConfig struct {
	Header      http.Header   `json:"header"`
	Cookies     []map[string]string `json:"cookies"`
	BeginUrl    string        `json:"begin_url"`
	RawData     utils.Data    `json:"raw_data"`
	Method      string        `json:"method"`
	ProxySource string        `json:"proxy_source"`
}

type Configure struct {
	FetConfig FetcherConfig `json:"Fetcher"`
	ParConfig ParserConfig	`json:"Parser"`
	SavConfig SaverConfig	`json:"Saver"`
}

//下面是各自的配置文件

type FetcherConfig struct {
	RequesterConfig
}

type ParserConfig struct {
}

type SaverConfig struct {
}

func GetConfigure() *Configure {
	once.Do(func() {
		if ! utils.Exists(configPath) {
			panic("Missing the config file:settings.json")
		}
		config = &Configure{}
		jsonParser := utils.ReadJsonFile(configPath)
		if err := jsonParser.Decode(config); err != nil {
			panic("configure file parsed failed!")
		}
	})
	return config
}
