package main

import "fmt"

// func FuncName(funcParamter funcParamterType) returnType{
// 	funcBody
// }

func sumTwoIntegers(a int, b int) int {
	var c = a + b
	return c
}

func main() {
	fmt.Println(sumTwoIntegers(5, 6))

	// multiple return values
	// perimeter, area, volume := multipleReturnValues(6, 3.2, 6) an alternative way
	fmt.Println(multipleReturnValues(6, 3.2, 6))

	// blank identifier. In this case the values of volume and perimeter are not returned
	_, area, _ := multipleReturnValues(6, 6, 0.0)
	fmt.Println("The new area value is", area)
}

// Multiple return values
func multipleReturnValues(length, width, height float64) (float64, float64, float64) {
	perimeter := length + width*2
	area := length * width
	volume := length * width * height
	return perimeter, area, volume
}
