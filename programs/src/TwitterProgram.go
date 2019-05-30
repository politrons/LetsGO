package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, welcome  %s!", r.URL.Path[1:])

}

func main() {
	http.HandleFunc("/tweet/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
