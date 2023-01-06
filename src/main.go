package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
