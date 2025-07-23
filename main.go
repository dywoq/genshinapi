package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("/static/"))

	// for debug purposes
	// fmt.Println("Starting server at 8080 port")

	http.HandleFunc("/", handle)
	http.HandleFunc("/characters/", handleCharacters)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
