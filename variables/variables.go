package main

import (
	"fmt"
)

// Variables can be instantiated using the keyword var followed by a list of the variable names, variable type comes last
// variables stick to one type, once declared only the value can change

var width, height int

// when declaring variables of different types in the same statment

var (
	name = "Osca"
	age  = 2021
)

func main() {
	// shorthand declaration of variables is another option, here type is set by golang (declaring and instantiang)
	// shorthand can only be done within a function
	rating := 9

	// shorthand declaration of multiple cariables
	name, age := "Osca", 2021

	// shorthand declaration of multiple variables can only be used when at least one of the variables is new
	a, b := 20, 30 // declare variables a and b
	b, c := 40, 50 // b is already declared but c is new
	b, c = 80, 90  // assign new values to already declared variables b and c

	// variables can also be assigned to results of computations

	d := (a + b + c)
	fmt.Println("My name is", name)
	fmt.Println("My age is ", age)
	fmt.Println("I have a rating of ", rating, "out of 10")
	fmt.Println(d)
}
