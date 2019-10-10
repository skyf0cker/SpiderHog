package Spider

import (
	"SpiderHog/conf"
	"SpiderHog/utils"
	"github.com/sirupsen/logrus"
	"net/url"
	"sync"
	"time"
)

var ParserMachine *Parser
//var PLog *logrus.Logger = utils.GetStdoutLogger()
var PLog *logrus.Logger = utils.GetFileLogger()


type Parser struct {
	ParseEngine ParserApi
}

func init() {
	Config := conf.GetConfigure()
	_ = Config.ParConfig // now is useless
	ParserMachine = &Parser{}
}

type ParserApi interface {
	//LinkExtract(response http.Response) (urllist []url.URL)
	//ContentExtract(response http.Response) (Contents []interface{})
	LinkExtract(response string) (urllist []url.URL)
	ContentExtract(response string) (Contents []interface{})
}

func (p *Parser)ActiveEngine(ParseEngine ParserApi)  {
	p.ParseEngine = ParseEngine
}

func (p *Parser)Parse(group *sync.WaitGroup) {
	defer group.Done()
	for {
		select {
		case content := <-ContentChan:
			PLog.Println("[*]:  begin parsing content...")
			html, _ := utils.Response2String(*content)
			go func() {
				for _, i := range p.ParseEngine.ContentExtract(html){
					SaveChan<-i
				}
			}()

			go func() {
				for _, j := range p.ParseEngine.LinkExtract(html){
					UrlChan<-j
				}
			}()
		case <-time.After(time.Duration(80)*time.Second):
			PLog.Println("[*]:  timeout, parser exit...")
			return
		}
	}
}


