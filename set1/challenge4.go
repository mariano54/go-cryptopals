package set1

import (
	"bufio"
	"log"
	"os"
)

func BreakRKX(bytes []byte) int {
	key := make([]byte, len(bytes))
	max := 0.0
	maxSolution := 0
	for char := 0; char < 256; char++ {
		for i, _ := range key {
			key[i] = byte(char)
		}
		res, _ := Xor(bytes, key)
		calcd := Score(string(res))
		if calcd >= max {
			max = calcd
			maxSolution = int(char)
		}
	}
	return maxSolution
}

func challenge4() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(dir + "/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes, _ := HexToBytes(scanner.Text())
		BreakRKX(bytes)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
