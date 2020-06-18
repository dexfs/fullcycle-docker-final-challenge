package main

import "testing"

func TestGreeting(t *testing.T) {
	result := greeting("test")
	if result != "<b>test</b>" {
		t.Error("Opa! o resultado deveria ser <b>test</b>", result)
	}

	emptyResult := greeting("")
	if emptyResult != "<b>Code.education Rocks!<b>" {
		t.Error("Opa! o resultado deveria ser <b>test</b>", emptyResult)
	}
}
