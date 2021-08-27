package main

import(
	"json"
)

// JSON is the standard way to communicate between the services of a REST API,
// marshaling means converting from a Go data type to an encoding, unmarshaling
// is the inverse process

// Use struct tags to add metadata
// Lets say we need to write and read this json
{ "id":"12345",
  "date_ordered":"2020-05-01T13:01:02Z",
  "customer_id":"3", 
  "items":[{"id":"xyz123","name":"Thing 1"},{"id":"abc789","name":"Thing 2"}] 
}

// We define types to map this data
// We specify the rules for processing our JSON with `struct tags`, strings 
// that are written after the fields in a struct. They are always like this:
// tagName: "tagValue". For JSON processing, we use json as the tag name, 
// if no json tag is provided, the default behaviour is to assume that the name
// of the JSON object field matches the name of the Go struct field

type Order struct {
	// Good practice to match the name of the field with the JSON onject
	ID string `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID string `json:"customer_id"`
	Items []Item `json:"items"`
}

type Item struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

// If a field should be ignored when marshaling or unmarshaling, use a dash (-)
// for the name, it the field should be left out when it is empty, add ,omitempty
// after the name (it is not the zero value), it works with zero-length slice or map
// Struct tags are never evaluated automatically, they are processed when a struct 
// instance is passed into a function.

// Unmarshaling and Marshaling

// The Unmarshal functoin is used to convert a slice of bytes into a struct
var o Order 
// populates data into an input parameter, we can reuse the same struct and it is 
// the only way to do it. Go does not have generics, so it can't specify what type 
// should be instantiated to store the bytes being read
err := json.Unmarshal([]byte(data), &o)
if err != nil {
	return err
}

// we use Marshal to write an Order onstance back as JSON, stored in a slice of bytes
out, err := json.Marshal(o)
// This two functions can r/w a struct of any type, this is thanks to reflection



func main() {

}