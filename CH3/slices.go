package main

import "fmt"

func main() {

	// To declare a slice, is almost the same as an array:
	var (
		sl1 []int // Inside the brackets is empty
		sl2 = []int{3, 6, 11}
		sl4 = []float32{1.2, 4.5, 6.3, 8.9, 11.1}
	)

	fmt.Printf("Slice of ints %v\n", sl1)

	// There is no value, so in this case it is nil
	// nil represent the lack of value, slices are only comparable to it

	fmt.Printf("Is sl1 empty?: %v\n", sl1 == nil)

	// sl1[0] = 1 This does not work because it it len 0 and the position is out of bounds

	// There are different functions: len, append

	fmt.Printf("Length of sl1: %v\n", len(sl1))

	sl1 = append(sl1, 12)
	fmt.Printf("Is sl1 empty?: %v\nsl1 is: %v\n", sl1 == nil, sl1)

	sl1 = append(sl1, sl2...) // ... is an operator
	fmt.Printf("sl1 with sl2 appended: %v\n", sl1)

	// Go is a call by value language, creates copies and then return them.

	// the cap() function returns the capacity in memory, if appending and more memory is requiered,
	// it doubles its capacity

	// make() function, it let declare a slice with a lenght and capacity
	sl3 := make([]int, 0, 8)
	fmt.Printf("sl3 cap is: %v, and length is: %v\n", cap(sl3), len(sl3))

	// Appending to sl3
	sl3 = append(sl3, sl1...)
	fmt.Printf("sl3 cap is: %v, and length is: %v\n", cap(sl3), len(sl3))

	// The goal is to minimize the number of times the slice needs to grow

	// Slicing is the same as other languages
	sl_sub1 := sl4[1:]  // from second element to end
	sl_sub2 := sl4[1:4] // from second to 5th-1
	fmt.Printf("slice 1 is: %v\n", sl_sub1)
	fmt.Printf("slice 2 is: %v\n", sl_sub2)
	// Slicing slices share memory, if one is changed, the other also will change.
	sl_sub1[1] = 0.5
	fmt.Printf("Slice[1] was 6.3, now is: %v, and sl4 now is: %v\n", sl_sub1, sl4)
	fmt.Printf("Slice 2 also shares memory: %v\n", sl_sub2) // Slice 2 was also affected.

	// Full slice expressions make memory independent slices
	sl_independent := sl4[2:5:5] // Last number indicates the last position avaible for the sub slice
	fmt.Printf("The independent slice is: %v\n", sl_independent)
	sl_independent = append(sl_independent, 10.0) // If we append, it does not share memory
	fmt.Printf("The independent slice is: %v and sl4 is: %v\n", sl_independent, sl4)
}
