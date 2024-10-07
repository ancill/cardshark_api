package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETCards(t *testing.T) {
	t.Run("returns quantity of cards for 'Spanish' deck", func(t *testing.T) {
		request := newGetDeckRequest("Spanish")
		response := httptest.NewRecorder()

		SharkyServer(response, request)

		got := response.Body.String()
		want := "20"

		assertResponseBody(t, got, want)
	})

	t.Run("returns quantity of cards for 'English' deck", func(t *testing.T) {
		request := newGetDeckRequest("English")
		response := httptest.NewRecorder()

		SharkyServer(response, request)

		got := response.Body.String()
		want := "30"

		assertResponseBody(t, got, want)
	})
}

func newGetDeckRequest(deck string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/decks/%s", deck), nil)
	return request
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}
