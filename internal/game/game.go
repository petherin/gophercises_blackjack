package game

import (
	"bufio"
	"fmt"
	"github.com/gophercises/deck"
	"github.com/petherin/gophercises_blackjack/internal/card"
	"github.com/petherin/gophercises_blackjack/internal/player"
	"log"
	"os"
	"regexp"
	"strconv"
)

var names = []string{"Fred", "Dan"}

type Game struct {
	CardsPerPlayer int
	Players        []*player.Player
	TotalTurns     int
}

func NewGame(numOfPlayers int, cardsPerPlayer int, totalTurns int) Game {
	// get the deck of cards and convert them to our cards
	deck := deck.New(deck.Deck(1), deck.Shuffle)
	var ourDeck []card.Card
	for _, c := range deck {
		card := card.Card{
			c, false,
		}
		ourDeck = append(ourDeck, card)
	}

	var players []*player.Player

	for i := 0; i < numOfPlayers; i++ {
		var newPlayer player.Player
		var dealer bool
		var deck []card.Card
		var ai bool

		if i == 0 {
			dealer = true
			deck = ourDeck
			ai = true
		}

		newPlayer = player.NewPlayer(names[i], deck, dealer, ai)
		players = append(players, &newPlayer)

	}

	return Game{CardsPerPlayer: cardsPerPlayer, Players: players, TotalTurns: totalTurns}
}

func (g *Game) Setup() error {
	var dealer *player.Player

	for _, player := range g.Players {
		if player.Dealer {
			dealer = player
			break
		}
	}
	fmt.Printf("%v", dealer)

	for i := 0; i < g.CardsPerPlayer; i++ {
		for _, player := range g.Players {
			card, err := dealer.DealACard()
			if err != nil {
				return err
			}

			if player.Dealer && i == 0 {
				card.Visible = true
			}

			player.Hand = append(player.Hand, *card)
		}
	}

	if len(g.Players) > 1 {
		g.Players[1].Turn = true
	}

	return nil
}

func (g *Game) Start() {
	if err := g.Setup(); err != nil {
		log.Printf("%v", err)
		return
	}

	// Keep playing until someone wins or the number of turns is reached
	turns := 0
	var dealer *player.Player
	for _, dealer = range g.Players {
		if dealer.Dealer {
			break
		}
	}

	for !g.Won() || turns < g.TotalTurns {
		// Player whose turn it is must decide what to do
		// Player could be an AI or a person.
		var player *player.Player
		for _, player = range g.Players {
			if player.Turn {
				break
			}
		}

		scanner := bufio.NewScanner(os.Stdin)
		pattern := regexp.MustCompile("[1-2]")

		if !player.AI {
			// Tell player what their cards are
			// Tell player what dealers visible card is
			// Ask player Hit or Stand
			// If Stand, turn is over
			// If Hit, deal a card and check if they've bust. If not, ask them to Hit or Stand again.
			fmt.Printf("\nIt's %s's turn. Your cards are:\n%s\nThe dealer's visible card is %s\n", player.Name, player.Hand, "get dealer's hand")
			fmt.Printf("Do you want to (1) hit or (2) stand? Choose and press Enter...")

			scanner.Scan()
			textInput := scanner.Text()
			if !pattern.MatchString(textInput) {
				fmt.Printf("Invalid choice\n")
				continue
			}

			action, err := strconv.Atoi(textInput)
			if err != nil {
				fmt.Printf("%v\n", err)
				continue
			}

			fmt.Printf("You chose %d", action)
		}
		turns++
	}
}

func (g *Game) Won() bool {
	playersLeft := len(g.Players)
	for _, player := range g.Players {
		if player.Bust {
			playersLeft--
		}
	}

	return playersLeft == 1
}
