package main

import "testing"

func TestGreetings(t *testing.T) {
	t.Run("test greet with person", func(t *testing.T) {
		result := greetings("testing")
		expected := "Hello, testing"
		if result != expected {
			t.Errorf("got result: %s, expected %s", result, expected)
		}
	})

	t.Run("test greet empty person says World", func(t *testing.T) {
		result := greetings("")
		expected := "Hello, World"
		if result != expected {
			t.Errorf("got result: %s, expected %s", result, expected)
		}
	})
}
