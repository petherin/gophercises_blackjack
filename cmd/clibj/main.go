package main

import (
	"fmt"
	"github.com/petherin/gophercises_blackjack/internal/game"
)

func main() {
	game := game.NewGame(2, 2, 1)
	fmt.Printf("%v", game)

	game.Start()
	// TODO
	//  There are some players
	//  They sit in a particular order
	//  One is the dealer
	//  They hold cardsSteve
	//  Cards are facedown or faceup
	//  A player has a 'turn' which is active or inactive
	//  There is a game rule saying each player is dealt 2 cards
	//  Players makes a play - Hit or Stand
	//  Hit triggers these actions:
	//  	the dealer deals them 1 more card
	//      calculate total of hand
	//		if they bust, game is over and dealer wins - so a player must have a won or lost field
	//		the player selects another play
	//  Stand triggers one result:
	//  	their turn ends and the next player takes their turn
	//	 In version 1 of game, dealer shows their had without taking a turn.
	//	 Determine winner
	//	   Calculate hand total - Ace is 1 or 11, whichever gets closer but not above 21, picture cards worth 10
	//	   Add up points of dealer's cards
	//     Add up points of all other player's hands
	//     Winner is player who gets closest to 21. If dealer and a player tie for highest, dealer wins

}
