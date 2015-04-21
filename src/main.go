package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleHookRequest)
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
