package main

import (
	"SpiderHog/Service/web"
	"SpiderHog/Spider"
	"SpiderHog/utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"net/url"
	"strings"
	"sync"
)

const host string = "https://code.fabao365.com"

type ParserKernel struct {

}

func (p ParserKernel) LinkExtract(response string) (urllist []url.URL) {

	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(response))
	selector := dom.Find("div.bs2l").Find("a[class=ch12],a[class=next]")
	selector.Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		href = strings.TrimSpace(href)
		if !strings.Contains(href, host) {
			href = host + href
		}
		Href, _ := url.Parse(href)
		urllist = append(urllist, *Href)
	})
	return
}

func (p ParserKernel) ContentExtract(response string) (Contents []interface{}) {
	//html, _ := utils.Response2String(response)
	selector := utils.ParseBySelector(response, ".bnr_s7")
	selector.Each(func(i int, selection *goquery.Selection) {
		content := strings.TrimSpace(selection.Text())
		Contents = append(Contents, content)
	})
	return
}

type SaverKernel struct {

}

func (s SaverKernel) Save(content interface{}) (error) {

	if str, ok := content.(string); ok {
		strs := utils.ParseByReg(str, "【法规名称】(.*?)\n")
		title := strs[1]
		content := str
		fmt.Printf("Title:%s\n", title)
		fmt.Printf("Content:%s\n", content)
		return nil
	} else {
		return errors.New("Type Assertion Error")
	}
}

func Run() {
	waitgroup := sync.WaitGroup{}
	pk := ParserKernel{}
	sk := SaverKernel{}
	Spider.Active(pk, sk)
	waitgroup.Add(1)
	go Spider.Crawl(&waitgroup)
	waitgroup.Add(1)
	go web.ServiceActive(&waitgroup)
	waitgroup.Wait()
}

