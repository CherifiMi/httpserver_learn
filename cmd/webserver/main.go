package main

import (
	poker "httpserver"
	"log"
	"net/http"
)

func main() {
	server := poker.NewPlayerServer(poker.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
