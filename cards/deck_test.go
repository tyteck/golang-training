package main

import (
	"fmt"
	"testing"
)

const TOTAL_NUMBER_OF_CARDS = 16

func TestGetRandomValue(t *testing.T) {
	v := getRandomValue(10)
	if !(0 <= v && v <= 10) {
		t.Errorf("Expected random value was between 0 and 10, obtained %v", v)
	}
}

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedNumberOfCards := TOTAL_NUMBER_OF_CARDS
	expectedFirstCard := "As de Carreau"
	expectedLastCard := "Trois de Trèfle"

	if len(d) != expectedNumberOfCards {
		t.Errorf("Expected number of cards was "+fmt.Sprint(expectedNumberOfCards)+", obtained %v", len(d))
	}

	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card was "+fmt.Sprint(expectedFirstCard)+", obtained %v", d[0])
	}

	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected last card was "+fmt.Sprint(expectedLastCard)+", obtained %v", d[len(d)-1])
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	expectedNumberOfCardsInDeck1 := 6
	expectedNumberOfCardsInDeck2 := TOTAL_NUMBER_OF_CARDS - expectedNumberOfCardsInDeck1

	deck1, deck2 := deal(d, expectedNumberOfCardsInDeck1)
	if len(deck1) != expectedNumberOfCardsInDeck1 {
		t.Errorf("Expected number of cards in deck 1 was "+fmt.Sprint(expectedNumberOfCardsInDeck1)+", obtained %v", len(deck1))
	}

	if len(deck2) != expectedNumberOfCardsInDeck2 {
		t.Errorf("Expected number of cards in deck 2 was "+fmt.Sprint(expectedNumberOfCardsInDeck2)+", obtained %v", len(deck2))
	}
}
