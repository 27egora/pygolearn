package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(ReverseString("привет"))
	fmt.Println(ReverseString("👋🌍"))
	fmt.Println(ReverseString("топот"))
}

func ReverseString(s string) (string, bool) { // исходное слово
	runes := []rune(s)    // слово в слайс ["_", "_", "_", "_" ...]
	slices.Reverse(runes) // переворот слова
	return string(runes), s == string(runes)
}
