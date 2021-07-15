package main

import "fmt"

func main() {

	// Array declaration
	var arr1 [5]int

	var (
		arr2  = [3]int{1, 2, 3}
		arr3  = [10]int{2, 4: 1, 5: 7}      // Alocate in positions and all left are 0's
		arr2d [3][3]string                  // 2D array
		arr4  = [...]float64{1.2, 1.1, 6.4} // Array size created by values
	)
	fmt.Printf("Array 1 with only size especified: %v\n", arr1)
	fmt.Printf("Array 2 has all values: %v\n", arr2)
	fmt.Printf("Array 3 has prealocated values: %v\n", arr3)

	// Asigning a value in n position
	arr1[3] = 1 // n start in 0
	fmt.Printf("Array 1 modified: %v\n", arr1)

	// 2D array
	fmt.Printf("2D array of empty strings: %v\n", arr2d)

	// Array sized by elements
	fmt.Printf("This array was created with [...]: %v\n", arr4)

	// Arrays of different size are different type
	// fmt.Printf("arr1 and arr2 are the same? : %v\n", arr1 == arr2) The compiler output an error

	// Also, we can't convert arrays of different sizes to identical types
	// nor assign to the same variable.

	// ONLY USE arrays if the exact length is known
}
