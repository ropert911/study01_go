package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//对字符串
func ioTest1() {
	{
		//ReadAll读取所有数据
		p, _ := ioutil.ReadAll(strings.NewReader("12345"))
		fmt.Println(string(p))
	}
	{
		r1 := strings.NewReader("aaa")
		//返回ReadCloser对象提供close函数
		rc1 := ioutil.NopCloser(r1)
		defer rc1.Close()
	}
}

func ioTest2() {
	//ReadDir返回目录下所有文件切片
	fileInfo, _ := ioutil.ReadDir("./")
	for _, data := range fileInfo {
		fmt.Println(data.Name())
	}
}

func ioTest3() {
	//读取整个文件数据
	data, _ := ioutil.ReadFile("./.gitignore")
	fmt.Println(string(data))
}
func main() {
	//bytesData := []byte("Hello World!")
	//reader = ioutil.NopCloser(bytes.NewReader(bytesData))

	//字符串Reader
	//ioTest1()

	//ReadDir返回目录下所有文件切片
	//ioTest2()

	//读取整个文件数据
	//ioTest3()

	//创建文件，存在清空文件
	ioutil.WriteFile("./1.txt", []byte("111"), 0655)

	//创建指定前缀的临时文件夹,返回文件夹名称
	dir, _ := ioutil.TempDir("./", "test")
	fmt.Println(dir)

	//创建test为前缀的临时文件，返回os.File指针
	f, _ := ioutil.TempFile("./", "test")
	f.Write([]byte("222"))
	f.Close()
}
