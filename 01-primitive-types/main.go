package main

import "fmt"

func main() {
	// we can declare variable by using keyword var
	// Go assign a default value to any unassigned variable
	// 0 for all integer and float types
	// false for booleans,
	// "" for strings,
	// nil for interfaces, slices, channels, maps, pointers and functions.

	// var x int
	// var y float32
	// var z string

	// declare multiple variables in the same time
	// var (
	// 	x int
	// 	y float32
	// 	z string
	// )

	// another way to declare a new varible by using := operator
	x := 1

	// := can assign value to existing variable, as long as we have a new variable
	// also we can assign many variables and types at the same time
	x, y, z := 10, 32.1, "string"

	fmt.Println(x, y, z)
}
