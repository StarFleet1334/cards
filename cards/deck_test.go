package main

import "testing"

func TestDeckCreatioon(t *testing.T) {
	deck := createDeck()

	if len(deck) != 16 {
		t.Errorf("Expected length of deck to be 16,but got %v", len(deck))
	}
}
