package Spider

import (
	"SpiderHog/utils"
	"net/http"
	"net/url"
)

type RequesterApi interface {
	SetHeader(http.Header)
	SetCookies(utils.Cook1es)
	SetUrl(url.URL)
	SetData(utils.Data)
	SetProxy(url.URL)
	SetProxySrc(url.URL)
	SetMethod(string)
	GetProxy()
	Get() (*http.Response, error)
	Post() (*http.Response, error)
	CheckProxy(u url.URL) (error)
}
