package main

import (
  "fmt"
  "io/ioutil"
  "path/filepath"
)

func CalculateHammingDistance(word1Bytes []byte, word2Bytes []byte) (int, error) {
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
  min := 99999.0
  minKeysize := 0
  for KEYSIZE:=8; KEYSIZE<40; KEYSIZE++ {
    total := 0.0
    for i:=0; i<5; i++ {
      for j:=i; j<30; j++ {
        if i != j {
          first := dat[(i)*KEYSIZE:(i+1)*KEYSIZE]
          second := dat[(i+1)*KEYSIZE:(i+2)*KEYSIZE]
          dist, _ := CalculateHammingDistance(first, second)
          total += float64(dist)/float64(KEYSIZE)
        }
      }        
    }
    if total < min {
      min = total
      minKeysize = KEYSIZE
    }
  }
  return minKeysize;
}

func Challenge6() {
  absPath, _ := filepath.Abs("ciphertext.txt")
  datBase64, _ := ioutil.ReadFile(absPath)
  dat, _ := Base64ToBytes(string(datBase64))
  KEYSIZE := FindKeySize(dat)
  
  var parsed [][]byte
  numBlocks := len(dat)/KEYSIZE;
  for index:=0; index<KEYSIZE; index++ {
    indexSlice := []byte{}
    for blockNum:=0; blockNum<numBlocks; blockNum++ {
      indexSlice = append(indexSlice, dat[blockNum*KEYSIZE + index])
    }
    parsed = append(parsed, indexSlice)
  }
  
  var fullKey []byte
  for _, block := range parsed {
    fullKey = append(fullKey, byte(BreakRKX(block)))
  }
  
  decrypted := []byte{}
  for i := 0; i < len(dat); i+=KEYSIZE {
    res := []byte{}
    if i+KEYSIZE > len(dat) {
      res, _ = Xor(fullKey[0:len(dat)-i], dat[i:])
    } else {
      res, _  = Xor(fullKey, dat[i:i+KEYSIZE])
    }
    decrypted = append(decrypted, res...)
  }
  fmt.Println(string(decrypted))
}

func main() {
  Challenge6()
}
