package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	totalCards := 16

	if len(d) != totalCards {
		t.Errorf("Expected deck length of %v, but got %v", totalCards, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFile := "_decktesting"
	os.Remove(testFile)

	deck := newDeck()
	deck.saveToFile(testFile)

	loadedDeck := newDeckFromFile(testFile)

	totalCards := 16
	if len(loadedDeck) != totalCards {
		t.Errorf("Expected %v cards, got %v", totalCards, len(loadedDeck))
	}

	os.Remove(testFile)
}
