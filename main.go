package main

import (
	"log"
	"net/http"
)

type InMemoryDeckStore struct{}

func (i *InMemoryDeckStore) GetDeckSize(deck string) int {
	return 123
}

func main() {
	server := &SharkyServer{&InMemoryDeckStore{}}
	handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5555", handler))
}
