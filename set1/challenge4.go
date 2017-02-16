package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

func BreakRKX(bytes []byte) {
  key := make([]byte, len(bytes))

  for char := 0; char < 256; char++ {
    for i, _ := range key {
      key[i] = byte(char)
    }
    res, _ := Xor(bytes, key)
    calcd := Score(string(res))
    if calcd >= 0.25 {
      fmt.Println(calcd, string(res))
    }
  }
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
