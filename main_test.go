package main

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
