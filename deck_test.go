package main

import (
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rand.Seed(1)
	os.Exit(m.Run())
}

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0].toString() != "A\u2660" {
		t.Errorf("Expected first card to be A\u2660, but got %s", d[0].toString())
	}

	if d[len(d)-1].toString() != "K\u2663" {
		t.Errorf("Expected last card to be K\u2663, but got %s", d[len(d)-1].toString())
	}
}

func TestSaveDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected loaded deck length of 52, but got %d", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestDeal(t *testing.T) {
	d := newDeck()
	hand, remainder := deal(d, 5)

	if len(hand) != 5 {
		t.Errorf("Expected hand length of 5, but got %d", len(hand))
	}

	if len(remainder) != 47 {
		t.Errorf("Expected remainder length of 47, but got %d", len(remainder))
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()

	d.shuffle()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0].toString() != "K\u2666" {
		t.Errorf("Expected first card to be K\u2666, but got %s", d[0].toString())
	}

	if d[len(d)-1].toString() != "6\u2666" {
		t.Errorf("Expected last card to be 6\u2666, but got %s", d[len(d)-1].toString())
	}
}

func TestToString(t *testing.T) {
	d := newDeck()

	s := d.toString()

	if len(s) != 263 {
		t.Errorf("Expected string representation length of 263, but got %d", len(s))
	}
}
