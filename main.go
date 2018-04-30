package main

import "fmt"

func main() {

	fmt.Printf("=> %+v\n", myFunctionUse())
}

func myFunction(a, b int) int {
	return a + b
}

var myFunc func(int, int) int

func myFunctionUse() string {
	return fmt.Sprint("=> ", myFunc(10, 20))
}

func init() {
	myFunc = myFunction

}
