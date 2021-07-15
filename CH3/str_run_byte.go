package main

import "fmt"

func main() {
	// Strings in go are not made out of runes, they are sequences of bytes
	// it is assumed they are composed of a secuence of UTF-8-encoded code points.

	var s string = "Engineering is cool"
	var b byte = s[8]       // Strings indexes are zero-based
	var bs string = s[5:13] // Slicing also works

	fmt.Printf("Original string: %v\n", s)
	fmt.Printf("Single byte of s: %v\n", b)
	fmt.Printf("Sub string: %v\n", bs)

	// When working with emojis or other languages, they are from 1 to 4 byte long
	var ship string = "Is a ⛴"
	fmt.Printf("String: %v\n", ship)
	fmt.Printf("Sub-string: %v\n", ship[3:7]) // All ship bytes were not included

	// Conversions
	var r rune = 'x' // Single quotes
	var st1 string = string(r)
	var by byte = 'y' // SIngle quotes
	var st2 string = string(by)
	fmt.Printf("st1 %v, comes from rune: %v\n", st1, r)
	fmt.Printf("st2 %v, comes from byte: %v\n", st2, by) // Prints the number in utf-8
	//Since Go 1.15, you can't convert from int to str
	var bse []byte = []byte(ship)
	var rs []rune = []rune(ship)
	fmt.Println(bse)
	fmt.Println(rs) // Slices of runes are uncommon

}
