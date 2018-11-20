package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// %v is string injection of the following value
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}
}
