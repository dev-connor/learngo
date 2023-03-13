package main

import (
	"fmt"

	"github.com/dev-connor/learngo/myDict"
)

func main() {
	dictionary := myDict.Dictionary{}
	dictionary["hello"] = "hello"
	fmt.Println(dictionary)
}
