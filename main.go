package main

import (
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Set the response status code to 502 Bad Gateway
	// w.WriteHeader(http.StatusBadGateway)
	// w.Write([]byte("<html><head><title>502 Bad Gateway</title></head><body><h1>502 Bad Gateway</h1></body></html>\n"))

	w.Write([]byte("<html><head><title>Cloud Run Demo</title></head><body><h1>Greetings from Cloud Run !!! YAY !!!</h1></body></html>\n"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
