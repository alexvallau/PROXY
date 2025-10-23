package server

import "net/http"

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	// Forwarding logic here
	w.Write([]byte("Request forwarded!"))
}
