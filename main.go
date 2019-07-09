package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/binarymason/deadbolt/internal/config"
	"github.com/binarymason/deadbolt/internal/routes"
)

const DEADBOLT_VERSION = "201907081400"

var router routes.Router

// Load the deadbolt.yml config file.
// If a file path is not specified, defaults to /etc/deadbolt/deadbolt.yml
func init() {
	c := flag.String("c", "/etc/deadbolt/deadbolt.yml", "Specify deadbolt.yml file")
	flag.Parse()

	router = routes.Router{
		Version: DEADBOLT_VERSION,
		Config:  config.Load(*c),
	}
}

func main() {
	http.HandleFunc("/", router.Default)
	http.HandleFunc("/unlock", router.Deadbolt)
	http.HandleFunc("/lock", router.Deadbolt)

	port := router.Port()

	fmt.Println("listening on port", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
