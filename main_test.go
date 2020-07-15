package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("should say hello to people", func(t *testing.T) {
		got := Hello("Yok", "")
		want := "Hello, Yok"

		assertCorrectMessage(t, got, want)
	})

	t.Run("should say hello world, if name empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("should hello in Spanish", func(t *testing.T) {
		got := Hello("Elodia", "Spanish")
		want := "Hola, Elodia"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should hello in French", func(t *testing.T) {
		got := Hello("Cal", "French")
		want := "Bonjour, Cal"
		assertCorrectMessage(t, got, want)
	})
}
