package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}


func PlayerServer(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "20")
}
