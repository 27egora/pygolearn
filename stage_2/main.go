package main

import (
	"fmt"
)

func printNumbers() {
	startnum := 0
	for startnum < 11 {
		fmt.Println(startnum)
		startnum += 1
	}
}

func main() {
	printNumbers()
}
