package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/verlandz/elasticsearch/scripts/handler"
)

func ignore(w http.ResponseWriter, r *http.Request) {}

func main() {
	http.HandleFunc("/search", handler.Search)
	http.HandleFunc("/favicon.ico", ignore)

	port := ":8080"
	fmt.Println("Listen and serve", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
