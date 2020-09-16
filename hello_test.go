package main

import "testing"

func TestGreeting(t *testing.T) {
	msg, err := greeting("smite")
	if msg != "Hi smite" || err != nil {
		t.Fatalf("msg mismatch")
	}
}
