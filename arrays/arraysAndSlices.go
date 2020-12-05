package main

import "fmt"

// Arrays a collection of elements of the same type

var names [5]string //Array name, length is 5 type is string

func main() {
	// =============================================ARRAYS======================================================
	names[0] = "Jim"
	names[1] = "Jones"
	names[2] = "Jane"
	names[3] = "Agatha"
	names[4] = "Sierra"
	fmt.Println(names)

	// Shorthand Notation
	keys := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(keys)
	// Declaring arrays where some indexes are partially unfilled
	keysV2 := [10]int{10, 11, 12}
	fmt.Println(keysV2)
	// Arrays without a set length, the compilers determines the length
	keysV3 := [...]int{20, 21, 22, 23, 24, 25, 26, 27, 28, 29}
	fmt.Println(keysV3)
	// The size of an array is part of its type, it cannot be resized

	// Multidimensional Arrays
	multi := [4][2]string{
		{"Kim", "L"},
		{"Kent", "Can"},
		{"Dau", "S"},
		{"Ms", "wee"}, // The last comma is important
	}
	fmt.Println(multi)
	// Use range to access key, value. Use of the blank identifier is allowed too
	for k, v := range keysV3 {
		fmt.Println(k, v, "for range loop")
	}
	for i := 0; i < len(keysV3); i++ {
		fmt.Println(i, "i for loop")
	}
	// ===========================================SLICES========================================================
	// References to data held in an array, they do not hold data on their own
	var sliceKeys3 []int = keysV3[0:5] // create a slice from existing array
	fmt.Println(sliceKeys3, "sliceKeys3")

	slice2 := []int{5, 6, 8} // Creates an array but returns a slice, Any changes to the mother array affects the slice. It does not own any data
	fmt.Println(slice2)
	slice3 := keysV3[8:10]
	fmt.Println(slice3)
}
