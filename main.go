package main

import (
	"fmt"
	"net/http"

	"github.com/binarymason/go-deadbolt/internal/config"
	"github.com/binarymason/go-deadbolt/internal/routes"
)

const DEADBOLT_VERSION = "201907081400"

func main() {
	router := routes.Router{
		Version: DEADBOLT_VERSION,
		Config:  config.Load("./testdata/simple_deadbolt_config.yml"),
	}

	http.HandleFunc("/", router.Default)
	http.HandleFunc("/unlock", router.Deadbolt)
	http.HandleFunc("/lock", router.Deadbolt)

	port := router.Port()

	fmt.Println("listening on port", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
