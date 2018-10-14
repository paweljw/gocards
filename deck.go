package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

type deck []card

func newDeck() deck {
	cardSuits := []string{"\u2660", "\u2665", "\u2666", "\u2663"}

	cards := deck{}

	for _, suit := range cardSuits {
		for i := 0; i < 13; i++ {
			cards = append(cards, card{suit: suit, value: i + 1})
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card.toString())
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toSliceOfString() []string {
	var slice []string

	for _, card := range d {
		slice = append(slice, card.toString())
	}

	return slice
}

func (d deck) toString() string {
	return strings.Join(d.toSliceOfString(), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

func newDeckFromFile(filename string) deck {
	strDeck, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(127)
	}

	var d deck

	for _, s := range strings.Split(string(strDeck), ",") {
		d = append(d, cardFromString(s))
	}

	return d
}

func (d deck) shuffle() {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
