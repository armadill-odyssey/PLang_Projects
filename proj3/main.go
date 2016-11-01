package main

import (
	"fmt"
	"os"
)

func main() {

	// #1 Dangling else
	// Go doesn't allow shorthand if statements, nor even ternary operators
	// Therefore it is impossible to have a dangling else or other ambiguity
	if 3 == 3 {
		fmt.Println("First is true")
		if false {
			fmt.Println("This won't print")
		}
	} else {
		fmt.Println("This won't print")
	}
	fmt.Println()

	//#2 a
	// This will print:
	// i = 1
	// i = 2
	// This is becuause Go will allow updating of the loop limit
	var i int
	n := 3
	for i = 1; i <= n; i++ {
		fmt.Printf("i = %d\n", i)
		n = 2
	}
	fmt.Println()

	// #2 b
	// This will print:
	// j = 1
	// This is becuause Go will allow updating of the loop index
	var j int
	for j = 1; j <= 3; j++ {
		fmt.Printf("j = %d\n", j)
		j = 3
	}

	// #3

	// #4

	fmt.Println()
	// #5
	// Go does not support traditional try/catch exception handling
	// Idiomatic go is to return a possible error from functions, since functions can return multiple values
	// This will return an error since this file does not exist
	file, err := os.Open("filename.ext")
	// Then check that the error is not nil
	if err != nil {
		// Handle error here
		fmt.Println("Could not find file")
		return
	}
	// Do something with file here
	fmt.Println(file)

}
