package set1

import "fmt"

var commonLetters = [...]rune{'a', 'e', 'i', 'o', 'u', 't', ' '}

func Score(input string) float64 {
	found := 0.0
	for _, r := range input {
		for _, r2 := range commonLetters {
			if r2 == r {
				found++
				break
			}
		}
	}
	return (found / float64(len(input)))
}

func challenge3() {
	hex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	bytes, _ := HexToBytes(hex)
	key := make([]byte, len(bytes))

	for char := 0; char < 256; char++ {
		for i, _ := range key {
			key[i] = byte(char)
		}
		res, _ := Xor(bytes, key)
		fmt.Println("Score", Score(string(res)), string(res))
	}
}
