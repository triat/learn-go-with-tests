package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("say 'Hello to people'", func(t *testing.T) {
		got := Hello("Tom", "")
		want := "Hello, Tom"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello world when no param is given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := Hello("Tom", "French")
		want := "Bonjour, Tom"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	// t.Helper() is needed to tell the test suite that this method is a helper.
	// By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
