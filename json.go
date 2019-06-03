package main

import (
	"encoding/json"
	"fmt"
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

func test1() {
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
func test2() {
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

func test3() {
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
func test4() {
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
func main() {
	test1()
	test2()
	test3()
	test4()
}
