package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func HttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(resp)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(resp)
	}
	fmt.Println(string(body))
}

func httpPostForm(urls string) {
	resp, err := http.PostForm(urls, url.Values{"Userlist": {"1145,1150"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println(string(body))
	//wg.Done()
}

func HttpRequest(url string) (*string, error) {
	//提交请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	//生成client 参数为默认
	client := &http.Client{}

	//处理返回结果
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("request error 1", err)
		return nil, err
	}

	defer response.Body.Close()

	body2, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("request error 2", err)
		return nil, err
	}

	var content = string(body2)
	return &content, nil
}

func httpDo(urls string) {
	v := url.Values{}
	v.Set("ApiUserId", "7")
	v.Set("token", "99be71bc9c")
	v.Set("UserID", "916")
	v.Set("OrderID", "20201804110907523237")
	body := ioutil.NopCloser(strings.NewReader(v.Encode()))

	client := &http.Client{}
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bear eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImEwMDIiLCJuYW1lIjoi5byg5LiJIiwiVGVsIjoiMTMzNjQ2NTg1ODUiLCJleHAiOjE1MjgwODQ2NDksImlzcyI6IueBq-WxseWPoyJ9.I2sDmSL17BnuDs8zi77ZBUAFxQYpouXIoKfZRfZLNRc")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body2, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body2))
}

func main() {
	//HttpGet("http://www.baidu.com")
	//httpPostForm("http://www.baidu.com")

	content, err := HttpRequest("http://www.baidu.com")
	if err == nil {
		fmt.Println(*content)
	}

}
