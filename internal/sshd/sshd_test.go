package sshd

import (
	"fmt"
	"testing"

	. "github.com/binarymason/deadbolt/internal/testhelpers"
)

func TestPermitRootLogin(t *testing.T) {
	var (
		r string
		x string
		m string
	)

	Given("without-password")
	When("PermitRootLogin is currently set to yes")
	Then("it should return 'PermitRootLogin without-password'")
	m = "without-password"
	x = fmt.Sprintf("PermitRootLogin %s", m)
	r = updatePermitRootLogin(m, "PermitRootLogin yes")
	Assert(r, x, t)

	When("PermitRootLogin is currently set to yes")
	And("it is commented out")
	Then("it should return 'PermitRootLogin without-password'")
	r = updatePermitRootLogin(m, "#PermitRootLogin no")
	Assert(r, x, t)

	When("PermitRootLogin is currently set to no")
	Then("it should return 'PermitRootLogin without-password'")
	r = updatePermitRootLogin(m, "PermitRootLogin no")
	Assert(r, x, t)

	When("PermitRootLogin is currently set to no")
	And("it is commented out")
	Then("it should return 'PermitRootLogin without-password'")
	r = updatePermitRootLogin(m, "#PermitRootLogin no")
	Assert(r, x, t)

	Given("no")
	When("PermitRootLogin is currently set to yes")
	Then("it should return 'PermitRootLogin no'")
	m = "no"
	x = fmt.Sprintf("PermitRootLogin %s", m)
	r = updatePermitRootLogin(m, "PermitRootLogin yes")
	Assert(r, x, t)

	When("PermitRootLogin is currently set to yes")
	And("it is commented out")
	Then("it should return 'PermitRootLogin no'")
	r = updatePermitRootLogin(m, "#PermitRootLogin yes")
	Assert(r, x, t)

	When("PermitRootLogin is currently set to no")
	Then("it should return 'PermitRootLogin no'")
	r = updatePermitRootLogin(m, "PermitRootLogin no")
	Assert(r, x, t)

	When("PermitRootLogin is currently set to no")
	And("it is commented out")
	Then("it should return 'PermitRootLogin no'")
	r = updatePermitRootLogin(m, "#PermitRootLogin no")
	Assert(r, x, t)
}

func TestGenerateConfig(t *testing.T) {
	var (
		cfg string
		r   string
		x   string
		m   string
	)

	Given("'no' and lines of text")
	m = "no"

	When("PermitRootLogin is currently yes")
	Then("it should return lines of text")
	And("PermitRootLogin should be set to no")

	cfg = `
foo
bar
baz bang
PermitRootLogin yes
booger
`

	x = `
foo
bar
baz bang
PermitRootLogin no
booger
`

	r = generateConfig(m, cfg)
	Assert(r, x, t)

	When("PermitRootLogin is already no")
	Then("it should return lines of text")
	And("PermitRootLogin should be set to no")
	cfg = `
foo
bar
baz bang
PermitRootLogin yes
booger
`
	r = generateConfig(m, cfg)
	Assert(r, x, t)

	Given("'without-password' and lines of text")
	m = "without-password"

	When("PermitRootLogin is no")
	Then("it should return lines of text")
	And("PermitRootLogin should be set to without-password")

	cfg = `
foo
bar
baz bang
PermitRootLogin no
booger
`

	x = `
foo
bar
baz bang
PermitRootLogin without-password
booger
`
	r = generateConfig(m, cfg)
	Assert(r, x, t)

	When("PermitRootLogin is already without-password")
	Then("it should return lines of text")
	And("PermitRootLogin should be set to without-password")

	cfg = `
foo
bar
baz bang
PermitRootLogin without-password
booger
`
	r = generateConfig(m, cfg)
	Assert(r, x, t)
}
