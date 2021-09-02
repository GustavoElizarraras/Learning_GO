package main 

import(
	"net/http" // Client type por make requests and receive responses
    "time"
)


// Go uses HTTP/2
// A default client instance (DefaultClient), do not use in production, 
// because it defaults to having no timeout. Instead instantiate your 
// own. Only one http.Client is necessary for your entire program, as it 
// properly handles multiple simultaneous requests across goroutines

client := &http.Client{
	Timeout: 30 * time.Second,
}

// When you want to make a request, you can create a new *http.Request
// instance with:

req, err := http.NewRequestWithContext(context.Background(), 
    http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
// (context, method, url, nil/io.reader)
// If you are making a PUT, POST, PATCH request, the last parameter is an io.Reader
if err != nil {
	panic(err)
}

// Setting headers
req.Header.Add("X-My-Client", "Learning Go")
res, err := client.Do(req)
// the result is returned in an http.Response thanks to Do()
if err != nil {
	panic(err)
}

// The response has several fields on the request:
// - response status: StatusCode field
// - response code: Status field 
// - response headers: Header field
// - returned content: Body field or type io.ReadCLoser
// This allows us to use it with json.Decoder to process REST API responses 

defer res.Body.Close()
if res.StatusCode != http.StatusOK {
	panic(fmt.Sprintf("unexpected status: got %v", res.Status))
}
fmt.Println(res.Header.Get("Content-Type"))
var data struct {
	UserID int `json:"userId"`
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}
err = json.NewDecoder(res.Body).Decode(&data)
if err != nil {
	panic(err)
}
fmt.Printf("%+v\n", data)

// There are function to make a GET, HEAD and POST calls, but they use the 
// default client, do not use them

func main() {
	
}