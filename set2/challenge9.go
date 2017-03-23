package main

import "fmt"

func PadBlock(block []byte, blockSize int) []byte {
  if len(block) >= blockSize {
    return block
  }
  numPaddingBytes := blockSize - len(block)
  for i := 0; i < numPaddingBytes; i++ {
    block = append(block, byte(numPaddingBytes))
  }
  return block
}

func Challenge9() {
  testBlock1 := []byte{12, 53, 43, 110, 12, 1}
  testBlock2 := []byte("YELLOW SUBMARINE")
  paddedBlock1 := PadBlock(testBlock1, 8)
  paddedBlock2 := PadBlock(testBlock2, 20)
  fmt.Println(paddedBlock1)
  fmt.Println(paddedBlock2)
}
