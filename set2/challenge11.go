package main

import (
	"fmt"
	"go-cryptopals/set1"
	"math/rand"
)

// RandomBytes generates a slice of random bytes
func RandomBytes(numBytes int) []byte {
	buf := make([]byte, numBytes)
	read, err := rand.Read(buf)
	if read != numBytes || err != nil {
		panic("Failed to generate random number")
	}
	return buf
}

func generateKey() []byte {
	return RandomBytes(BlockSize)
}

func randomEncrypt(plaintext []byte) []byte {
	plaintext = append(RandomBytes(rand.Int()%5), plaintext...)
	plaintext = append(plaintext, RandomBytes(rand.Int()%5)...)

	if len(plaintext)%16 != 0 {
		missingBytes := 16 - len(plaintext)%16
		plaintext = append(plaintext, RandomBytes(missingBytes)...)
	}
	if rand.Int()%2 == 0 {
		// Encrypt CBC
		return CBCEncrypt(RandomBytes(16), RandomBytes(16), plaintext)
	}
	// Encrypt ECB
	return set1.ECBEncrypt(RandomBytes(16), plaintext)
}

// DetectECB detects if an encryption function uses ECB mode
func DetectECB(encrypter func([]byte) []byte) bool {
	plaintext := make([]byte, 320)
	ciphertext := encrypter(plaintext)
	for i := 100; i < 200; i++ {
		if ciphertext[i] != ciphertext[i+BlockSize] {
			return false
		}
	}
	return true
}

func challenge11() {
	fmt.Println(DetectECB(randomEncrypt))
	fmt.Println(DetectECB(randomEncrypt))
	fmt.Println(DetectECB(randomEncrypt))
	fmt.Println(DetectECB(randomEncrypt))
	fmt.Println(DetectECB(randomEncrypt))
}
