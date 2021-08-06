package main

import "fmt"

func checkAllSymbolsUnique(s string) bool {
	tmp := make(map[rune]bool)
	for _, v := range s {
		if _, ok := tmp[v]; ok {
			return false
		} else {
			tmp[v] = true
		}
	}
	return true
}

func main() {
	allUniqueString := "汉字漢zxcvbn"
	duplicateSymbolsString := "asd汉cvbkj汉"
	if checkAllSymbolsUnique(allUniqueString) {
		fmt.Println("All symbols in string are unique.")
	} else {
		fmt.Println("String has duplicate symbols.")
	}
	if checkAllSymbolsUnique(duplicateSymbolsString) {
		fmt.Println("All symbols in string are unique.")
	} else {
		fmt.Println("String has duplicate symbols.")
	}
}
