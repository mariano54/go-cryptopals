func CalculateHammingDistance(word1 string, word2 string) {
  substitutionsRequired := 0
  for i, _ := range word1 {
    if strings.CharAt(word1, i) != string.CharAt(word2, i) {
      substitutionsRequired++
    }
  }
}
