package main

import (
	"fmt"
	"net/http"
	"strings"
)

type SharkyStore interface {
	// Counting provided deck length and return size
	GetDeckSize(deck string) int
	RecordDeck(deck string)
}

type SharkyServer struct {
	store SharkyStore
}

func (s *SharkyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	deck := strings.TrimPrefix(r.URL.Path, "/decks/")
	switch r.Method {
	case http.MethodPost:
		s.processCard(w, deck)
	case http.MethodGet:
		s.showSize(w, deck)
	}

}

func (s *SharkyServer) showSize(w http.ResponseWriter, deck string) {
	score := s.store.GetDeckSize(deck)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (s *SharkyServer) processCard(w http.ResponseWriter, deck string) {
	s.store.RecordDeck(deck)
	w.WriteHeader(http.StatusAccepted)
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
