package crypt

// golang AES-CBC with IV encrypt & decrypt
// author: LIHKG Foresight

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func Encrypt(key, data []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	dataWithPadding := PKCS7Padding(data)
	buffer := make([]byte, aes.BlockSize+len(dataWithPadding))

	iv := buffer[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	encrypter := cipher.NewCBCEncrypter(block, iv)
	encrypter.CryptBlocks(buffer[len(iv):], dataWithPadding)

	return buffer
}

func Decrypt(key, data []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	iv := data[0:aes.BlockSize]

	buffer := make([]byte, len(data)-len(iv))
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypter.CryptBlocks(buffer, data[len(iv):])

	return PKCS7Unpadding(buffer)
}

func PKCS7Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7Unpadding(plantText []byte) []byte {
	length := len(plantText)
	padding := int(plantText[length-1])
	return plantText[:(length - padding)]
}
