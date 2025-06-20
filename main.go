package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		arg1 := args[1]
		reversed, isPalindrome := ReverseString(arg1)
		fmt.Println("Исходное слово:", arg1)
		fmt.Println("Перевернутое:", reversed)
		fmt.Println("Палиндром:", isPalindrome)
	} else {
		fmt.Println("Указать ОДНО слово в качестве аргумента!")
	}
}

func ReverseString(s string) (string, bool) {
	runes := []rune(s)
	slices.Reverse(runes)
	reversed := string(runes)
	return reversed, s == reversed
}
