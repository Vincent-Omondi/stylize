// main.go
package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Vincent-Omondi/stylize/handlers"
)

func main() {
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Use the handler function for routing
	http.HandleFunc("/", handler)

	log.Println("Server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

// handler function for routing HTTP requests
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handlers.IndexHandler(w, r)
	case "/ascii-art":
		handlers.AsciiArtHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}
