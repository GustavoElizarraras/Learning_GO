package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {

	// Encoding and Decoding JSON Streams
	// When we have multiple JSON structs to read or write at once, json.Decoder
	// and json.Encoder can be used. Lets assume we have the following data:

	const data = `
	{"name": "Fred", "age": 40} 
	{"name": "Mary", "age": 21} 
	{"name": "Pat", "age": 30}
	`
	var t struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	dec := json.NewDecoder(strings.NewReader(data))
	var b bytes.Buffer         // write into  variable of this type
	enc := json.NewEncoder(&b) // write out single or multiple values

	// We use the More method on json.Decoder as a for loop condition,
	// one object at a time
	for dec.More() {
		// we will store data in the t variable
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}
		// process t
		fmt.Println(t)
		err = enc.Encode(t)
		if err != nil {
			panic(err)
		}
	}

	out := b.String()
	fmt.Println(out)

	// // Writting out multiple values with json.Encoder

	// for _, input := range allInputs {
	// 	t := process(input)
	// 	err = enc.Encode(t)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

}
