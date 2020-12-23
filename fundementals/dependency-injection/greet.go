package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet says hi
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hi %s", name)
}

// MyGreeterHandler greets the www
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// Greet(os.Stdout, "Nicholas")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
