package config

import (
	"os"
	"testing"

	. "github.com/binarymason/go-deadbolt/internal/testhelpers"
)

func TestLoad(t *testing.T) {
	var (
		p string
		r Config
	)

	p = "../../testdata/simple_deadbolt_config.yml"

	Given("a deadbolt config file")
	Then("values are parsed correctly")

	r = Load(p)

	Assert(r.Secret, "supersecret", t)
	Assert(r.Whitelisted[0], "127.0.0.1", t)
	Assert(r.Whitelisted[1], "127.0.0.2", t)

	When("DEADBOLT_SECRET is an environment variable")
	Then("environment variable takes precedence")

	os.Setenv("DEADBOLT_SECRET", "foo")
	r = Load(p)
	Assert(r.Secret, "foo", t)
	os.Setenv("DEADBOLT_SECRET", "") // teardown

	When("deadbolt_secret is NOT in config file")
	And("DEADBOLT_SECRET is an environment variable")
	p = "../../testdata/missing_secret_deadbolt_config.yml"
	os.Setenv("DEADBOLT_SECRET", "bar")
	r = Load(p)
	Assert(r.Secret, "bar", t)
	os.Setenv("DEADBOLT_SECRET", "") // teardown

	When("deadbolt_secret is NOT in config file")
	And("DEADBOLT_SECRET is NOT an environment variable")
	Then("panic")
	defer func() {
		if recover() == nil {
			t.Fatal("expected a panic, but received none")
		}
	}()

	r = Load(p)
}
