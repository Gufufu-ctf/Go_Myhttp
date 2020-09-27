package Myhttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Myhttp struct{
	Url string
	Header map[string]string
	Post_value map[string]string
	IsProxy bool
	Timeout int
}

var proxyConf = "127.0.0.1:1080"


func NewMyhttp()* Myhttp{
	return &Myhttp{
		Url : "",
		Header: map[string]string{},
		Post_value: map[string]string{},
		IsProxy: false,
		Timeout: 10,
	}
}



func buildHttpClient(mytimeout int,isProxy bool) *http.Client {			//是否使用代理
	var proxy func(*http.Request) (*url.URL, error) = nil
	if isProxy {
		proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse("socks5://" + proxyConf)
		}
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport, Timeout: time.Duration(mytimeout) * time.Second  }		//设置 超时时间，一般用于时间盲注
	return client
}


func (this*Myhttp)Get()string{			//post_value = "username=amdin&passwd=admin"
	client :=buildHttpClient(this.Timeout,this.IsProxy)
	req, err := http.NewRequest("GET", this.Url, strings.NewReader(""))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	for i,v := range this.Header{
		req.Header.Set(i,v)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}


func (this*Myhttp)Post()string{			//post_value = "username=amdin&passwd=admin"
	client :=buildHttpClient(this.Timeout,this.IsProxy)
	req, err := http.NewRequest("POST", this.Url, strings.NewReader(this.Poststr()))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for i,v := range this.Header{
		req.Header.Set(i,v)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(body)
}



func (this*Myhttp)Poststr()string{
	res := ""
	for i,v := range  this.Post_value{
		res += i+"="+url.QueryEscape(v)+"&"
	}
	res = res[:len(res)-1]
	return res
}