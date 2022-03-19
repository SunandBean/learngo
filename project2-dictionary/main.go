package main

import (
	"fmt"
	"learngo/project2-dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// word := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, _ := dictionary.Search(word)
	// fmt.Println("found ", word, "definition: ", hello)
	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }

	// baseWord := "hello"
	// dictionary.Add(baseWord, "First")
	// err := dictionary.Update(baseWord, "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// word, _ := dictionary.Search(baseWord)
	// fmt.Println(word)

	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Delete(baseWord)
	if err != nil {
		fmt.Println(err)
	}

	word, err2 := dictionary.Search(baseWord)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(word)
	}
}
