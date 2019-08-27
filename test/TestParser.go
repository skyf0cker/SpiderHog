//package Test
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"../spider"
)

func ParserByRegTest() {
	url := "http://www.sxu.edu.cn"
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	response, _ := client.Do(req)
	bContent, _ := ioutil.ReadAll(response.Body)
	content := string(bContent)
	fmt.Println(content)
	//spider.ParseByReg(content, "<a href=.*?>")
}

func ParseBySelectorTest() {
	url := "http://www.sxu.edu.cn"
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	response, _ := client.Do(req)
	bContent, _ := ioutil.ReadAll(response.Body)
	content := string(bContent)
	spider.ParseBySelector(content, ".indextitlemore")
}


func main () {
	ParseBySelectorTest()
}
