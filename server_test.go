package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubDeckStore struct {
	size  map[string]int
	decks []string
}

func (s *StubDeckStore) GetDeckSize(deck string) int {
	size := s.size[deck]
	return size
}

func (s *StubDeckStore) RecordDeck(deck string) {
	s.decks = append(s.decks, deck)
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
	t.Run("returns 404 on missing deck", func(t *testing.T) {
		request := newGetDeckRequest("Chinise")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestDeckStore(t *testing.T) {
	store := StubDeckStore{
		map[string]int{},
		nil,
	}

	server := &SharkyServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		deck := "Spanish"
		request := newPostDeckRequest(deck)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.decks) != 1 {
			t.Errorf("got %d calls to RecordDeck want %d", len(store.decks), 1)
		}

		if store.decks[0] != deck {
			t.Errorf("did not store correct deck got %q want %q", store.decks[0], deck)
		}
	})
}

func newPostDeckRequest(deck string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/decks/%s", deck), nil)
	return req
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
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
