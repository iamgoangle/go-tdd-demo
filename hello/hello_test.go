package hello

import "testing"

func TestHello(t *testing.T) {
	assert := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("it should say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assert(t, got, want)
	})

	t.Run("it should say 'Hello, Golf'", func(t *testing.T) {
		got := Hello("Golf", "")
		want := "Hello, Golf"

		assert(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Golf", "spanish")
		want := "Hola, Golf"
		assert(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Golf", "french")
		want := "Bonjour, Golf"
		assert(t, got, want)
	})
}
