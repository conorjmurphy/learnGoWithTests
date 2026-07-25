package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Conor")
	want := "Hello, Conor!"

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
