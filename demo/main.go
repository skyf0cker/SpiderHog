package Dome

import (
	"../conf"
	"../spider"
	"../utils"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"sync"
)

type myParser struct {
	spider.Parser
}

type mySaver struct {
	spider.Saver
}

func (p *myParser)LinkExtract(content string) []string {

	var resultList []string
	selection := utils.ParseBySelector(content, ".erjimenu")
	selection.Each(func(i int, s *goquery.Selection) {
		result := strings.TrimSpace(s.Text())
		//fmt.Println(result)
		resultList = append(resultList, result)
	})
	return resultList
}

func (p *myParser)MiddleLinkRule(content string) (urlList []string) {
	return urlList
}

func (s *mySaver)SaveRule(content string)  {
	//fmt.Println(content)
	s.StrSave(utils.Substr(content, 0, 2)+".txt", content)
}

func main() {
	wg := sync.WaitGroup{}
	f := spider.Fetcher{}
	p := myParser{}
	s := mySaver{}
	sp := spider.Spider{}
	j := conf.Configure()
	f.InitFetcher(j)
	p.InitParser(j)
	s.InitSaver(j)
	sp.InitSpider(j)

	wg.Add(1)
	go f.FetchContent(&wg)
	wg.Add(1)
	go spider.Parse(&wg, &p)
	wg.Add(1)
	go spider.Save(&wg, &s)
	wg.Wait()
	log.Println("[*]；  all jobs have been done!")
	//time.Sleep(50000000000000000)
}
