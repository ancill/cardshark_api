package main

import (
	"log"
	"net/http"
)

func main() {
	server := &SharkyServer{NewInMemoryDeckStore()}
	log.Fatal(http.ListenAndServe(":5555", server))
}
