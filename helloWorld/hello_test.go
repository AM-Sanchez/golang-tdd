package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Bub")
	want := "Hello Bub"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
