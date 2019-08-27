package spider

import (
	"../conf"
)

type Requester struct {
	Header conf.HttpHeader
	ProxyUrl string
	RequestMethod string
	Retry int
}
