package variables

import "testing"

func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != 20 {
		t.Errorf("Expected deck length of 20, but got %v", len(d))
	}

	expectedFirstCard := "Ace of Spades"
	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card to be '%s', but got '%s'", expectedFirstCard, d[0])
	}

	expectedLastCard := "Five of Clubs"
	if d[len(d)-1] != expectedLastCard {
		t.Errorf("Expected last card to be '%s', but got '%s'", expectedLastCard, d[len(d)-1])
	}

	d.shuffle()

}
