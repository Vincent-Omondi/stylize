// main.go
package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"stylize/handlers"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Usage: go run .")
		os.Exit(1)
	}

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Use the handler function for routing
	http.HandleFunc("/", handler)
	port := ":8080"
	log.Printf("Server started on http://localhost%s", port)

	err := http.ListenAndServe(port, nil)
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
		handlers.ErrorHandler(w, r, http.StatusNotFound, []string{"Page not found"})
	}
}
