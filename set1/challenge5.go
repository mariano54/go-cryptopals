package set1

import (
	"fmt"
	"os"
)

func RKXEncrypt(key []byte, message []byte) []byte {
	result := make([]byte, len(message))
	for i, b := range message {
		result[i] = key[i%len(key)] ^ b
	}
	return result
}

func challenge5() {
	key := []byte(os.Args[1])
	message := []byte(os.Args[2])
	c := RKXEncrypt(key, message)
	hex := BytesToHex(c)
	fmt.Println(hex)
}
