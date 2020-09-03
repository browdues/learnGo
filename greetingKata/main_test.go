package main

import "testing"

func TestGreetBob(t *testing.T) {
	actualGreeting := Greet("Bob")
	expectedGreeting := "Hello, Bob."

	if actualGreeting != expectedGreeting {
		t.Errorf("Got: '%s', want '%s'", actualGreeting, expectedGreeting)
	}
}

func TestGreetEmpty(t *testing.T) {
	actualGreeting := Greet("")
	expectedGreeting := "Hello, my friend."

	if actualGreeting != expectedGreeting {
		t.Errorf("Got: '%s', want '%s'", actualGreeting, expectedGreeting)
	}
}

func TestGreetBOB(t *testing.T) {
	actualGreeting := Greet("BOB")
	expectedGreeting := "HELLO BOB!"

	if actualGreeting != expectedGreeting {
		t.Errorf("Got: '%s', want '%s'", actualGreeting, expectedGreeting)
	}
}

func TestGreetJillandJane(t *testing.T) {
	actualGreeting := Greet("Jill", "Jane")
	expectedGreeting := "Hello, Jill and Jane."

	if actualGreeting != expectedGreeting {
		t.Errorf("Got: '%s', want '%s'", actualGreeting, expectedGreeting)
	}
}

func TestGreetJillJaneAmy(t *testing.T) {
	actualGreeting := Greet("Jill", "Jane", "Amy")
	expectedGreeting := "Hello, Jill, Jane, and Amy."

	if actualGreeting != expectedGreeting {
		t.Errorf("Got: '%s', want '%s'", actualGreeting, expectedGreeting)
	}
}

func TestGreetJillJANEandAmy(t *testing.T) {
	actualGreeting := Greet("Jill", "JANE", "Amy")
	expectedGreeting := "Hello, Jill and Amy. AND HELLO JANE!"

	if actualGreeting != expectedGreeting {
		t.Errorf("Got: '%s', want '%s'", actualGreeting, expectedGreeting)
	}
}

func TestGreetAmyBRIANBobMAX(t *testing.T) {
	actualGreeting := Greet("Amy", "BRIAN", "Bob", "MAX")
	expectedGreeting := "Hello, Amy and Bob. AND HELLO BRIAN AND MAX!"

	if actualGreeting != expectedGreeting {
		t.Errorf("Got: '%s', want '%s'", actualGreeting, expectedGreeting)
	}
}

func TestGreetAmyBRIAN(t *testing.T) {
	actualGreeting := Greet("Amy", "BRIAN")
	expectedGreeting := "Hello, Amy. AND HELLO BRIAN!"

	if actualGreeting != expectedGreeting {
		t.Errorf("Got: '%s', want '%s'", actualGreeting, expectedGreeting)
	}
}
