package main

import (
	"log"
	"net/http"
	"os"
)

func getPort() string {
	port := os.Getenv("PORT")

	if port != "" {
		return ":" + port
	}

	return ":80"
}

func main() {
	http.HandleFunc("/", handleHookRequest)

	port := getPort()
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
