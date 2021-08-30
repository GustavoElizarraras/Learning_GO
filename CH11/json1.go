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
// for the name, if the field should be left out when it is empty, add ,omitempty
// after the name (it is not the zero value), it works with zero-length slices or maps.
// Struct tags are never evaluated automatically, they are processed when a struct 
// instance is passed into a function.

// Unmarshaling and Marshaling

// The Unmarshal function is used to convert a slice of bytes into a struct
var o Order // Order is a struct
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

// JSON, Readers and Writers

// ioutil.ReadAll to copy the contents of an io.Reader into a byte slice so it can 
// be read by json.Unmarshal. Simalarly, writing to an inmemory byte slice buffer 
// using json.Marshal and then write to the disk or network; these are inefficient. 
// json.Decoder and json.Encoder types read from and write to anything that meets the 
// io.Reader and io.Writer interfaces

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
toFile := Person {
	Name: "Fred",
	Age: 40,
}

tmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
if err != nil {
	panic(err)
}
defer os.Remove(tmpFile.Name())
// toFile to a temporary file by passing tmpFile 
err = json.NewEncoder(tmpFile).Encode(toFile)
if err != nil {
	panic(err)
}
err = tmpFile.Close()
if err != nil {
	panic(err)
}

// Once toFile is written, we can read the JSON back in by passing a
// reference to the temp file to json.NewDecoder and then calling the 
// Decode method on it with a variable of type Person

tmpFile2, err := os.Open(tmpFile.Name())
if err != nil {
	panic(err)
}
var fromFile Person 
err = json.NewDecoder(tmpFile2).Decode(&fromFile)
if err != nil {
	panic(err)
}
err = tmpFile2.Close()
if err != nil {
	panic(err)
}
fmt.Printf("%+v\n", fromFile)

func main() {

}