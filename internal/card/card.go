package card

import "github.com/gophercises/deck"

type Card struct {
	deck.Card
	Visible bool
}
