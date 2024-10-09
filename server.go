package main

import (
	"fmt"
	"net/http"
	"strings"
)

type SharkyStore interface {
	// Counting provided deck length and return size
	GetDeckSize(deck string) int
}

type SharkyServer struct {
	store SharkyStore
}

func (s *SharkyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	deck := strings.TrimPrefix(r.URL.Path, "/decks/")

	fmt.Fprint(w, s.store.GetDeckSize(deck))
}

func GetDeckSize(deck string) int {
	if deck == "Spanish" {
		return 20
	}

	if deck == "English" {
		return 30
	}
	return 0
}
