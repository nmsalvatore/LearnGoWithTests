package main

import (
	"fmt"
)

func main() {
	word1 := "hello"
	word2 := "theramin"
	merged := mergeAlternately(word1, word2)
	fmt.Printf("%q and %q merged make %q\n", word1, word2, merged)
}

func mergeAlternately(word1 string, word2 string) string {
	minLen := min(len(word1), len(word2))
	merged := make([]byte, 0, len(word1)+len(word2))

	for i := 0; i < minLen; i++ {
		merged = append(merged, word1[i], word2[i])
	}

	if len(word1) > minLen {
		merged = append(merged, word1[minLen:]...)
	} else if len(word2) > minLen {
		merged = append(merged, word2[minLen:]...)
	}

	return string(merged)
}
