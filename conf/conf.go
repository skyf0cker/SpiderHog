package conf

import (
	"../utils"
	"log"
)

const (
	configPath string = "./settings.json"
)

type JsonConfig struct {
	FConfig FetcherConfig
	PConfig ParserConfig
	SConfig SaverConfig
	SpConfig SpiderConfig
	PrConfig ProxierConfig
}

type SpiderConfig struct {
	Header HttpHeader
	TargetUrl string
	Method string
	Retry int
}

type FetcherConfig struct {
}

type ParserConfig struct {
}

type SaverConfig struct {
	SavePath string
	SaveMethod string
}

type ProxierConfig struct {
	Ip   string
	Port string
}

type HttpHeader struct {
	UserAgent string `json:"User-Agent"`
	Referer string
	Host string
	Connection string
	Pragma string
	CacheControl string `json:"Cache-Control"`
	SecFetchMode string `json:"Sec-Fetch-Mode"`
	Accept string
	SecFetchSite string `json:"Sec-Fetch-Site"`
	Cookies string
}

func Configure() JsonConfig {
	j := JsonConfig{}
	if ! utils.Exists(configPath) {
		log.Fatal("[*]:  Missing the config file:settings.json")
	} else {
		jsonParser := utils.ReadJsonFile(configPath)
		err := jsonParser.Decode(&j)
		if !utils.Check(err) {
			log.Fatal("[*]:  json parse failed",err)
		}
	}
	return j
}



