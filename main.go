package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const port string = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf("Server running on port '%s'\n", port)
	http.ListenAndServe(port, nil)
}
