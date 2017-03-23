package main

import "fmt"

var blockSize = 16

func challenge10() {
	// key := []byte("YELLOW SUBMARINE")
	// cipher, err1 := aes.NewCipher(key)
	// if err1 != nil {
	// 	log.Panic(err)
	// }
	// dat, err2 := set1.ReadBase64DataFromFile("ciphertext-10.txt")
	//
	IV := "0000000000000000"
	for i := 0; i < len(IV); i++ {
		fmt.Printf("%x", IV[i])
	}
	fmt.Println()
	fmt.Printf("%x", IV)
	// encryptedBlocks := make([]byte, len(data))
	// for blockIndex := 0 {
	//   // If there is no more data to read
	//   if (blockIndex + 1) * blockSize > len(  data) - (blockIndex * blockSize) {
	//     break
	//   }
	//   encryptedBlocks = append(encryptedBlocks, )
	// }
}

func main() {
	challenge10()
}
