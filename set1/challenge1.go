package main

import (
  "bytes"
  "errors"
  "strings"
)

const numOffset byte = 48
const letterOffset byte = 97

var hexByteValues = [...]byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 97, 98, 99, 100, 101, 102}
var base64Chars = [...]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
  'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
  'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2',
  '3', '4', '5', '6', '7', '8', '9', '+', '/'}

func IndexOf(slice [64]rune, value rune) int {
  for p, v := range slice {
    if v == value {
      return p
    }
  }
  return -1
}

func HexToBytes(hex string) ([]byte, error) {
  hex = strings.ToLower(hex)
  output := make([]byte, 0)
  for i, charByte := range hex {
    byteValue := bytes.IndexByte(hexByteValues[:], byte(charByte))
    if byteValue < 0 {
      return nil, errors.New("Invalid character")
    }
    if i%2 == 0 {
      output = append(output, byte(byteValue<<4))
    } else {
      output[i/2] += byte(byteValue)
    }
  }
  return output, nil
}

func BytesToHex(bytes []byte) string {
  output := make([]rune, 0)
  for _, b := range bytes {
    value1 := b >> 4
    value2 := (b << 4) >> 4
    output = append(output, rune(hexByteValues[value1]))
    output = append(output, rune(hexByteValues[value2]))
  }
  return string(output)
}

func BytesToBase64(input []byte) (string, error) {
  bytes := make([]byte, len(input))
  numPadding := 0
  for len(bytes)%3 != 0 {
    bytes = append(bytes, 0)
    numPadding++
  }
  copy(bytes, input)
  output := make([]rune, 0)
  for i := 0; i < len(bytes); i += 3 {
    firstChar := bytes[i] >> 2
    secondChar := ((bytes[i] << 6) >> 2) + bytes[i+1]>>4
    thirdChar := ((bytes[i+1] << 4) >> 2) + bytes[i+2]>>6
    fourthChar := ((bytes[i+2] << 2) >> 2)
    output = append(output, base64Chars[firstChar])
    output = append(output, base64Chars[secondChar])
    output = append(output, base64Chars[thirdChar])
    output = append(output, base64Chars[fourthChar])
  }
  if numPadding > 0 {
    output[len(output)-1] = '='
  }
  if numPadding > 1 {
    output[len(output)-2] = '='
  }
  return string(output), nil
}

func Base64ToBytes(input string) ([]byte, error) {
  // Iterate through base64 runes of the string
  if len(input)%4 != 0 {
    return nil, errors.New("Invalid padding")
  }
  inputNoNL := ""
  for _, r := range input {
    if r != 10 {
      inputNoNL += string(r)
    }
  }

  strIndex := 0
  totalBytes := make([]byte, 3*len(inputNoNL)/4)
  totalBytesIndex := 0
  for _, r := range inputNoNL {
    base64Index := IndexOf(base64Chars, r)
    if base64Index == -1 {
      if r == '=' {
        base64Index = 0 // Padding
      } else {
        return nil, errors.New("Invalid character")
      }
    }
    switch strIndex % 4 {
    case 0:
      totalBytes[totalBytesIndex] = byte(base64Index) << 2
    case 1:
      totalBytes[totalBytesIndex] += byte(base64Index) >> 4
      totalBytesIndex++
      totalBytes[totalBytesIndex] = byte(base64Index) << 4
    case 2:
      totalBytes[totalBytesIndex] += byte(base64Index) >> 2
      totalBytesIndex++
      totalBytes[totalBytesIndex] = byte(base64Index) << 6
    case 3:
      totalBytes[totalBytesIndex] += byte(base64Index)
      totalBytesIndex++
    }
    strIndex++
  }
  return totalBytes, nil
}

func HexToBase64(hexString string) (string, error) {
  bytes, err := HexToBytes(hexString)
  if err != nil {
    return "", errors.New("Hex conversion failed")
  }
  return BytesToBase64(bytes)
}
