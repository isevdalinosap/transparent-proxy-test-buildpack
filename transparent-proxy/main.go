package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Im the transparent proxy!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	log.Printf("Transparent proxy server starting on port 8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
