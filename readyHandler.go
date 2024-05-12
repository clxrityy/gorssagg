package main

import (
	"net/http"
)

func readyHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}