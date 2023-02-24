// This code is to be used as a sort of cheat sheet any time I need to write a program in Go.

// MULTILINE COMMENT:
/* This is a multiline comment. */

// START OF PROGRAM:
package main; 

// PRINTING TO CONSOLE:
import ("fmt");

func printSomething() {
	fmt.Println("Hello, world!");
}

func returnSomething() int {
	return 5;
}

// FUNCTION SYNTAX:
func main() {
	// PRINTING TO CONSOLE:
	fmt.Println("Hello, world!");
	
	// VARIABLES:
	var x int = 5;
	var y int = 6;
	var z int = x + y;
	fmt.Println(z);

	// ARRAYS:
	var a [5]int;
	a[2] = 7;
	fmt.Println(a);

    // STRINGS:
	var s string = "Hello, world!";
	fmt.Println(s);

	// FOR LOOPS:
	for i := 0; i < 5; i++ {
		fmt.Println(i);
	}

	printSomething();

	returnSomething();
}
