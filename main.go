package main

import "fmt"

func main() {
	var age uint8 = 23;
	name := "jono";
	
	// Print
	fmt.Print("Hello, ");
	fmt.Print("world! \n");
	fmt.Print("new line \n");

	// Println
	fmt.Println("Hey guys!");
	fmt.Println("Later!");
	fmt.Println("My age is", age, "and my name is", name);

	// Printf
	fmt.Printf("My age is %d and my name is %s\n", age, name);
	fmt.Printf("My age is %d and my name is %q\n", age, name);
	fmt.Printf("age is of type %T\n", age);
	fmt.Printf("You scored %f ponints!\n", 225.55);
	fmt.Printf("You scored %0.0f ponints!\n", 225.55);

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %v\n", age, name);
	fmt.Println("the saved string is", str);

}
