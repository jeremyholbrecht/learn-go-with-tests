package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jere", "")
		want := "Hello, Jere"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Jere", spanish)
		want := "Hola, Jere"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jere", french)
		want := "Bonjour, Jere"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In dutch", func(t *testing.T) {
		got := Hello("Jere", dutch)
		want := "Goeiendag, Jere"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
