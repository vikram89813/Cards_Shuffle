package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	l := len(d)

	if l != 52 {
		t.Errorf("Expected deck length of 52 but got : %v", l)
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of spades but got : %v", d[0])
	}

	if d[l-1] != "king of Clubs" {
		t.Errorf("Expected first card to be king of Clubs but got : %v", d[l-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52 but got : %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
