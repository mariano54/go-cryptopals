package main

import "errors"

func Xor(arr1 []byte, arr2 []byte) ([]byte, error) {
  if len(arr1) != len(arr2) {
    return nil, errors.New("Arrays are different lengths")
  }
  output := make([]byte, len(arr1))
  for i, _ := range arr1 {
    output[i] = arr1[i] ^ arr2[i]
  }
  return output, nil
}
