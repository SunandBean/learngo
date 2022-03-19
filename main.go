package main // main -> only for compiling

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 20
	fmt.Println(a)
}
