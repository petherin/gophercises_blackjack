package player

import (
	"fmt"
	"github.com/petherin/gophercises_blackjack/internal/card"
)

type Player struct {
	Name   string
	Hand   []card.Card
	Turn   bool
	Winner bool
	Bust   bool
	Deck   []card.Card
	Dealer bool
	AI     bool
}

func NewPlayer(name string, deck []card.Card, dealer bool, ai bool) Player {
	return Player{
		Name:   name,
		Deck:   deck,
		Dealer: dealer,
		AI:     ai,
	}
}

func (p *Player) DealACard() (*card.Card, error) {
	// DealACard a card from deck
	if !p.Dealer {
		return nil, fmt.Errorf("Player not a dealer")
	}

	if p.Deck == nil {
		return nil, fmt.Errorf("Deck is nil")
	}

	if len(p.Deck) == 0 {
		return nil, fmt.Errorf("Deck is empty")
	}

	var dealtCard card.Card
	dealtCard, p.Deck = p.Deck[len(p.Deck)-1], p.Deck[:len(p.Deck)-1]

	return &dealtCard, nil
}

func (p *Player) Hit() {
	// Dealer deals 1 more card
	// Calculate value of cards
	// If bust, set Bust property
	// If one player remains who isn't Bust, they Win
}

func (p *Player) AddUp() {

}

func (p *Player) DetermineWinner() {

}
