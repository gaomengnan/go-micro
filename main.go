package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		log.Println("test")
		fmt.Fprintf(w, "hello")
	})
	log.Println("web server started: " + port)
	http.ListenAndServe(":"+port, nil)

}
