package main

//import "os"
import "net/http"
import "log"
//import "./hookhandler"

func main() {
	http.HandleFunc("/", HookHandler)
	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
