package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Gambôa", "Spanish")
		want := "Hola, Gambôa"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Gambôa", "French")
		want := "Bonjour, Gambôa"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	//tell that method is a helper, and error will indicate the line that was called in the test
	//ex: error is not on line 26, but in line 17 or 10
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
