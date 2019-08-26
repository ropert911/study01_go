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

//输出反馈模式（Output FeedBack (OFB)）
