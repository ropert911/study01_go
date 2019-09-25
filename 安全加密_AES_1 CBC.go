package main

//安全加密的参考文章：https://studygolang.com/articles/15642?fr=sidebar
//对称加密, 加解密都使用的是同一个密钥, 其中的代表就是AES

//AES：高级加密标准（Advanced Encryption Standard），又称Rijndael加密法，这个标准用来替代原先的DES。AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
//块：对明文进行加密的时候，先要将明文按照128bit进行划分。
//填充方式：因为明文的长度不一定总是128的整数倍，所以要进行补位，我们这里采用的是PKCS7填充方式
//AES实现的方式多样, 其中包括ECB、CBC、CFB、OFB等

//加密模式	对应加解密方法
//CBC		NewCBCDecrypter, NewCBCEncrypter
//CTR		NewCTR
//CFB		NewCFBDecrypter, NewCFBEncrypter
//OFB		NewOFB

//密码分组链接模式（Cipher Block Chaining (CBC)）		最常见的使用的方式
//  将明文分组与前一个密文分组进行XOR运算，然后再进行加密。每个分组的加解密都依赖于前一个分组。而第一个分组没有前一个分组，因此需要一个初始化向量

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

//补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//解密
func AesDecrypt(cryted string, key string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize]) //这里第二个参数是向量，这里没有批量，直接用了一个数
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

//加密
func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32. 即 128、192、256bi
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize() //这里得到的是16字节-128位
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize]) //这里第二个参数是向量，这里没有批量，直接用了一个数
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}

func main() {
	orig := "hello world hello world"
	key := "0123456789012345"
	fmt.Println("原文：", orig)
	encryptCode := AesEncrypt(orig, key)
	fmt.Println("密文：", encryptCode)
	decryptCode := AesDecrypt(encryptCode, key)
	fmt.Println("解密结果：", decryptCode)
}
