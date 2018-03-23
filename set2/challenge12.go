package main

import (
	"fmt"
	"go-cryptopals/set1"
	"go-cryptopals/util"
)

var randomKey = RandomBytes(BlockSize)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func randomECBEncrypt(plaintext []byte) []byte {
	base64Str := "Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd2" +
		"4gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBz" +
		"dGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IH" +
		"N0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK"

	strBytes, err := set1.Base64ToBytes(base64Str)

	if err != nil {
		panic("Failed to convert to base64")
	}
	plaintext = append(plaintext, strBytes...)
	return set1.ECBEncrypt(randomKey, plaintext)
}

func challenge12() {
	prefixFull := "AAAAAAAAAAAAAAAA"
	solution := ""
	blockAddition := 0
	initialBlock := true
	for i := 0; i < 10000; i++ {
		prefix := prefixFull[:BlockSize-1-(i%BlockSize)]
		if (i)%BlockSize == 0 && i > 0 {
			blockAddition++
			initialBlock = false
		}
		for j := 0; j < 255; j++ {
			length := len(randomECBEncrypt([]byte(prefix)))
			startIndex := len(prefix)/BlockSize + blockAddition
			endIndex := min(len(prefix)/BlockSize+blockAddition+BlockSize, length)
			fmt.Println(length)
			fmt.Println(startIndex, endIndex)
			target := randomECBEncrypt([]byte(prefix))[startIndex:endIndex]
			var attemptBlock []byte
			if initialBlock {
				attemptBlock = randomECBEncrypt([]byte(prefix + solution + string(j)))[0:BlockSize]
			} else {
				attemptBlock = randomECBEncrypt([]byte(solution[len(solution)-BlockSize+1:] + string(j)))[0:BlockSize]
			}
			if util.DeepEqual(target, attemptBlock) {
				solution += string(j)
				break
			}
		}
		fmt.Println(solution)
	}
	fmt.Println(solution)
}
