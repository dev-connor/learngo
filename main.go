package main

import (
	"fmt"

	"github.com/dev-connor/learngo/myDict"
)

type Dictionary map[string]string

func main() {
	dictionary := myDict.Dictionary{}
	dictionary["hello"] = "hello"
	fmt.Println(dictionary)

}
