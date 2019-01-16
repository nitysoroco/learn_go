package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Whoa! Go is neat!!")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is my about")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)
}
