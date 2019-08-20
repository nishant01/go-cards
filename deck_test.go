package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of the Spades" {
		t.Errorf("Expected first card of Ace of the Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of the Clubs" {
		t.Errorf("Expected first card of Four of the Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.toSaveFile("_decktesting")
	lodedDeck := newDeckFromFile("_decktesting")
	if len(lodedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(lodedDeck))
	}
	os.Remove("_decktesting")

}
