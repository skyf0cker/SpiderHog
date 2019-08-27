package spider

import (
	"../conf"
	"../utils"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

type Fetcher struct {
	Requester
}

func (f *Fetcher)InitFetcher(j conf.JsonConfig)  {
	f.Retry = j.SpConfig.Retry
	f.Header = j.SpConfig.Header
	f.RequestMethod = j.SpConfig.Method
	//f.ProxyUrl = j.PrConfig.Ip + j.PrConfig.Port
	f.ProxyUrl = ""
}

func (f *Fetcher) FetchContent (group *sync.WaitGroup) {
	defer group.Done()
	for {
		select {
		case url := <-urlChan:
			log.Println("[*]:  begin fetching " + url)
			retryTime := f.Retry
			proxyUrl := f.ProxyUrl
			header := utils.Struct2Map(f.Header)
			Strheader := make(map[string]string)
			for key, val := range header{
				if val != "" {
					Strheader[key] = val.(string)
				}
			}
			response := utils.Request(url, "GET", proxyUrl, Strheader, retryTime)
			s, _ := ioutil.ReadAll(response.Body)
			contentChan<-string(s)
		case <-time.After(time.Duration(5)*time.Second):
			log.Println("[*]:  timeout, fetcher exit...")
			return
		}
	}
}

