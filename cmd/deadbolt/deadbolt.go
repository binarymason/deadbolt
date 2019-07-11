package main

import (
	"flag"
	"log"
	"os"

	"github.com/binarymason/deadbolt/internal/deadbolt"
)

var config *string

func init() {
	config = flag.String("c", "/etc/deadbolt/deadbolt.yml", "Specify deadbolt.yml file")
	flag.Parse()
}

func main() {
	dblt, err := deadbolt.New(*config)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	dblt.Listen()
}
