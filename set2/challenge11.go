package main

import (
	"fmt"
	"go-cryptopals/set1"
	"math/rand"
)

func randomBytes(numBytes int) []byte {
	buf := make([]byte, numBytes)
	read, err := rand.Read(buf)
	if read != numBytes || err != nil {
		panic("Failed to generate random number")
	}
	return buf
}

func generateKey() []byte {
	return randomBytes(BlockSize)
}

func randomEncrypt(plaintext []byte) []byte {
	plaintext = append(randomBytes(rand.Int()%5), plaintext...)
	plaintext = append(plaintext, randomBytes(rand.Int()%5)...)

	if len(plaintext)%16 != 0 {
		missingBytes := 16 - len(plaintext)%16
		plaintext = append(plaintext, randomBytes(missingBytes)...)
	}
	if rand.Int()%2 == 0 {
		// Encrypt CBC
		return CBCEncrypt(randomBytes(16), randomBytes(16), plaintext)
	} else {
		// Encrypt ECB
		return set1.ECBEncrypt(randomBytes(16), plaintext)
	}
}
func challenge11() {
	for {
		fmt.Print("Enter plaintext: ")
		var input string
		fmt.Scanln(&input)
		fmt.Println(randomEncrypt([]byte(input)))

		fmt.Println(randomEncrypt([]byte("00000000000000000000000000000000000000000")))
	}
}

func main() {
	challenge11()
}
