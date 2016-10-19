package main

import "fmt"

var a int = 4

func main() {
	fmt.Println("Test scoping")
	a := 1
	testLexical()
	
	// prints the a in the local scope
	fmt.Println(a)

	// Go cannot alter scope
	// This is an illegal operation, there is no function in go that can do something like this.
	// testLexical().bind(main)

	//testing overloading
	// Go does not allow method overloading, this won't compile
	//testOverload("Hey")
	testOverload(3)

	// test operator overloading
	fmt.Println("Test if Go permits operator overloading")
	b := 1
	fmt.Println(b)
	// Go allows operator overloading
	b += 1
	fmt.Println(b)

	fmt.Println("Test implicit type conversion")
	var c int = 6/4
	// Should cast to int by dropping off decimal places of the float
	fmt.Println(c)

	// test for alias
	fmt.Println("Test for aliases")
	d := 4
	e := &d
	f := e
	// f and e are aliases to the references of d
	fmt.Println("e: ", *e, " f: ", *f)
	d = 5
	// e and f should have changed also
	fmt.Println("e: ", *e, " f: ", *f)

	// Test for dangling references
	g := &d
	fmt.Println("d: ", d, " g: ", g, " *g: ", *g)
	// This action is not allowed in go, you can not assign to nil nor delete objects
	// Go has garbage collection built in, so you can't run into dangling references
	//d = nil
	fmt.Println("d: ", d, " g: ", g, " *g: ", *g)
}

// Would print the a defined in main if it was dynamically scoped
func testLexical() {
	fmt.Println(a)
}

//func testOverload(a string) {
//	fmt.Println("This function, testOverload, prints a type string")
//	fmt.Println(a)
//}

func testOverload(a int) {
        fmt.Println("This function, testOverload, takes a type int and prints it")
        fmt.Println(a)
}
