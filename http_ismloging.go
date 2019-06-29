package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpPostRequestBytesHeaderRest2(url string, bytesData []byte) (*string, error) {
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

	str, err := HttpPostRequestBytesHeaderRest2("http://192.168.20.56:20010/ias/auth/mobilelogin", bytesData)
	if nil == err {
		fmt.Println(*str)
	}

	var vMap map[string]string
	json.Unmarshal([]byte(*str), &vMap)
	fmt.Println("token is :", vMap["data"])
}
