package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/home", homeHandler)
	mux.HandleFunc("/profile", profileHandler)

	log.Println("Starting web on 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Home"))
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Profile"))
}
