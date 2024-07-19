package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var postMutex sync.Mutex

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Printf("GET request,\nPath: %s\nHeaders:\n%s\n", r.URL.Path, r.Header)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "GET request for %s", r.URL.Path)
	case "POST":
		postMutex.Lock()
		defer postMutex.Unlock()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Cannot read body", http.StatusBadRequest)
			return
		}
		log.Printf("POST request,\nPath: %s\nHeaders:\n%s\n\nBody:\n%s\n", r.URL.Path, r.Header, body)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "POST request for %s", r.URL.Path)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handler)

	port := "80"
	log.Printf("Starting httpd on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
