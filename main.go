package main

import (
	"log"
	"net/http"
)

func main() {
	server := &SharkyServer{}
	handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
