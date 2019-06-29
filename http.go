package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"unsafe"
)

type GpsInfo struct {
	DeviceId  string  `bson:"deviceid,omitempty" json:"deviceid"`
	Direction uint16  `bson:"direction,omitempty" json:"direction"`
	Lat       float32 `bson:"lat,omitempty" json:"lat"`
	Lng       float32 `bson:"lng,omitempty" json:"lng"`
	Speed     float32 `bson:"speed,omitempty" json:"speed"`
	Time      string  `bson:"time,omitempty" json:"time"`
	Vehicle   string  `bson:"vehicle,omitempty" json:"vehicle"`
}

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

func httpPostForm(urls string, v url.Values) {
	resp, err := http.PostForm(urls, v)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println(string(respBytes))
}

func HttpRequest(url string) (*string, error) {
	//提交请求
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	//处理返回结果
	client := &http.Client{}
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

func SamplePost1(urls string, v url.Values) {
	reader := strings.NewReader(v.Encode())
	req, err := http.NewRequest("POST", urls, ioutil.NopCloser(reader))
	if err != nil {
		// handle error
	}

	//Header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bear eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImEwMDIiLCJuYW1lIjoi5byg5LiJIiwiVGVsIjoiMTMzNjQ2NTg1ODUiLCJleHAiOjE1MjgwODQ2NDksImlzcyI6IueBq-WxseWPoyJ9.I2sDmSL17BnuDs8zi77ZBUAFxQYpouXIoKfZRfZLNRc")

	//发送
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	//回应
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(respBytes))
}

func SamplePost2(url string, bytesData []byte) {
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Header
	request.Header.Set("Content-Type", "application/json")

	//发送
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	//回应
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
}

func main() {
	HttpGet("http://www.baidu.com")
	content, _ := HttpRequest("http://www.baidu.com")
	fmt.Println(*content)

	httpPostForm("http://192.168.20.56:18004/itc/api/schoolbus/gps", url.Values{"Userlist": {"1145,1150"}})

	var gps = GpsInfo{DeviceId: "7070730", Direction: 120, Lat: 103.95077, Lng: 30.777979, Speed: 50.125, Time: "2019-06-19 17:48:40", Vehicle: "川A123"}
	bytesData, _ := json.Marshal(gps)
	SamplePost2("http://192.168.20.56:18004/itc/api/schoolbus/gps", bytesData)

	v := url.Values{}
	v.Set("ApiUserId", "7")
	v.Set("token", "99be71bc9c")
	v.Set("UserID", "916")
	v.Set("OrderID", "20201804110907523237")
	SamplePost1("http://192.168.20.56:18004/itc/api/schoolbus/gps", v)
}
