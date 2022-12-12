package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	Host = "0.0.0.0"
	Port = "8080"
)

func main() {
	fmt.Println("[+] The application has started")
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Access-Control-Allow-Origin", "*")
                w.Header().Set("Content-Type", "application/json")
		
	        fmt.Fprintf(w, "This is a Simple HTTP Web Server!")
	})

	err := http.ListenAndServe(Host+":"+Port, nil)
	if err != nil {
		log.Fatal("Error Starting the HTTP Server : ", err)
		return
	}
}
