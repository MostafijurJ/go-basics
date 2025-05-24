package variables

import (
	"fmt"
	"math/rand"
)

func PrintVariables() {
	cards := newDeck()
	cards.print()
}

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, "~~ ", card)
	}
}

func newDeck() deck {
	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	ranks := []string{"Ace", "Two", "Three", "Four", "Five"}

	cards := deck{}
	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, rank+" of "+suit)
		}
	}

	return cards
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	fmt.Println("Deck shuffled")
}
