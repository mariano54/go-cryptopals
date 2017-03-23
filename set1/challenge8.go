package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func hash(arr []byte) int {
	product := 1
	for _, item := range arr {
		product *= int(item)
	}
	return product
}

func challenge8() {
	absPath, err1 := filepath.Abs("ciphertext-8.txt")
	HandleError(err1)
	file, err2 := os.Open(absPath)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ciphertexts := [][]byte{}
	for scanner.Scan() {
		ciphertext, err3 := Base64ToBytes(scanner.Text())
		HandleError(err3)
		ciphertexts = append(ciphertexts, ciphertext)
	}

	for x, c := range ciphertexts {
		m := map[int]int{}
		for i := 0; i < len(c)/16; i++ {
			block := c[i*16 : (i+1)*16]
			m[hash(block)] = m[hash(block)] + 1
		}
		for _, value := range m {
			if value > 1 {
				fmt.Println(x, value, true)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	challenge8()
}
