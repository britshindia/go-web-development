package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>I'm an awesome Programmer ..!!! ^_^ :)</h1>")
	return
}

func main() {
	fmt.Println("Waiting to accept requests ...")

	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe("localhost:8080", nil)
	return
}
