package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func main() {
	server := &PlayerServer{store: &InMemoryPlayerStore{}}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

func GetPlayerScore(name string) int {
	if name == "Pepper" {
		return 20
	}

	if name == "Floyd" {
		return 10
	}

	return 0
}
