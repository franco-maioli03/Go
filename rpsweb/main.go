package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {
	// Create Handler
	router := http.NewServeMux()

	// Handler to serve static files
	fs := http.FileServer(http.Dir("static"))
	// Path to access static files
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	// Configure routes

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	log.Printf("Server listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
