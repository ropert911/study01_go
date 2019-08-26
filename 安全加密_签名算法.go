package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

//签名算法, 如MD5、SHA1、HMAC等, 主要用于验证，防止信息被修改, 如：文件校验、数字签名、鉴权协议

// sha256加密文件内容
func fileSha156(fpath string) {
	file, err := os.OpenFile(fpath, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	h := sha256.New()
	// 将文件内容拷贝到sha256中
	io.Copy(h, file)
	fmt.Printf("%x\n", h.Sum(nil))
}

func main() {
	// sha256加密字符串   Secure Hash Algorithm，缩写为SHA 安全散列算法
	str := "hello world"
	sum := sha256.Sum256([]byte(str))
	fmt.Printf("SHA256：%x\n", sum)
	sum1 := sha256.Sum256([]byte(str))
	fmt.Printf("SHA256：%x\n", sum1)

	// md5加密
	result := md5.Sum([]byte(str))
	fmt.Printf("MD5：%x\n", result)
	result1 := md5.Sum([]byte(str))
	fmt.Printf("MD5：%x\n", result1)

	fileSha156("D:\\source_code\\go_work\\src\\study01_go\\安全加密_签名算法.go")
}
