package main

import "fmt"

func revertString(s string) string {
	runeString := []rune{}
	//this is used for correct unicode handling
	for _, v := range s {
		runeString = append(runeString, v)
	}
	for i := 0; i < len(runeString)/2; i++ {
		runeString[i], runeString[len(runeString)-1-i] = runeString[len(runeString)-1-i], runeString[i]
	}
	return string(runeString)
}

func main() {
	s := "ab⌘d日本語e⌘"
	fmt.Println(revertString(s))
}
