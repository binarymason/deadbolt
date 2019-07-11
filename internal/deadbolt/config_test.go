package deadbolt

import (
	"os"
	"testing"

	. "github.com/binarymason/deadbolt/internal/testhelpers"
)

func TestLoad(t *testing.T) {
	var (
		p string
		r *Deadbolt
	)

	p = "../../test/testdata/simple_deadbolt_config.yml"

	Given("a deadbolt config file")
	Then("values are parsed correctly")

	r, _ = New(p)

	Assert(r.Secret, "supersecret", t)
	Assert(r.Whitelisted[0], "127.0.0.1", t)
	Assert(r.Whitelisted[1], "127.0.0.2", t)

	And("SSHDConfigPath defaults to /etc/ssh/sshd_config")
	Assert(r.SSHDConfigPath, "/etc/ssh/sshd_config", t)

	When("DEADBOLT_SECRET is an environment variable")
	Then("environment variable takes precedence")

	os.Setenv("DEADBOLT_SECRET", "foo")
	r, _ = New(p)
	Assert(r.Secret, "foo", t)
	os.Setenv("DEADBOLT_SECRET", "") // teardown

	When("deadbolt_secret is NOT in config file")
	And("DEADBOLT_SECRET is an environment variable")
	p = "../../test/testdata/missing_secret_deadbolt_config.yml"
	os.Setenv("DEADBOLT_SECRET", "bar")
	r, _ = New(p)
	Assert(r.Secret, "bar", t)
	os.Setenv("DEADBOLT_SECRET", "") // teardown

	When("deadbolt_secret is NOT in config file")
	And("DEADBOLT_SECRET is NOT an environment variable")
	Then("an error is returned")
	_, err := New(p)

	if err == nil {
		t.Fatal("expected an error, but received none")
	}
}
