package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Starting server at port 8000\n")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world from Go Server!")
	})
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi")
	})
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Error is :::", err)
	}
}
