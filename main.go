package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/health-check", healthCheck)
	http.ListenAndServe(":9290", nil)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
