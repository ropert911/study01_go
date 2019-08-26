package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

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

//4.密码反馈模式（Cipher FeedBack (CFB)）
//  前一个密文分组会被送回到密码算法的输入端。

func ExampleNewCFBDecrypter(cryted1 string, key1 string) []byte {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	key, _ := hex.DecodeString(key1)
	ciphertext, _ := hex.DecodeString(cryted1)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	//fmt.Printf("%s", ciphertext)
	// Output: some plaintext
	return ciphertext
}

func ExampleNewCFBEncrypter(content1 string, key1 string) []byte {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	key, _ := hex.DecodeString(key1)
	plaintext := []byte(content1)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	//fmt.Printf("%x\n", ciphertext)
	return ciphertext
}

func main() {
	orig := "some plaintext"
	fmt.Printf("加密前：%s\n", orig)

	encryptCode := ExampleNewCFBEncrypter("some plaintext", "6368616e676520746869732070617373")
	fmt.Printf("加密后的结果：%x\n", encryptCode)
	decryptCode := ExampleNewCFBDecrypter(fmt.Sprintf("%x", encryptCode), "6368616e676520746869732070617373")
	fmt.Println("解密结果：", string(decryptCode))
}
