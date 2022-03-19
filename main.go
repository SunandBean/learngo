package main // main -> only for compiling

import "fmt"

func main() {
	names := []string{"nico", "lynn", "dal"} // no number in definition -> slice
	names = append(names, "flynn")
	fmt.Println(names)
}
