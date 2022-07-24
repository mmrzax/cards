package cards

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of 'Ace of Spades', but got '%v'", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of 'Four of Clubs', but got '%v'", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {

	// Remove existing "_decktesting"
	os.Remove("_decktesting")

	// Create a new deck
	deck := NewDeck()

	// Save deck to "_decktesting"
	deck.SaveToFile("_decktesting")

	// Load deck from "_decktesting"
	loadedDeck := NewDeckFromFile("_decktesting")

	// Assert deck Length
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	// Remove _decktesting
	os.Remove("_decktesting")
}
