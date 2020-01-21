package main

import (
	"io"
	"net/http"

	"github.com/sasidakh/employee/employee"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

// replace this with http.HandlerFunc
var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {

	server := http.Server{
		Addr:    ":8080",
		Handler: &myHandler{},
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))

	addHandler("/", hello)
	addHandler("/employee", employee.Handler)

	server.ListenAndServe()
}

type myHandler struct{}

//ServeHTTP function to pass to myHandler
func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My Server: "+r.URL.String())
}

func addHandler(route string, handler http.HandlerFunc) {
	if _, ok := mux[route]; ok {
		panic("Route exists.")
	}
	mux[route] = handler
}
