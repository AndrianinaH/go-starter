package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")

	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var firstVar = "fomby"
	fmt.Println(firstVar)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// The := syntax is shorthand for declaring and initializing a variable,
	// e.g. for var wanyKamo string = "fomby foana" in this case.
	wanyKamo := "fomby foana"
	fmt.Println("fomby ve ? " + wanyKamo)

}
