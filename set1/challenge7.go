package set1

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

func ECBDecrypt(key []byte, ciphertext []byte) []byte {
	cipher, err := aes.NewCipher(key)
	HandleError(err)

	plaintext := make([]byte, len(ciphertext))
	numBlocks := len(ciphertext) / 16
	for i := 0; i < numBlocks; i++ {
		cipher.Decrypt(plaintext[i*16:], ciphertext[i*16:(i+1)*16])
	}
	return plaintext
}

func ECBEncrypt(key []byte, plaintext []byte) []byte {
	cipher, err := aes.NewCipher(key)
	HandleError(err)

	ciphertext := make([]byte, len(plaintext))
	numBlocks := len(plaintext) / 16
	for i := 0; i < numBlocks; i++ {
		cipher.Encrypt(ciphertext[i*16:], plaintext[i*16:(i+1)*16])
	}
	return ciphertext
}
func challenge7() {
	data, err := ReadBase64DataFromFile("ciphertext-7.txt")
	HandleError(err)
	key := []byte("YELLOW SUBMARINE")

	fmt.Println(string(ECBDecrypt(key, data)))
}
