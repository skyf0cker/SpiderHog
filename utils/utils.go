package utils

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strings"
)

func Check(err error) bool {
	if err != nil{
		return false
	} else {
		return true
	}
}

func Exists(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func GetRandomUa () string {
	content, error := ioutil.ReadFile("./user-agent-list/ua.txt")
	if error != nil{
		log.Fatal("ua文件打开失败")
		panic("ua文件打开失败...")
	}
	str := string(content)
	ua_array := strings.Split(str, "\n")
	length := len(ua_array);

	if length<0{
		panic("ua文件没有数据")
	}

	index := rand.Intn(length)
	return ua_array[index-1]
}


func ParseByReg(html string, regPattern string) []string{
	r, error := regexp.Compile(regPattern)
	if error != nil{
		panic("正则表达式语法有误")
	} else {
		strList := r.FindStringSubmatch(html)
		return strList
	}
}
func ParseBySelector(html string, selector string) *goquery.Selection{

	dom, error := goquery.NewDocumentFromReader(strings.NewReader(html))
	if error != nil{
		log.Fatal("解析错误")
		return nil
	} else {
		test := dom.Find(selector)
		return test
	}
}

func Request(reqUrl string, meth string, proxy string, header map[string]string, retry int) *http.Response{
	var client http.Client
	if proxy != ""{
		pFunc, _ := url.Parse(proxy)
		tr := &http.Transport{
			Proxy: http.ProxyURL(pFunc),
		}
		client = http.Client{
			Transport: tr,
		}
	} else {
		log.Println("[*]:  没有使用代理")
		client = http.Client{}
	}
	req, _ := http.NewRequest(meth, reqUrl, nil)
	for key, val := range header{
		req.Header.Add(key, val)
	}
	times := 1
retry:
	response, err := client.Do(req)
	if !Check(err){
		if times < retry{
			times++
			log.Println("[*]:  正在重新尝试"+reqUrl)
			goto retry
		} else {
			log.Fatal("[*]:  请求"+reqUrl+"失败", err)
		}
	}
	return response
}

func ReadJsonFile(path string) *json.Decoder{
	configFile, _ := os.Open(path)
	jsonParser := json.NewDecoder(configFile)
	return jsonParser
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func Substr(str string, begin int, end int) string {
	return string([]rune(str)[begin:end])
}
