package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Says hello to people", func(t *testing.T) {
		got := Hello("Advaiya", "")
		want := "Hello, Advaiya!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Gives greetings in Spanish", func(t *testing.T) {
		got := Hello("Advaiya", "Spanish")
		want := "Hola, Advaiya!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Gives greetings in French", func(t *testing.T) {
		got := Hello("Advaiya", "French")
		want := "Bonjour, Advaiya!"
		assertCorrectMessage(t, got, want)
	})
}
