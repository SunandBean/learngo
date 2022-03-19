package main // main -> only for compiling

import "fmt"

func multiply(a int, b int) int {
	return a * b
}

func multiply2(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(multiply(2, 2))
	fmt.Println(multiply2(2, 2))
}
