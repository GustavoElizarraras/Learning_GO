package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Custom JSON Parsing
// Overriding functionality, maybe to deal with other time formats. We
// create a new type that implements two interfaces: json.Marshaler and
// json.Unmarshaler

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID          string      `json:"id"`
	Items       []Item      `json:"items"`
	DateOrdered RFC822ZTime `json:"date_ordered"`
	CustomerID  string      `json:"customer_id"`
}

type RFC822ZTime struct {
	// embedded a time.Time so we can access to other methods
	time.Time
}

func (rt RFC822ZTime) MarshalJSON() ([]byte, error) {
	// the method that reads is declared on a value receiver
	out := rt.Time.Format(time.RFC822Z)
	return []byte(`"` + out + `"`), nil
}

func (rt *RFC822ZTime) UnmarshalJSON(b []byte) error {
	// the method that modifies is declared on pointer receiver
	if string(b) == "null" {
		return nil
	}
	t, err := time.Parse(`"`+time.RFC822Z+`"`, string(b))
	if err != nil {
		return err
	}
	*rt = RFC822ZTime{t}
	return nil
}

func main() {

	data := `
	{
		"id": "12345",
		"items": [
			{
				"id": "xyz123",
				"name": "Thing 1"
			},
			{
				"id": "abc789",
				"name": "Thing 2"
			}
		],
		"date_ordered": "01 May 20 13:01 +0000",
		"customer_id": "3"
	}`

	var o Order
	err := json.Unmarshal([]byte(data), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
	fmt.Println(o.DateOrdered.Month())
	out, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	// We have allowed the date formar of the JSON we are processing to change
	// types of the fields in our data structure. This is a drawback to the
	// encoding/json approach, but with json.Marshaler/Unmarshaler is too complicated
	// because we need to handle all fields

	// Use one struct for converting to and from JSON, another for data processing.
	// Read in JSON to your JSON-aware type, and then copy it to the other; for
	// writing out JSON, do the reverse. Business logic separated

	// A posibility is passing a map[string]interface{} to json.Marshal/Unmarshal
	// but only do this for the exploratory phase of your code
}
