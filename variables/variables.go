package variables

import "fmt"

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
	ranks := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cards := deck{}
	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, rank+" of "+suit)
		}
	}

	return cards
}
