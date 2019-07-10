package validate

import (
	"testing"

	"github.com/binarymason/deadbolt/internal/config"
	. "github.com/binarymason/deadbolt/internal/testhelpers"
)

func TestValidRequest(t *testing.T) {
	c := config.Config{Secret: "foo", Whitelisted: []string{"127.0.0.3"}}

	Given("a remote IP")
	And("an Authorization Header")
	When("IP is whitelisted")
	And("Authorization is correct")
	Then("valid request")
	Assert(ValidRequest("127.0.0.3", "foo", c), true, t)

	When("IP is NOT whitelisted")
	And("Authorization is correct")
	Then("invalid request")
	Assert(ValidRequest("nope", "foo", c), false, t)

	When("IP is whitelisted")
	And("Authorization is NOT correct")
	Then("invalid request")
	Assert(ValidRequest("127.0.0.3", "nope", c), false, t)

	When("IP is NOT whitelisted")
	And("Authorization is NOT correct")
	Then("invalid request")
	Assert(ValidRequest("nope", "nope", c), false, t)
}
