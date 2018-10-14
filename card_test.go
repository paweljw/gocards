package main

import "testing"

func TestCardToString(t *testing.T) {
	c := card{suit: "\u2660", value: 1}

	if c.toString() != "A\u2660" {
		t.Errorf("Expected string representation to be A\u2660, got %s", c.toString())
	}
}

func TestHumanValue(t *testing.T) {
	c := card{suit: "F", value: 1}

	if c.humanValue() != "A" {
		t.Errorf("Expected string representation of 1 to be A, got %s", c.humanValue())
	}

	c.value = 11

	if c.humanValue() != "J" {
		t.Errorf("Expected string representation of 11 to be J, got %s", c.humanValue())
	}

	c.value = 12

	if c.humanValue() != "Q" {
		t.Errorf("Expected string representation of 12 to be Q, got %s", c.humanValue())
	}

	c.value = 13

	if c.humanValue() != "K" {
		t.Errorf("Expected string representation of 13 to be K, got %s", c.humanValue())
	}

	c.value = 6

	if c.humanValue() != "6" {
		t.Errorf("Expected string representation of 6 to be \"6\", got %s", c.humanValue())
	}
}

func TestCardNumericValue(t *testing.T) {
	if cardNumericValue("A") != 1 {
		t.Errorf("Expected numeric representation of A to be 1, got %d", cardNumericValue("A"))
	}

	if cardNumericValue("J") != 11 {
		t.Errorf("Expected numeric representation of J to be 11, got %d", cardNumericValue("J"))
	}

	if cardNumericValue("Q") != 12 {
		t.Errorf("Expected numeric representation of Q to be 12, got %d", cardNumericValue("Q"))
	}

	if cardNumericValue("K") != 13 {
		t.Errorf("Expected numeric representation of K to be 13, got %d", cardNumericValue("K"))
	}

	if cardNumericValue("6") != 6 {
		t.Errorf("Expected numeric representation of \"6\" to be 6, got %d", cardNumericValue("6"))
	}
}

func TestCardFromString(t *testing.T) {
	c := cardFromString("10\u2666")

	if c.suit != "\u2666" {
		t.Errorf("Expected suit to read as \u2666, got %s", c.suit)
	}

	if c.value != 10 {
		t.Errorf("Expected card value to be 10, got %d", c.value)
	}
}
