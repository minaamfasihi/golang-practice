package main

import "fmt"

var i int = 23

func main() {
	fmt.Println(i) // 23
	var i int = 13
	fmt.Println(i) // prints 13. variable in the innermost block takes precedence. this is called shadowing	

	var j int = 12 // wont compile because j is declared but not used.
}