package select_racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	defer slowServer.Close()
	defer fastServer.Close()

	want := fastServer.URL
	got := Racer(slowServer.URL, fastServer.URL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
