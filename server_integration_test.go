package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryDeckStore()
	server := SharkyServer{store}
	deck := "English"

	server.ServeHTTP(httptest.NewRecorder(), newPostDeckRequest(deck))
	server.ServeHTTP(httptest.NewRecorder(), newPostDeckRequest(deck))
	server.ServeHTTP(httptest.NewRecorder(), newPostDeckRequest(deck))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetDeckRequest(deck))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}
