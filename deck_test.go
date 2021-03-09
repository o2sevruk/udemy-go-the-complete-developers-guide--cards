package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	qty := 16

	if len(d) != qty {
		t.Errorf("Expected deck length %v, but it %v", qty, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	qty := 16

	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	lodadedDeck := newDeckFromFile(filename)

	if len(lodadedDeck) != qty {
		t.Errorf("Expected deck length %v, but it %v", qty, len(lodadedDeck))
	}

	os.Remove(filename)
}
