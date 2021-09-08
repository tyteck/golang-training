package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	cards := newDeck()
	enmain := "./enmain.deck"

	enMain, laPioche := deal(cards, 4)

	title("En main")
	enMain.print()
	title("la pioche")
	laPioche.print()

	fmt.Println("As a string : " + laPioche.toString())

	fmt.Println("==== saving file ====")
	enMain.toFile(enmain)
	laPioche.toFile("./lapioche.deck")

	fmt.Println("==== reading file " + enmain + "====")
	d := fromFile(enmain)
	d.print()

	fmt.Println("==== randomizing ====")
	d.shuffle()
	d.print()

}

func title(s string) {
	fmt.Println("======== " + s + " ========")
}

func getNewSeed() *rand.Rand {
	source := rand.NewSource(time.Now().UnixNano())
	return rand.New(source)
}

func getRandomValue(max int) int {
	r := getNewSeed()
	return r.Intn(max)
}
