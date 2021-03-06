package main

import "fmt"

// #1
// Go lang has support for Cartesian product types
// Similar to C, Golang can define structs of related data
type Cartesian struct {
	t uint
	u int
	v rune // alias for int32, represents a Unicode code point
}

// #2
// Golang does not have Enumerated types, but they can be replicated in a round about way

// We can create a type to only be used for the enum
type enumInt int

// Here the values for the Enum are established
const (
	// When declaring constants, a variable iota starts at 0 and increments everytime it is used
	ZERO enumInt = iota // ZERO = 0
	ONE                 // ONE = 1
	TWO
	THREE
	FOUR
)

// Here are the strings for the enums
var numbers = [...]string{
	"ZERO",
	"ONE",
	"TWO",
	"THREE",
	"FOUR",
}

// This is the two string method for type enumInt, it return the respective index in the numbers array
func (n enumInt) String() string { return numbers[n] }

func main() {
	// #2 cont.
	// we can assign the above constants to a type enumInt
	var test1 enumInt = ONE
	// Go will also cast an int to an enumInt
	var test2 enumInt = 2
	// Go falls short of perfect enums, because this is a valid assignment,
	var test3 enumInt = 8

	fmt.Println(test1) // prints "ONE"
	fmt.Println(test2) // prints "TWO"
	// This will throw an out of bounds error becuase it can't the index '8' does not exist in the numbers array
	fmt.Println(test3)
	fmt.Println("")

	//#3 Go has static typing.
	var x int = 24

	fmt.Println(x)
	// var x int = "Hello!"
	//- Does not compile: .\main.go:43: cannot use "Hello!" (type string) as type int in assignment
	// If Go used dynamic typing this wouldn't cause an error until run time
	fmt.Println("")

	//#4 Go does indeed have short-circuit evaluation; the boolean value of 'shortcircuittest' is never evaluated because a is false.
	a := false
	if a && shortcircuittest() {
		fmt.Println("This will print if something goes horribly wrong.")
	} else {
		fmt.Println("This ALONE will print if everything is ok.")
	}

	fmt.Println("")

	// #5
	// Golang performs applicative order evaluation
	// If it was normal order, it would copy the fn call 'double(2)' to each x in the square method.
	// Therefore making two calls to the double method, so "Doubling" would be printed twice
	// But since Golang is applicative order, double(2) is evaluated first and then that value is passed on to the square() method
	result := square(double(2))
	fmt.Println(result)
}

func shortcircuittest() bool {
	fmt.Println("This will print of the second statement is evaluated.")
	return true
}

func double(x int) int {
	fmt.Println("Doubling")
	return x + x
}
func square(x int) int {
	fmt.Println("Squaring")
	return x * x
}
