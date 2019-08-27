package spider

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"../utils"
	)

type Proxier struct {
	Ip string
	Port string
}

type ip struct {
	ip string
	port int
}

func (p *Proxier)GetProxy() string {
	client := http.Client{}
	request, _ := http.NewRequest("GET", p.Ip + p.Port, nil)
	response, _ := client.Do(request)
	var i ip
	err := json.NewDecoder(response.Body).Decode(&i)
	if utils.Check(err){
		log.Fatal("[*]:  json解析失败")
	}
	return i.ip+string(i.port)
}

func (p *Proxier)CheckProxyAlive(proxy string) bool {

	pUrl, _ := url.Parse(proxy)
	tr := &http.Transport{
		Proxy:                  http.ProxyURL(pUrl),
	}
	
	client := &http.Client{
		Transport: tr,
	}
	request,_ := http.NewRequest("get", "http://www.sxu.edu.cn", nil)

	_, err := client.Do(request)

	if utils.Check(err){
		return true
	} else {
		return false
	}
}

