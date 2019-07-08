package main

import (
	"fmt"
	"net/http"
)

const DEADBOLT_VERSION = "201907081400"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path != "/" {
		http.Error(w, "invalid route", http.StatusNotFound)
		return
	}
	w.Write([]byte("Deadbolt version: " + DEADBOLT_VERSION + "\n"))
}

func deadboltHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if r.Method != http.MethodPost {
		http.Error(w, "invalid HTTP method. expected POST", http.StatusBadRequest)
		return
	}
	w.Write([]byte(path + "\n"))
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/unlock", deadboltHandler)
	http.HandleFunc("/lock", deadboltHandler)

	port := ":8080"
	fmt.Println("listening on port", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
