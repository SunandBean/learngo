package main // main -> only for compiling

import (
	"fmt"

	"github.com/sunandbean/learngo/something"
)

func main() {
	fmt.Println("Hello world!")
	something.SayHello() // Start with Uppercase -> export from other package
	// something.sayBye()   // Start with Lowercase -> private function
}
