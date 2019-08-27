package spider

import (
	"../conf"
)

var urlChan chan string = make(chan string, 20)
var contentChan chan string  = make(chan string, 20)
var saveChan chan string = make(chan string, 20)

type Spider struct {
	BeginUrl string
	Header conf.HttpHeader
	Proxy string
}

func (s *Spider)InitSpider(j conf.JsonConfig) {
	s.BeginUrl = j.SpConfig.TargetUrl
	s.Header = j.SpConfig.Header
	s.Proxy = j.PrConfig.Ip+j.PrConfig.Port
	urlChan<-s.BeginUrl
}

//func (sp *Spider)Work(f Fetcher, p Parser, s Saver) {
//	go f.FetchContent()
//	go p.Parse()
//	go s.Save()
//}
// 这里可以有一个类叫做conf，把所有的配置信息导入
