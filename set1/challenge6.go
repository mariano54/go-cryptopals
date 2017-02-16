package main

import (
  "errors"
  "fmt"
  "io/ioutil"
  "path/filepath"
)

func CalculateHammingDistance(word1 string, word2 string) (int, error) {
  if len(word1) != len(word2) {
    return 0, errors.New("Wrong length")
  }
  word1Bytes := []byte(word1)
  word2Bytes := []byte(word2)
  substitutionsRequired := 0
  for i, _ := range word1Bytes {
    if word1Bytes[i] != word2Bytes[i] {
      diff := int(word1Bytes[i] ^ word2Bytes[i])
      numDiff := 0
      for diff != 0 {
        numDiff += (diff & 0x1)
        diff >>= 1
      }
      substitutionsRequired += numDiff
    }
  }
  return substitutionsRequired, nil
}

func FindKeySize(dat []byte) int {
  for KEYSIZE:=2; KEYSIZE<40; KEYSIZE++ {
    total := 0.0
    for i:=0; i<30; i++ {
      for j:=i; j<30; j++ {
        if i != j {
          first := string(dat[i*KEYSIZE:(i+1)*KEYSIZE])
          second := string(dat[j*KEYSIZE:(j+1)*KEYSIZE])
          dist, _ := CalculateHammingDistance(first, second)
          // fmt.Println(KEYSIZE, ": ", i, j, float64(dist)/float64(KEYSIZE))
          total += float64(dist)/float64(KEYSIZE)
        }
      }        
    }
    fmt.Println(KEYSIZE, total)
  }
  return 2;
}

func main() {
  absPath, _ := filepath.Abs("ciphertext.txt")
  datBase64, _ := ioutil.ReadFile(absPath)
  dat, err := Base64ToBytes(string(datBase64))
  fmt.Println(dat, err)
  fmt.Println(BytesToBase64(dat))
  KEYSIZE := FindKeySize(dat)
  
  for KEYSIZE=2; KEYSIZE<40; KEYSIZE++ {
    fmt.Println(KEYSIZE)
    var parsed [][]byte
    numBlocks := len(dat)/KEYSIZE;
    for index:=0; index<KEYSIZE; index++ {
      indexSlice := []byte{}
      for blockNum:=0; blockNum<numBlocks; blockNum++ {
        indexSlice = append(indexSlice, dat[blockNum*KEYSIZE + index])
      }
      parsed = append(parsed, indexSlice)
    }
    
    for _, block := range parsed {
      // fmt.Println(i, block)
      BreakRKX(block)
    }
  }
  // fmt.Println(parsed)
}
