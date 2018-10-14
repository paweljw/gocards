package main

import (
	"strconv"
)

type card struct {
	suit  string
	value int
}

func (c card) toString() string {
	return c.humanValue() + c.suit
}

func (c card) humanValue() string {
	switch c.value {
	case 1:
		return "A"
	case 11:
		return "J"
	case 12:
		return "Q"
	case 13:
		return "K"
	default:
		return strconv.Itoa(c.value)
	}
}

func cardNumericValue(s string) int {
	switch s {
	case "A":
		return 1
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	default:
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return num
	}
}

func cardFromString(s string) card {
	val, suit := s[:len(s)-3], s[len(s)-3:]
	return card{value: cardNumericValue(val), suit: suit}
}
