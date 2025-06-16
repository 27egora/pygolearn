package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(ReverseString("привет"))
	fmt.Println(ReverseString("👋🌍"))
}

func ReverseString(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}
