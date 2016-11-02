package main

import (
	"fmt"
	"os"
)

var paramTest1 int = 1
var paramTest2 int = 2

type fn func(int, int)

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
	fmt.Println()
	// #3 Go has pass by value parameter passing. The values of paramTest1 and 2 (the values 1 and 2)
	// are being passed to ParamPassTest, and even when paramTest1 is changed in the function, the values
	// of the parameters remain the same.
	fmt.Println(ParamPassTest(paramTest1, paramTest2))
	fmt.Println()
	
	// #4 Go passes the actual function as a parameter, rather than a call to a function.

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

// #3 if paramters are evaluated at the start of the expression, as they are in pass by value,
// then x + y will be 1 + 2, and not 4 + 2
func ParamPassTest(x int, y int) int {
	paramTest1 = 4
	//if pass by value, this function should return 3 instead of 6.
	return x + y
}
func FuncPassTest(x fn) {
	x(4,2)
	// #4 no errors; Go is happy to pass X as a generic function of two int parameters.
}
