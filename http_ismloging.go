package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func RESTHttpPostRequestBytes(url string, urlArg string, bytesData []byte) (*string, error) {
	if len(urlArg) > 0 {
		url = url + "?" + urlArg
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	//Header
	request.Header.Set("Content-Type", "application/json")

	//发送
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	//回应
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var content = string(respBytes)
	return &content, nil
}

func RESTHttpGetRequest(url string, url_arg string) (*string, error) {
	url = url + "?" + url_arg
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	//Header
	request.Header.Set("Content-Type", "application/json")

	//发送
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	//回应
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var content = string(respBytes)
	return &content, nil
}

type LogInfo struct {
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	IsRemember string `json:"isRemember"`
}

func main() {
	//方法一：要构造对象
	//var lInfo = LogInfo{UserName: "admin", Password: "admins", IsRemember: "true"}
	//bytesData, _ := json.Marshal(lInfo)

	//方法二：用map进行构建
	values := map[string]string{"UserName": "admin", "Password": "admins", "IsRemember": "true"}
	bytesData, _ := json.Marshal(values)

	//`{"UserName":"admin","Password":"admins", "IsRemember":"true"}`   这是对象的字符串

	str, err := RESTHttpPostRequestBytes("http://192.168.20.56:20010/ias/auth/mobilelogin", "", bytesData)
	if nil == err {
		fmt.Println(*str)
	}

	var vMap map[string]interface{}
	json.Unmarshal([]byte(*str), &vMap)
	fmt.Println("success is :", vMap["success"].(bool))
	fmt.Println("token is :", vMap["data"].(string))

	//建一级区域
	values = map[string]string{
		"address":        "in the hell",
		"businessArea":   "1",
		"businessTypeID": "1",
		"categoryID":     "1",
		"cityID":         "1",
		"contact":        "auto_tester",
		"email":          "xxx@163.com",
		"name":           "xxxxxxxxxxxxxxxxxxxx",
		"passengers":     "1",
		"phone":          "10010",
		"provinceID":     "1",
		"remark":         "comm",
		"scene":          "IOT",
	}
	bytesData, _ = json.Marshal(values)
	str, err = RESTHttpPostRequestBytes("http://192.168.20.56:20010/ias/top/add", "token="+vMap["data"].(string), bytesData)
	if nil == err {
		fmt.Println(*str)
	}

	//
	str, err = RESTHttpGetRequest("http://192.168.20.56:20010/ias/auth/logout", "tokens=%s"+vMap["data"].(string))
	if nil == err {
		fmt.Println(*str)
	}
}
