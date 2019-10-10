package utils

import (
	"crypto/tls"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Cook1es []*http.Cookie

func (c Cook1es) SetCookies(u *url.URL, cookies []*http.Cookie) {
	c = cookies
}

func (c Cook1es) Cookies(u *url.URL) []*http.Cookie {
	return c
}

type Data map[string]string

func ReqGet(tarurl url.URL, proxy url.URL, header http.Header, cookies http.CookieJar) (response *http.Response, err error) {
	var client *http.Client
	if tarurl.String() == "" {
		err = errors.New("url not found!")
		return
	} else {
		trans := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		if proxy.String() != "" {
			trans.Proxy = func(request *http.Request) (Url *url.URL, e error) {
				Url = &proxy
				return
			}
		}

		client = &http.Client{
			Timeout:15*time.Second,
			Transport: trans,
		}

		if cookies != nil {
			client.Jar = cookies
		}

		if req, temp := http.NewRequest("GET", tarurl.String(), nil); temp != nil {
			err = temp
			return
		} else {
			req.Header = header
			var temp2 error = nil
			if response, temp2 = client.Do(req); temp2 != nil {
				err = temp2
				return
			} else {
				return
			}
		}
	}
	return
}

func ReqPost(tarurl url.URL, proxy url.URL, header http.Header, cookies http.CookieJar, data map[string]string) (response *http.Response, err error) {
	var client *http.Client
	if tarurl.String() == "" {
		err = errors.New("url not found!")
		return
	} else {
		if len(data) != 0 {
			urlData := url.Values{}
			for k, v := range data {
				urlData.Set(k, v)
			}
			trans := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			if proxy.String() != "" {
				trans.Proxy = func(request *http.Request) (Url *url.URL, e error) {
					Url = &proxy
					return
				}
			}
			client = &http.Client{
				Transport: trans,
			}

			if cookies != nil {
				client.Jar = cookies
			}

			body := ioutil.NopCloser(strings.NewReader(urlData.Encode())) //把form数据编下码

			if req, temp := http.NewRequest("POST", tarurl.String(), body); temp != nil {
				err = temp
				return
			} else {
				req.Header = header
				req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
				var temp2 error = nil
				if response, temp2 = client.Do(req); temp2 != nil {
					err = temp2
					return
				} else {
					return
				}
			}
		}
	}
	return
}

func Response2String(response http.Response) (string, error) {
	str, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return string(str),err
	}
	return string(str), err
}
