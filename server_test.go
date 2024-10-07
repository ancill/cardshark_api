package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETCards(t *testing.T) {
	t.Run("returns quantity of cards for 'Spanish' deck", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/decks/Spanish", nil)
		response := httptest.NewRecorder()

		SharkyServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns quantity of cards for 'English' deck", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/decks/English", nil)
		response := httptest.NewRecorder()

		SharkyServer(response, request)

		got := response.Body.String()
		want := "30"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
