package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

type Animal struct {
	Name  string
	Order string
}

type Message struct {
	Name  string
	Body  string
	Time  int64
	inner string
}

func jsonTest1() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	//对像转json格式
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("编码后的json:")
	os.Stdout.Write(b)
	fmt.Println("")

}
func jsonTest2() {
	//json传对象
	var jsonBlob = []byte(`[
    {"Name": "Platypus", "Order": "Monotremata"},
    {"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)

	var animals []Animal
	err2 := json.Unmarshal(jsonBlob, &animals)
	if err2 != nil {
		fmt.Println("error:", err2)
	}
	fmt.Printf("解析后的对象：%+v \n", animals)
}

func jsonTest3() {
	//json对大小写不敏感，但对象里的必须是大写开头
	var m = Message{
		Name:  "Alice",
		Body:  "Hello",
		Time:  1294706395881547000,
		inner: "ok",
	}
	b3 := []byte(`{"nAmE":"Bob","Food":"Pickle", "inner":"changed"}`)

	err3 := json.Unmarshal(b3, &m)
	if err3 != nil {
		fmt.Printf(err3.Error())
		return
	}
	fmt.Printf("%v\n", m)
}

//StructTag  如果希望手动配置结构体的成员和JSON字段的对应关系，可以在定义结构体的时候给成员打标签：
//	使用omitempty熟悉，如果该字段为nil或0值（数字0,字符串"",空数组[]等），则打包的JSON结果不会有这个字段
func jsonTest4() {
	fmt.Println("jsonTest4===========")
	type Message struct {
		Name string `json:"msg_name"`       // 对应JSON的msg_name
		Body string `json:"body,omitempty"` // 如果为空置则忽略字段
		Time int64  `json:"-"`              // 直接忽略字段
	}
	var m = Message{
		Name: "Alice",
		Body: "",
		Time: 1294706395881547000,
	}
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Println(string(data))
}

func jsonTest5() {
	fmt.Println("jsonTest5===========")
	b3 := []byte(`{"name":"xq","cmd":"ping a host"}`)
	var response map[string]interface{}
	json.Unmarshal(b3, &response)

	deviceName := response["name"].(string)
	fmt.Println(deviceName, response["cmd"].(string))
}

type GwData struct {
	Channel  int     `json:"channel"`
	Sf       int     `json:"sf"`
	Time     string  `json:"time"`
	Gwip     string  `json:"gwip"`
	Gwid     string  `json:"gwid"`
	Repeater string  `json:"repeater"`
	Systype  int     `json:"systype"`
	Rssi     float32 `json:"rssi"`
	Snr      float32 `json:"snr"`
	Snr_max  float32 `json:"snr_max"`
	Snr_min  float32 `json:"snr_min"`
	MacAddr  string  `json:"macAddr"`
	Data     string  `json:"data"`
	FrameCnt int     `json:"frameCnt"`
	Fport    int     `json:"fport"`
}

func jsonTest6() {
	fmt.Println("jsonTest6===========")
	data := GwData{
		Channel:  488500000,
		Sf:       12,
		Time:     "2019-06-06T15:58:11+08:00",
		Gwip:     "10.10.21.84",
		Gwid:     "000080029c09e987",
		Repeater: "00000000ffffffff",
		Systype:  3,
		Rssi:     -27.0,
		Snr:      21.3,
		Snr_max:  37.0,
		Snr_min:  9.5,
		MacAddr:  "000000000301095d",
		Data:     "031e19f931fea80408",
		FrameCnt: 4,
		Fport:    5,
	}
	dataStr, err := json.Marshal(data)
	if err != nil {
		fmt.Println("ERROR", err.Error())
		return
	}
	fmt.Println(string(dataStr))
}

type User struct {
	Id      string
	Balance uint64
}

func main() {
	//jsonTest1() //json.Marshal(group) 对象=>json格式
	//jsonTest2() //json.Unmarshal  json=>对象
	//jsonTest3()
	//jsonTest4()
	//jsonTest5()
	//jsonTest6()

	//对象=>json
	u := User{Id: "www.361way.com", Balance: 8}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	io.Copy(os.Stdout, b)

	//json==>对象
	var u2 User
	json.NewDecoder(b).Decode(&u2)
	fmt.Println(u2)
}
