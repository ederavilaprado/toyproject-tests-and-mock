package main

import (
	"fmt"

	"github.com/ederavilaprado/toyproject-tests-and-mock/pkg/example"
)

func main() {

	fmt.Println(myFunctionUse())

	e := example.New()
	fmt.Println(runExampleFunction(e))
}

func myFunction(a, b int) int {
	return a + b
}

var myFunc = myFunction

func myFunctionUse() string {
	return fmt.Sprint("=> ", myFunc(10, 20))
}

func runExampleFunction(e example.Example) string {
	if err := e.Set("my_key", "my_value"); err != nil {
		return ""
	}
	v, err := e.Get("my_key")
	if err != nil {
		return ""
	}

	return v
}
