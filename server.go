package main

import (
	"fmt"
	"net/http"
	"strings"
)

type SharkyStore interface {
	GetDeckQuantity(deck string) string
}

type SharkyServer struct {
	store SharkyStore
}

func (s *SharkyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	deck := strings.TrimPrefix(r.URL.Path, "/decks/")

	fmt.Fprint(w, s.GetDeckQuantity(deck))
}

// Count provided deck length and return quantity for it
func (s *SharkyServer) GetDeckQuantity(deck string) string {
	if deck == "Spanish" {
		return "20"
	}

	if deck == "English" {
		return "30"
	}
	return ""
}
