package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("Saying hello to someone (an argument is passed in.", func(t *testing.T) {
		got := Hello("Bub")
		want := "Hello, Bub"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to everyone (empty string passed in).", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!!!"
		assertCorrectMessage(t, got, want)
	})
}
