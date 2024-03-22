package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("saying hello to a person", func(t *testing.T) {
		got := Hello("Aksh", "")
		want := "Hello, Aksh"
		// fmt.Printf("%T\n", t)
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to anything", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Aksh", "Spanish")
		want := "Hola, Aksh"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Aksh", "French")
		want := "Bonjour, Aksh"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got string, want string) {
	// fmt.Printf("%T\n", t)
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
