package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubDeckStore struct {
	size map[string]int
}

func (s *StubDeckStore) GetDeckSize(deck string) int {
	size := s.size[deck]
	return size
}

func TestGETCards(t *testing.T) {
	store := StubDeckStore{
		size: map[string]int{
			"Spanish": 20,
			"English": 30,
		},
	}
	server := &SharkyServer{
		&store,
	}
	t.Run("returns quantity of cards for 'Spanish' deck", func(t *testing.T) {
		request := newGetDeckRequest("Spanish")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assertResponseBody(t, got, want)
	})

	t.Run("returns quantity of cards for 'English' deck", func(t *testing.T) {
		request := newGetDeckRequest("English")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

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
