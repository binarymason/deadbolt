package main

import (
	"fmt"
	"net/http"

	"github.com/binarymason/go-deadbolt/internal/routes"
)

func main() {
	http.HandleFunc("/", routes.Default)
	http.HandleFunc("/unlock", routes.Deadbolt)
	http.HandleFunc("/lock", routes.Deadbolt)

	port := ":8080"
	fmt.Println("listening on port", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
