package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func createDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func deal(d deck, sizeToDeal int) (deck, deck) {
	return d[:sizeToDeal], d[sizeToDeal:]
}

func (d deck) print() {

	for index, deck := range d {
		fmt.Println(index, deck)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {

	bs, err := ioutil.ReadFile(fileName)
	if err != nil {

	}

}
