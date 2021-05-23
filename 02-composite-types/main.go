package main

import "fmt"

func main() {

	// --- array ---
	// array, fixed length
	// name [size]type{elements}
	// var arr [3]int
	// [0 0 0]
	// var arr = [3]int{1, 2, 3}
	// [1 2 3]

	// [size]type{index:value}
	// var arr = [10]int{5: 1}
	// [0 0 0 0 0 1 0 0 0 0]

	// multi-dimensions array
	var arr = [2][3]int{}
	fmt.Println(arr)
	// [[0 0 0] [0 0 0]]

	// compare array, must be the same type
	foo := [3]int{1, 2, 3}
	bar := [...]int{1, 2, 3}
	fmt.Println(foo == bar)
	// true

	// --- slice ---
	// slice, like array but dynamic length
	var slice = []int{1, 2, 3}
	// [1 2 3]

	var secondSlice = []int{4, 5, 6}
	slice = append(slice, secondSlice...)
	fmt.Println(slice)
	// [1 2 3 4 5 6]

	//using `make` to create a new slice. make(type, size, capacity)
	testMake := make([]int, 5, 10)
	fmt.Println(testMake)
	// [0 0 0 0 0]

	// if we create a slice from the slice, it shared memory.
	a := slice[1]
	b := slice[:2]
	slice[1] = 20
	fmt.Println(a, b)
	// 2 [1 20]

	// create a slice independently
	c := make([]int, 2)
	copy(c, slice)
	slice[1] = 200
	fmt.Println(c)
	// [1, 20]
}
