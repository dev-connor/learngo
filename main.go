package main

import (
	"fmt"

	"github.com/dev-connor/learngo/myDict"
)

type Dictionary map[string]string

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

}
