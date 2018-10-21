package main

import (
	"fmt"
)

//Most powerful cards in game
const (
	PericaValue = 15
	Perica      = "Perica"
	PericoValue = 16
	Perico      = "Perico"
)

//Hand 3 cards owned by a player
type Hand struct {
	card1 Card
	card2 Card
	card3 Card
}

func transform(card *Card, vira Card) {

	if vira.value == 10 && vira.suit == card.suit && card.value == 11 {
		card.gameValue = PericaValue
		card.face = fmt.Sprintf("%-13s", Perica)
	}

	if vira.value == 10 && vira.suit == card.suit && card.value == 12 {
		card.gameValue = PericoValue
		card.face = fmt.Sprintf("%-13s", Perico)
	}

	if vira.value == 11 && vira.suit == card.suit && card.value == 10 {
		card.gameValue = PericaValue
		card.face = fmt.Sprintf("%-13s", Perica)
	}

	if vira.value == 11 && vira.suit == card.suit && card.value == 12 {
		card.gameValue = PericoValue
		card.face = fmt.Sprintf("%-13s", Perico)
	}

	if vira.suit == card.suit && card.value == 10 {
		card.gameValue = PericaValue
		card.face = fmt.Sprintf("%-13s", Perica)
	}

	if vira.suit == card.suit && card.value == 11 {
		card.gameValue = PericoValue
		card.face = fmt.Sprintf("%-13s", Perico)
	}
}

//NewHand Constructor
func NewHand(cards []Card, vira Card) *Hand {
	hand := new(Hand)
	hand.card1 = cards[0]
	hand.card2 = cards[1]
	hand.card3 = cards[2]

	transform(&hand.card1, vira)
	transform(&hand.card2, vira)
	transform(&hand.card3, vira)

	return hand
}
