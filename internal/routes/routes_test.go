package routes

import (
	"testing"

	"github.com/binarymason/deadbolt/internal/config"
	. "github.com/binarymason/deadbolt/internal/testhelpers"
)

func TestPort(t *testing.T) {
	var (
		c config.Config
		r Router
	)
	Given("a Deadbolt config")
	When("Port is specified")
	port := "8090"
	c = config.Config{Port: port}
	r = Router{Config: c}

	Then("Port with colon is specified")
	Assert(r.Port(), ":"+port, t)

	When("No port is specified in config")
	c = config.Config{}
	r = Router{Config: c}

	Then("Port defailts to 8080")
	Assert(r.Port(), ":8080", t)
}
