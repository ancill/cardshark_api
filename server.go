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
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	deck := strings.TrimPrefix(r.URL.Path, "/decks/")

	score := s.store.GetDeckSize(deck)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
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
