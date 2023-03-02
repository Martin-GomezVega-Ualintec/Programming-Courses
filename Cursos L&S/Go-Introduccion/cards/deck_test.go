package main

import "testing"
import "strconv"
import "os"

// t tiene el valor si algo salió mal
func TestNewDeck(t *testing.T) {
	d := newDeck()
	var cadena string = "Expected deck length of 20, but got %v" + strconv.Itoa(len(d))
	if len(d) != 16 {
		t.Errorf(cadena)
	}

	// Primer carta
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// Ultima carta
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
