package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	err := http.ListenAndServe(":8000", http.HandlerFunc(MyGreeterHandler))

	if err != nil {
		fmt.Println(err)
	}
}
