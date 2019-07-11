package main

import (
	"flag"

	"github.com/binarymason/deadbolt/internal/deadbolt"
)

var config *string

func init() {
	config = flag.String("c", "/etc/deadbolt/deadbolt.yml", "Specify deadbolt.yml file")
	flag.Parse()
}

func main() {
	dblt := deadbolt.New(*config)
	if err := dblt.Listen(); err != nil {
		panic(err)
	}
}
