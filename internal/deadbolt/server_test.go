package deadbolt

import (
	"testing"

	. "github.com/binarymason/deadbolt/internal/testhelpers"
	"github.com/go-delve/delve/pkg/config"
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
	r = Router{Config: &c}

	Then("Port with colon is specified")
	Assert(r.Port(), ":"+port, t)

	When("No port is specified in config")
	c = config.Config{}
	r = Router{Config: &c}

	Then("Port defailts to 8080")
	Assert(r.Port(), ":8080", t)
}

func TestValidRequest(t *testing.T) {
	var (
		r request
	)
	c := config.Config{Secret: "foo", Whitelisted: []string{"127.0.0.3"}}

	Given("a remote IP")
	And("an Authorization Header")
	When("IP is whitelisted")
	And("Authorization is correct")
	Then("valid request")
	r = request{ip: "127.0.0.3", auth: "foo"}
	Assert(r.isValid(&c), true, t)

	When("IP is NOT whitelisted")
	And("Authorization is correct")
	Then("invalid request")
	r = request{ip: "nope", auth: "foo"}
	Assert(r.isValid(&c), false, t)

	When("IP is whitelisted")
	And("Authorization is NOT correct")
	Then("invalid request")
	r = request{ip: "127.0.0.3", auth: "nope"}
	Assert(r.isValid(&c), false, t)

	When("IP is NOT whitelisted")
	And("Authorization is NOT correct")
	Then("invalid request")
	r = request{ip: "nope", auth: "nope"}
	Assert(r.isValid(&c), false, t)
}
