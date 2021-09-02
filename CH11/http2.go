package main 

import(
	"net/http" // Client type por make requests and receive responses
    "time"
)

// The Server
// It is resposible for listening for HTTP request, it is a performant 
// HTTP/2 server that supports TLS.
// A request to a server is handled by an implementation of the http.Handler
// interface: 

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// *http.Request is the exact same type that's used to send a request to
// an HTTP server. The http.ResponseWriter is an interface with 3 methods:

type ResponseWriter interface {
	Header() http.Header
	WriteHeader(statusCode int)
	Write([]byte) (int, error)
}

// These methods must be called in the order presented:
// - Header: get an instance of http.Header and set any response header
// - WriteHeader: call it with the HTTP status code for the response
// - Write: set the body for the response

type HelloHandler struct{}

func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}

// You instantiate a new http.Server just like any other struct:

s := http.Server{
	Addr: ":8080"
	// specifying timeouts, set these to handle malicious or broken HTTP clients
	ReadTimeout: 30 * time.Second
	WriteTimeout: 90 * time.Second
	IdleTimeout: 120 * time.Second
	Handler: HelloHandler{},
}
err != s.ListenAndServe()
if err != nil {
	if err !=http.ErrServerClosed {
		panic(err)
	}
}

// A server that only handles a single request isn't useful, Go includes a
// request router *http.ServeMux, it meets the http.Handler interface and
// includes two methods for dispatch requests.

mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello\n"))
})
// This method takes in a function or closure and converts it to a http.HandlerFunc

// For simple handlers, a closure is sufficient, for more complex handlers,
// use a method on a struct 

// Avoid using package level functions outside of trivial test programs, like 
// http.Handle, http.HandleFunc, http.ListenAndServe and http.ListenAndServeTLS,
// they work with http.DefaultServeMux. This means you can't configure server properties. 

// An *http.ServeMux dispatches requests to http.Handler instances (this is implemented), 
// multiple related requests can be created and registered with a parent *http.ServeMux

person := http.NewServeMux()
// A request for /person/greet is handled
person.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("greetings\n"))
})
dog := http.NewServeMux()
// A request for /dog/greet is handled
dog.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("good puppy\n"))
})
mux := http.NewServeMux()
// registering person and dog with mux, and removing part of the path
// processed by mux
mux.Handle("/person/", http.StripPrefix("/person", person))
mux.Handle("/dog/", http.StripPrefix("/dog", dog))




func main() {

}