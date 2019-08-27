package spider

import (
	"../conf"
	"log"
	"sync"
	"time"
)

type Parser struct {
	Requester
}

type ParseInterface interface {
	LinkExtract(content string) []string
	MiddleLinkRule(content string) []string
}

func (p *Parser)InitParser(j conf.JsonConfig)  {
	p.Retry = j.SpConfig.Retry
	p.Header = j.SpConfig.Header
	p.ProxyUrl = j.PrConfig.Ip + j.PrConfig.Port
	p.RequestMethod = j.SpConfig.Method
}

func (p *Parser)LinkExtract(content string) (urlList []string) {
	return urlList
}

func (p *Parser)MiddleLinkRule(content string) (urlList []string) {
	return urlList
}

func Parse(group *sync.WaitGroup, parseInterface ParseInterface) {
	defer group.Done()
	for {
		select {
		case content := <-contentChan:
			log.Println("[*]:  begin parsing content...")
			go func() {
				for _, i := range parseInterface.LinkExtract(content){
					saveChan<-i
				}
			}()

			go func() {
				for _, j := range parseInterface.MiddleLinkRule(content){
					urlChan<-j
				}
			}()
		case <-time.After(time.Duration(5)*time.Second):
			log.Println("[*]:  timeout, parser exit...")
			return
		}
	}
}