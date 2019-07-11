package deadbolt

import (
	"testing"

	. "github.com/binarymason/deadbolt/internal/testhelpers"
)

func TestValidRequest(t *testing.T) {
	var (
		r request
	)
	d := Deadbolt{Secret: "foo", Whitelisted: []string{"127.0.0.3"}}

	Given("a remote IP")
	And("an Authorization Header")
	When("IP is whitelisted")
	And("Authorization is correct")
	Then("valid request")
	r = request{ip: "127.0.0.3", auth: "foo"}
	Assert(d.authorizedRequest(&r), true, t)

	When("IP is NOT whitelisted")
	And("Authorization is correct")
	Then("invalid request")
	r = request{ip: "nope", auth: "foo"}
	Assert(d.authorizedRequest(&r), false, t)

	When("IP is whitelisted")
	And("Authorization is NOT correct")
	Then("invalid request")
	r = request{ip: "127.0.0.3", auth: "nope"}
	Assert(d.authorizedRequest(&r), false, t)

	When("IP is NOT whitelisted")
	And("Authorization is NOT correct")
	Then("invalid request")
	r = request{ip: "nope", auth: "nope"}
	Assert(d.authorizedRequest(&r), false, t)
}
