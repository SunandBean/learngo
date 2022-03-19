package main // main -> only for compiling

import "fmt"

func main() {
	name := "nico" // == var name string = "nico", go will automatically find the type, only can use in function
	name = "lynn"
	fmt.Println(name)
}
