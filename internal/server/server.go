package server

import (
	"net/http"
)

func Start() {
	http.HandleFunc("/", proxyHandler)
	http.HandleFunc("/health", healthCheck)
	http.ListenAndServe(":8080", nil)
}
