package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Запуск сервера на http://127.0.0.1:4000")
	if err := http.ListenAndServe(":4000", mux); err != nil {
		log.Fatal(err)
	}
}
