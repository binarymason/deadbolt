package sshd

import (
	"fmt"
	"testing"
)

func Given(s string) {
	fmt.Println("Given", s)
}

func When(s string) {
	fmt.Println("  When", s)
}

func Then(s string) {
	fmt.Println("    Then", s)
}

func And(s string) {
	fmt.Println("    And", s)
}

func Assert(a, x string, t *testing.T) {

	a = fmt.Sprintf("%v", a)
	x = fmt.Sprintf("%v", x)
	if a != x {
		t.Errorf("Expected %s, but got: %s", x, a)
	}
}

func TestToggle(t *testing.T) {
	var (
		r string
		x string
		s string
	)

	Given("PermitRootLogin yes")
	When("locking")
	Then("it should return 'PermitRootLogin no'")
	r = toggle("lock", "PermitRootLogin yes")
	x = "PermitRootLogin no"
	Assert(r, x, t)

	Given("PermitRootLogin no")
	When("locking")
	Then("it should return 'PermitRootLogin no'")
	r = toggle("lock", "PermitRootLogin no")
	x = "PermitRootLogin no"
	Assert(r, x, t)

	Given("a random string")
	When("locking")
	Then("it should return the string")
	s = "foobarbaz"
	r = toggle("lock", s)
	Assert(r, s, t)

	And("it works with empty strings")
	s = ""
	r = toggle("lock", s)
	Assert(r, s, t)

	Given("PermitRootLogin no")
	When("unlocking")
	Then("it should return 'PermitRootLogin yes'")
	r = toggle("unlock", "PermitRootLogin no")
	x = "PermitRootLogin yes"
	Assert(r, x, t)

	Given("PermitRootLogin yes")
	When("unlocking")
	Then("it should return 'PermitRootLogin yes'")
	r = toggle("unlock", "PermitRootLogin yes")
	x = "PermitRootLogin yes"
	Assert(r, x, t)
}

func TestLockConfig(t *testing.T) {
	var (
		cfg string
		r   string
		x   string
	)

	Given("lines of text")
	When("locking")
	And("PermitRootLogin is yes")
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
	r = lockConfig(cfg)
	Assert(r, x, t)

	When("locking")
	And("PermitRootLogin is already no")
	Then("it should return lines of text")
	And("PermitRootLogin should be set to no")
	cfg = `
foo
bar
baz bang
PermitRootLogin yes
booger
  `
	r = lockConfig(cfg)
	Assert(r, x, t)

}

func TestUnlockConfig(t *testing.T) {
	var (
		cfg string
		r   string
		x   string
	)

	Given("lines of text")
	When("unlocking")
	And("PermitRootLogin is no")
	Then("it should return lines of text")
	And("PermitRootLogin should be set to yes")

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
PermitRootLogin yes
booger
  `
	r = unlockConfig(cfg)
	Assert(r, x, t)

	When("unlocking")
	And("PermitRootLogin is already yes")
	Then("it should return lines of text")
	And("PermitRootLogin should be set to yes")

	cfg = `
foo
bar
baz bang
PermitRootLogin yes
booger
  `
	r = unlockConfig(cfg)
	Assert(r, x, t)
}
