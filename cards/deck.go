package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

const separator = ", "

func newDeck() deck {
	cards := deck{}

	cardsColors := []string{"Carreau", "Coeur", "Pique", "Trèfle"}
	cardsValues := []string{"As", "Un", "Deux", "Trois"}
	for _, color := range cardsColors {
		for _, value := range cardsValues {
			cards = append(cards, value+" de "+color)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, size int) (deck, deck) {
	return d[:size], d[size:]
}

func (d deck) toString() string {
	return strings.Join(d, separator)
}

func (d deck) toFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

func fromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error : ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), separator)
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := getRandomValue(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
