package main

import (
	"fmt"
	"sort"
)

//Most powerful cards in game
const (
	PericaValue        = 15
	Perica             = "Perica"
	PericoValue        = 16
	Perico             = "Perico"
	PericaEnvisteValue = 29
	PericoEnvisteValue = 30
	EnvisteValue       = 20
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
	belongsTo         *Player
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
		return
	}

	if vira.value == 10 && vira.suit == card.suit && card.value == 12 {
		card.gameValue = PericoValue
		card.face = fmt.Sprintf("%-13s", Perico)
		hand.perico = card
		return
	}

	if vira.value == 11 && vira.suit == card.suit && card.value == 10 {
		card.gameValue = PericaValue
		card.face = fmt.Sprintf("%-13s", Perica)
		hand.perica = card
		return
	}

	if vira.value == 11 && vira.suit == card.suit && card.value == 12 {
		card.gameValue = PericoValue
		card.face = fmt.Sprintf("%-13s", Perico)
		hand.perico = card
		return
	}

	if vira.suit == card.suit && card.value == 10 {
		card.gameValue = PericaValue
		card.face = fmt.Sprintf("%-13s", Perica)
		hand.perica = card
		return
	}

	if vira.suit == card.suit && card.value == 11 {
		card.gameValue = PericoValue
		card.face = fmt.Sprintf("%-13s", Perico)
		hand.perico = card
		return
	}
}

//NewHand Constructor
func NewHand(cards []Card, vira Card) *Hand {

	hand := new(Hand)
	cards[0].belongsTo = hand
	cards[1].belongsTo = hand
	cards[2].belongsTo = hand

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

	if len(hand.suitMapCards) == 2 {
		for _, cards := range hand.suitMapCards {
			if len(cards) == 2 {
				for _, card := range cards {
					hand.envisteValue += card.envisteValue
				}

				if hand.perica != nil { //flor con perica
					hand.flor = true
					hand.florValue = PericaEnvisteValue + cards[0].envisteValue + cards[1].envisteValue
				}

				if hand.perico != nil {
					hand.flor = true
					hand.florValue = PericoEnvisteValue + cards[0].envisteValue + cards[1].envisteValue
				}

				hand.envisteValue += EnvisteValue
			}
		}

		if hand.perica != nil {
			//vale 29
			hand.envisteValue = PericaEnvisteValue + hand.cardsOrderByValue[0].envisteValue
		}

		if hand.perico != nil {
			//vale 30
			hand.envisteValue = PericoEnvisteValue + hand.cardsOrderByValue[0].envisteValue
		}
	}

	if len(hand.suitMapCards) == 1 {
		hand.flor = true
		hand.florValue = cards[0].envisteValue + cards[1].envisteValue + cards[2].envisteValue + EnvisteValue
		//hay flor normal y corriente
	}

	return hand

}
