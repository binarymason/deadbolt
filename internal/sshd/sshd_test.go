package sshd

import (
	"testing"

	. "github.com/binarymason/deadbolt/internal/testhelpers"
)

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
