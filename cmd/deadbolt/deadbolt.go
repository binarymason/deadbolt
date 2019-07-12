package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/binarymason/deadbolt/internal/deadbolt"
)

var config *string
var showVersion *bool

func init() {
	config = flag.String("c", "/etc/deadbolt/deadbolt.yml", "Specify deadbolt.yml file")
	showVersion = flag.Bool("version", false, "Print deadbolt version")

	flag.Parse()
}

func main() {
	if *showVersion {
		fmt.Println(deadbolt.GetVersion())
		return
	}

	dblt, err := deadbolt.New(*config)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	dblt.Listen()
}
