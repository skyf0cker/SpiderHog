package Spider

import (
	"log"
	"net/http"
	"net/url"
	"sync"
)

var UrlChan chan url.URL = make(chan url.URL, 20)
var ContentChan chan *http.Response = make(chan *http.Response, 20)
var SaveChan chan interface{} = make(chan interface{}, 20)

func Active(api ParserApi, saveApi SaveApi)  {
	ParserMachine.ActiveEngine(api)
	SaverMachine.ActiveEngine(saveApi)
}

func Crawl(group *sync.WaitGroup) {

	defer group.Done()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go FetcherMachine.Fetch(&wg)
	wg.Add(1)
	go ParserMachine.Parse(&wg)
	//wg.Add(1)
	//go SaverMachine.Saved(&wg)
	wg.Wait()
	log.Println("[*]；  all jobs have been done!")
}