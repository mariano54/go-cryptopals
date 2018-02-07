package main

import (
	"crypto/aes"
	"fmt"
	"go-cryptopals/set1"
	"log"
)

const BlockSize = 16

func PKCS7Padding(data []byte) []byte {
	if len(data)%BlockSize != 0 {
		bytesMissing := BlockSize - len(data)%BlockSize
		for i := 0; i < bytesMissing; i++ {
			data = append(data, byte(bytesMissing))
		}
	}
	return data
}

func CBCEncrypt(iv []byte, key []byte, plaintext []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		log.Panic(err)
	}
	plaintext = PKCS7Padding(plaintext)
	ciphertext := make([]byte, len(plaintext))
	numBlocks := (len(ciphertext) / BlockSize)

	prev := iv
	for i := 0; i < numBlocks; i++ {
		start, end := i*BlockSize, (i+1)*BlockSize
		xored, err2 := set1.Xor(plaintext[start:end], prev)
		if err2 != nil {
			panic("Failed xor")
		}
		cipher.Encrypt(ciphertext[start:end], xored)
		prev = ciphertext[start:end]
	}
	return ciphertext
}

func CBCDecrypt(iv []byte, key []byte, ciphertext []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		log.Panic(err)
	}
	ciphertext = PKCS7Padding(ciphertext)
	plaintext := make([]byte, len(ciphertext))
	numBlocks := (len(ciphertext) / BlockSize)
	prev := iv
	for i := 0; i < numBlocks; i++ {
		start, end := i*BlockSize, (i+1)*BlockSize
		cipher.Decrypt(plaintext[start:end], ciphertext[start:end])
		xored, err2 := set1.Xor(plaintext[start:end], prev)
		if err2 != nil {
			panic("Failed xor")
		}
		copy(plaintext[start:end], xored)
		prev = ciphertext[start:end]
	}
	return plaintext
}

func challenge10() {
	key := []byte("YELLOW SUBMARINE")

	data, err := set1.ReadBase64DataFromFile("ciphertext-10.txt")
	if err != nil {
		log.Panic(err)
	}
	iv := make([]byte, 16)
	decrypted := CBCDecrypt(iv, key, data)
	encrypted := CBCEncrypt(iv, key, decrypted)
	decrypted = CBCDecrypt(iv, key, encrypted)

	fmt.Println(string(decrypted))
}
