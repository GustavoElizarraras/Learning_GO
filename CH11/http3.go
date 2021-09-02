package main 

import(
	"net/http" // Client type por make requests and receive responses
    "time"
)

// Middleware pattern
// It is used to perfom multiple handlers, it uses a function that takes in
// an http.Handler instance and returns an http.Handler (usually a closure)
// that is converted to an http.HandlerFunc 

func RequestTimer(h http.Handler) http.Handler {
	// this middleware generator provides timing of requests
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		log.Printf("request time for %s: %v", r.URL.Path, end.Sub(start))
	})
}

var securityMsg = []byte("Incorrect password\n")
// this middleware implementation uses the worst access control imaginable
func TerribleSecurityProvider(password string) func(http.Handler) http.Handler {
	// We pass the password, the returned functions uses it and then returns 
	// a closure inside a closure
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Secret-Password") != password {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write(securityMsg)
				return
		}
		h.ServeHTTP(w, r)
		})
	}		
}

// First, we do setup operations or checks, if checks don't pass, we write the 
// output in the middleware (an error code) and return. If all good, we call
// the handler's ServeHTTP method and when that returns, we clean it up

// We add middle to out requests handlers by chaining them:
terribleSecurity := TerribleSecurityProvider("Gopher")

mux.Handle("/hello", terribleSecurity(RequestTimer(
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello\n"))
	}))))

// We get back out middleware from TerribleSecurityProvider and then wrap our
// handler in a series of function calls that follow this order: 
// terribleSecurity closure -> RequestTimer -> actual request handler

// Because *http.ServeMux implements the http.Handler interface, you can apply 
// middleware to all of the handlers registered with a single request router:

wrappedMux := terribleSecurity(RequestTimer(mux))
s := http.Server{
	Addr: ":8080",
	Handler: wrappedMux,
}


// Use idiomati third-party modules to enhance the server
// If you don't like the function chains for middleware, you can use alice, it
// allows the following syntax:

helloHandler := func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
chain := alice.New(terribleSecurity, RequestTimer).ThenFunc(helloHandler)
mux.Handle("/hello", chain)

// A disadvantage of *http.ServeMux is that it does not allow you to specify
// handlers based on an HTTP verb or header, nor provides support for variables
// in the URL path. Also, nesting instances is clunky. Other projects to replace
// it are gorilla mux and chi.

func main() {
	
}