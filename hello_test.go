package main

import "testing"

func TestHello(t *testing.T) {
    assertCorrectMessage := func(t testing.TB, got, want string) {
        t.Helper() // if error occurs, error comes from subtest not this helper function 
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

    t.Run("Saying hello to people", func(t *testing.T) {
        got := Hello("Sam", "")
        want := "Hello, Sam"
        assertCorrectMessage(t, got, want)
    })

    t.Run("Empty name string defaults to 'World'", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

    t.Run("Unrecognized language defaults to English", func (t *testing.T) {
        got := Hello("Mayowa", "German")
        want := "Hello, Mayowa"
        assertCorrectMessage(t, got, want)
    })

    t.Run("Check if language of french returns correct output", func (t *testing.T) {
        got := Hello("Jake", "french")
        want := "Bonjour, Jake"
        assertCorrectMessage(t, got, want)
    })
}
