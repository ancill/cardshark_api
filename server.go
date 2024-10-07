package main

import (
	"fmt"
	"net/http"
	"strings"
)

func SharkyServer(w http.ResponseWriter, r *http.Request) {
	deck := strings.TrimPrefix(r.URL.Path, "/decks/")

	fmt.Fprint(w, GetDeckQuantity(deck))
}

// Count provided deck length and return quantity for it
func GetDeckQuantity(deck string) string {
	if deck == "Spanish" {
		return "20"
	}

	if deck == "English" {
		return "30"
	}
	return ""
}
