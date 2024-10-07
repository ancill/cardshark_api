package main

import (
	"fmt"
	"net/http"
	"strings"
)

func SharkyServer(w http.ResponseWriter, r *http.Request) {
	deck := strings.TrimPrefix(r.URL.Path, "/decks/")

	if deck == "Spanish" {
		fmt.Fprint(w, "20")
	}

	if deck == "English" {
		fmt.Fprint(w, "30")
	}
}
