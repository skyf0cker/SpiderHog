package utils

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"regexp"
	"strings"
)

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

func ReadJsonFile(path string) *json.Decoder{
	configFile, _ := os.Open(path)
	jsonParser := json.NewDecoder(configFile)
	return jsonParser
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