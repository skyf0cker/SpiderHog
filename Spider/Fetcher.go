package Spider

import (
	"SpiderHog/conf"
	"SpiderHog/utils"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var FetcherMachine *Fetcher
//var FLog *logrus.Logger = utils.GetStdoutLogger()
var FLog *logrus.Logger = utils.GetFileLogger()

type Fetcher struct {
	Header   http.Header
	Cookies  http.CookieJar
	Tarurl   url.URL
	RawData  utils.Data
	Proxy    url.URL
	ProxySrc url.URL
	Method   string
}

//{"proxy": "113.65.5.6:8118", "fail_count": 0, "region": "", "type": "", "source": "freeProxy03", "check_count": 1, "last_status": 1, "last_time": "2019-10-08 01:12:59"}
type ProxyStruct struct {
	Proxy      string `json:"proxy"`
	FailCount  int    `json:"fail_count"`
	Region     string `json:"region"`
	Type       string `json:"type"`
	Source     string `json:"source"`
	CheckCount int    `json:"check_count"`
	LastStatus int    `json:"last_status"`
	LastTime   string `json:"last_time"`
}

//init 里面要做的就是加载配置文件，初始化变量

func init() {
	Config := conf.GetConfigure()
	fet := Config.FetConfig
	header := fet.Header
	cookies := utils.Cook1es{}
	for _, cookie := range fet.Cookies {
		cookies = append(cookies, &http.Cookie{
			Name:  cookie["Name"],
			Value: cookie["Value"],
		})
	}
	tarurl, _ := url.Parse(fet.BeginUrl)
	data := fet.RawData
	proxysrc, _ := url.Parse(fet.ProxySource)
	method := fet.Method

	FetcherMachine = &Fetcher{
		Header:   header,
		Cookies:  cookies,
		Tarurl:   *tarurl,
		RawData:  data,
		ProxySrc: *proxysrc,
		Method:   method,
	}
	FetcherMachine.GetProxy()
	UrlChan <- FetcherMachine.Tarurl
}

func (f *Fetcher) SetHeader(h http.Header) {
	f.Header = h
}

func (f *Fetcher) SetCookies(c utils.Cook1es) {
	f.Cookies = c
}

func (f *Fetcher) SetUrl(url url.URL) {
	f.Tarurl = url
}

func (f *Fetcher) SetData(d utils.Data) {
	f.RawData = d
}

func (f *Fetcher) SetProxy(p url.URL) {
	f.Proxy = p
}

func (f *Fetcher) SetProxySrc(u url.URL) {
	f.ProxySrc = u
}

func (f *Fetcher) SetMethod(m string) {
	f.Method = m
}

func (f *Fetcher) GetProxy() {
	if f.ProxySrc.String() == "" {
		return
	}

	hosturl := f.ProxySrc.String()
	geturl := hosturl + "/get"

	for {
		if resp, err := http.Get(geturl); err != nil {
			panic("proxy pool wrong")
		} else {
			var p ProxyStruct
			if jerr := json.NewDecoder(resp.Body).Decode(&p); jerr != nil {
				panic(jerr)
			} else {
				rawurl := p.Proxy
				surl := "http://" + rawurl
				if uurl, perr := url.Parse(surl); perr != nil {
					return
				} else {
					if err2 := f.CheckProxy(*uurl); err2 != nil {
						//这里应该想代理池发出请求，删除不可用代理
						delurl := hosturl + "/delete?proxy=" + rawurl
						_, err3 := http.Get(delurl)
						if err3 != nil {
							FLog.Println("proxy deleted failed!")
						} else {
							FLog.Println("delete the url " + delurl)
						}
						continue
					} else {
						FLog.Printf("using proxy:%s", (*uurl).String())
						f.SetProxy(*uurl)
						break
					}
				}
			}
		}
	}
}

func (f *Fetcher) Get() (rp *http.Response, e error) {
	rp, e = utils.ReqGet(f.Tarurl, f.Proxy, f.Header, f.Cookies)
	return
}

func (f *Fetcher) Post() (rp *http.Response, e error) {
	rp, e = utils.ReqPost(f.Tarurl, f.Proxy, f.Header, f.Cookies, f.RawData)
	return
}

func (f *Fetcher) CheckProxy(u url.URL) (err error) {
	if testUrl, errt := url.Parse("http://www.baidu.com"); errt != nil {
		err = errt
		return
	} else {
		if _, errt2 := utils.ReqGet(*testUrl, u, f.Header, f.Cookies); errt2 != nil {
			err = errors.New("proxy can not be used!")
			return
		}
	}
	return
}

func (f *Fetcher) Fetch(group *sync.WaitGroup) {
	defer group.Done()
	for {
		select {
		case url := <-UrlChan:
			FLog.Println("[*]:  begin fetching " + url.String())
			f.Tarurl = url
			if f.Method == "GET" {
				if response, err := f.Get(); err == nil {
					ContentChan <- response
					//newResponse := <-ContentChan
					//str, _ := utils.Response2String(*newResponse)
					//fmt.Println(str)
				} else {
					FLog.Println("request failed....")
					FLog.Println("Trying to change the proxy...")
					f.GetProxy()
					UrlChan<-f.Tarurl
				}
			} else {
				if response, err := f.Post(); err == nil {
					ContentChan <- response
				}
			}
		case <-time.After(time.Duration(60) * time.Second):
			FLog.Println("[*]:  timeout, fetcher exit...")
			return
		}
	}
}
