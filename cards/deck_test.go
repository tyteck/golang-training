package main

import (
	"fmt"
	"testing"
)

func TestGetRandomValue(t *testing.T) {
	v := getRandomValue(10)
	if !(0 <= v && v <= 10) {
		t.Errorf("Expected random value was between 0 and 10, obtained %v", v)
	}
}

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedNumberOfCards := 16
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
