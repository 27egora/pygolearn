package main

import "fmt"

func fibonacci(limit int) []int {
	sequence := []int{0, 1}
	for {
		next := sequence[len(sequence)-1] + sequence[len(sequence)-2]
		if next > limit {
			break
		}
		sequence = append(sequence, next)
	}
	return sequence
}

func main() {
	fib := fibonacci(100)
	fmt.Println("Ряд Фибоначчи до 100:")
	fmt.Println(fib)
}
