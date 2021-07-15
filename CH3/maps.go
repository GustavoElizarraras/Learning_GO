package main

import "fmt"

func main() {

	// Maps are decleared: map[keyType]valueType
	var nilMap map[string]int // nil map (zero length) with string keys and int values
	// Attempting to write to a nil map causes a panic, and read return zero value for the map's value type
	fmt.Println(nilMap)
	// Using := to assign a map literal
	totalSubjects := map[string]int{} // Empty map literal, you can read and write
	fmt.Println(totalSubjects)
	companies := map[string][]string{
		"Apple":     []string{"Iphone", "Mac"},
		"Microsoft": []string{"Surface", "Windows"},
		"Google":    []string{"Pixel", "ChromeOS"}, // The last comma is neccesary
	}
	// Without redundancies
	companies2 := map[string][]string{
		"Apple":     {"Iphone", "Mac"},
		"Microsoft": {"Surface", "Windows"},
		"Google":    {"Pixel", "ChromeOS"}, // The last comma is neccesary
	}
	fmt.Println(companies)
	fmt.Println(companies2)
	fmt.Println(companies2["Microsoft"]) // Calling the values of the key "Microsoft"

	// If known how many key value pair will be, but do not know the values, use make(
	ages := make(map[int][]string, 10)
	fmt.Println(ages)
	ages[5] = []string{"Aldo", "Pancho"} // Populating map
	fmt.Println(ages)
	fmt.Printf("This map has %v value-pairs\n", len(ages)) // Using len()
	// maps are not comparable

	// comma ok (, ok) Idiom
	// It is used to check if a key is in a map
	value, ok := companies2["Google"]
	fmt.Println(value, ok)
	value, ok = companies2["IBM"] // Also, if key is not present, return the zero value
	fmt.Println(value, ok)

	// Deleting
	delete(ages, 5)
	fmt.Printf("ages map has %v value-pairs\n", len(ages)) // Using len()

	// Simulating sets
	intSet := map[int]bool{} // The values are boolean
	vals := []int{1, 3, 2, 4, 7, 9, 6, 10, 2, 7}
	for _, v := range vals {
		intSet[v] = true // Assigning true to the values
	}
	fmt.Printf("Length of the slice of values: %v, length of the map: %v\n", len(vals), len(intSet))
	// The length of the map is less, because duplicates are not allowed in a map
	fmt.Printf("Does the \"set\" include 5?: %v\n", intSet[5])
	fmt.Printf("Does the \"set\" include 10?: %v\n", intSet[10])

}
