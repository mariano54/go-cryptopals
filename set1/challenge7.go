package main

import (
	"crypto/aes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ReadBase64DataFromFile(filename string) ([]byte, error) {
	absPath, err1 := filepath.Abs(filename)
	if err1 != nil {
		return nil, err1
	}
	datBase64, err2 := ioutil.ReadFile(absPath)
	if err2 != nil {
		return nil, err2
	}
	dat, err3 := Base64ToBytes(string(datBase64))
	if err3 != nil {
		return nil, err3
	}
	return dat, nil
}

func challenge7() {
	dat, err := ReadBase64DataFromFile("ciphertext-7.txt")
	HandleError(err)

	key := []byte("YELLOW SUBMARINE")

	cipher, err4 := aes.NewCipher(key)
	HandleError(err4)

	decrypted := make([]byte, len(dat))
	numBlocks := len(dat) / 16
	for i := 0; i < numBlocks; i++ {
		cipher.Decrypt(decrypted[i*16:], dat[i*16:(i+1)*16])
	}
	fmt.Println(string(decrypted))
}
