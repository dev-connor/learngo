package main

import (
	"fmt"

	"github.com/dev-connor/learngo/myDict"
)

type Dictionary map[string]string

func main() {
	dictionary := myDict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, err := dictionary.Search(word)
	fmt.Println(hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
