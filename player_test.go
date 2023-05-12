package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("it can get palyer", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players", nil)
		response := httptest.NewRecorder()
		PlayerServer(response, request)
		got := response.Body.String()
		want := "30"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
