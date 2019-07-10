package main

import (
	"flag"

	"github.com/binarymason/deadbolt/internal/server"
)

var s *server.Server

// Load the deadbolt.yml config file.
// If a file path is not specified, defaults to /etc/deadbolt/deadbolt.yml
func init() {
	c := flag.String("c", "/etc/deadbolt/deadbolt.yml", "Specify deadbolt.yml file")
	flag.Parse()
	s = server.New(c)
}

func main() {
	if err := s.Serve(); err != nil {
		panic(err)
	}
}
