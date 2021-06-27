package blackjack

import (
	"fmt"
	"gophercises/deck"
)

type AI interface {
	Bet(shuffled bool) int
	Play(hands []deck.Card, dealer deck.Card) Move
	Results(hands [][]deck.Card, dealer []deck.Card)
}

type dealerAI struct{}

func (ai dealerAI) Bet(shuffled bool) int {
	// noop
	return 1
}

func (ai dealerAI) Play(hands []deck.Card, dealer deck.Card) Move {
	dScore := Score(hands...)
	if dScore <= 16 || (dScore == 17 && Soft(hands...)) {
		return MoveHit
	}
	return MoveStand
}

func (ai dealerAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	// noop
}

type humanAI struct{}

func HumanAI() AI {
	return humanAI{}
}

func (ai humanAI) Bet(shuffled bool) int {
	if shuffled {
		fmt.Println("The deck was just shuffled")
	}
	fmt.Println("What would you like to bet?")
	var bet int
	fmt.Scanf("%d\n", &bet)
	return bet
}

func (ai humanAI) Play(hands []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player:", hands)
		fmt.Println("Dealer:", dealer)
		fmt.Println("What will you do? (h)it, (s)tand, (d)ouble, s(p)lit")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit
		case "s":
			return MoveStand
		case "d":
			return MoveDouble
		case "p":
			return MoveSplit
		default:
			fmt.Println("Invalid option:", input)
		}
	}
}

// TODO:
func (ai humanAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:")
	for _, h := range hands {
		fmt.Println(" ", h)
	}
	fmt.Println("Dealer:", dealer)
	fmt.Println()
}
