package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun"
	fmt.Println(reverseWords(s))
}
