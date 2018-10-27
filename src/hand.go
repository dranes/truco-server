package main

import (
	"fmt"
	"sort"
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
	card1             *Card
	card2             *Card
	card3             *Card
	cardsOrderByValue []Card
	enviste           bool
	envisteValue      int
	flor              bool
	florValue         int
	suitMapCards      map[int][]*Card
	perico            *Card
	perica            *Card
}

type byEnviste []Card

func (s byEnviste) Len() int {
	return len(s)
}

func (s byEnviste) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byEnviste) Less(i, j int) bool {
	return s[i].envisteValue > s[j].envisteValue
}

func (hand *Hand) transform(card *Card, vira Card) {

	if vira.value == 10 && vira.suit == card.suit && card.value == 11 {
		card.gameValue = PericaValue
		card.face = fmt.Sprintf("%-13s", Perica)
		hand.perica = card
	}

	if vira.value == 10 && vira.suit == card.suit && card.value == 12 {
		card.gameValue = PericoValue
		card.face = fmt.Sprintf("%-13s", Perico)
		hand.perico = card
	}

	if vira.value == 11 && vira.suit == card.suit && card.value == 10 {
		card.gameValue = PericaValue
		card.face = fmt.Sprintf("%-13s", Perica)
		hand.perica = card
	}

	if vira.value == 11 && vira.suit == card.suit && card.value == 12 {
		card.gameValue = PericoValue
		card.face = fmt.Sprintf("%-13s", Perico)
		hand.perico = card
	}

	if vira.suit == card.suit && card.value == 10 {
		card.gameValue = PericaValue
		card.face = fmt.Sprintf("%-13s", Perica)
		hand.perica = card
	}

	if vira.suit == card.suit && card.value == 11 {
		card.gameValue = PericoValue
		card.face = fmt.Sprintf("%-13s", Perico)
		hand.perico = card
	}
}

//NewHand Constructor
func NewHand(cards []Card, vira Card) *Hand {

	hand := new(Hand)
	hand.card1 = &cards[0]
	hand.card2 = &cards[1]
	hand.card3 = &cards[2]

	sort.Sort(byEnviste(cards))
	hand.cardsOrderByValue = cards

	hand.transform(hand.card1, vira)
	hand.transform(hand.card2, vira)
	hand.transform(hand.card3, vira)

	hand.suitMapCards = make(map[int][]*Card, 0)
	if hand.card1 != hand.perico && hand.card1 != hand.perica {
		hand.suitMapCards[hand.card1.suit] = append(hand.suitMapCards[hand.card1.suit], hand.card1)
	}

	if hand.card2 != hand.perico && hand.card2 != hand.perica {
		hand.suitMapCards[hand.card2.suit] = append(hand.suitMapCards[hand.card2.suit], hand.card2)
	}

	if hand.card3 != hand.perico && hand.card3 != hand.perica {
		hand.suitMapCards[hand.card3.suit] = append(hand.suitMapCards[hand.card3.suit], hand.card3)
	}

	//pendiente de calcular flor reservada

	if len(hand.suitMapCards) == 3 {
		//no hay enviste a menos que haya perico o perica

	}

	if len(hand.suitMapCards) == 2 {
		for _, cards := range hand.suitMapCards {
			if len(cards) == 2 {
				for _, card := range cards {
					hand.envisteValue += card.envisteValue
				}
				hand.envisteValue += 20
			}
		}

		if hand.perica != nil {
			//vale 29
			hand.envisteValue = 29 + hand.cardsOrderByValue[0].envisteValue
		}

		if hand.perico != nil {
			//vale 30
			hand.envisteValue = 30 + hand.cardsOrderByValue[0].envisteValue
		}
		//hay enviste o flor a menos que haya perico o perica
	}

	if len(hand.suitMapCards) == 1 {
		//hay flor normal y corriente
	}

	//calculate enviste value
	return hand

}
