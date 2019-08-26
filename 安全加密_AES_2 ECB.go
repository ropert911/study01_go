package main

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

//1.电码本模式（Electronic Codebook Book (ECB)）
//   将明文分组加密之后的结果直接称为密文分组。

//ECB模式: mysql中AES_DECRYPT函数的实现方式
//主要关注三点:
//1.调用aes.NewCipher([]byte)是加密关键字key的生成方式, 即下面的generateKey方法
//2.分组分块加密的加密方式
//3.mysql中一般需要HEX函数来转化数据格式
//加密: HEX(AES_ENCRYPT('关键信息', '***—key'))
//解密: AES_DECRYPT(UNHEX('关键信息'), '***-key’)
//所以调用AESEncrypt或者AESDecrypt方法之后, 使用hex.EncodeToString()转化
import (
	"crypto/aes"
	"fmt"
)

func AESEncrypt(src []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

func AESDecrypt(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func main() {
	orig := "hello world hello world"
	key := "0123456789012345"
	fmt.Println("原文：", orig)
	encryptCode := AESEncrypt([]byte(orig), []byte(key))
	fmt.Println("密文：", string(encryptCode))
	decryptCode := AESDecrypt(encryptCode, []byte(key))
	fmt.Println("解密结果：", string(decryptCode))
}
