package truco

import (
	"math/rand"
	"time"
)

//Deck represents a collection of cards
type Deck struct {
	cards []Card
}

//NewDeck constructor for Deck
func NewDeck() *Deck {
	deck := new(Deck)
	deck.cards = make([]Card, 0, 40)

	for _, suit := range Suits {
		for counter := 1; counter <= 10; counter++ {
			deck.cards = append(deck.cards, NewCard(suit, counter))
		}
	}

	return deck
}

//Shuffle cards in deck
func (deck Deck) Shuffle() {

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.cards), func(i, j int) { deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i] })

}

//Deal cards
func (deck Deck) Deal() ([]Card, []Card, []Card, []Card, Card) {
	//Deal one card per player 3 times
	player1 := make([]Card, 0, 3)
	player2 := make([]Card, 0, 3)
	player3 := make([]Card, 0, 3)
	player4 := make([]Card, 0, 3)

	for counter := 0; counter <= 2; counter++ {
		player1 = append(player1, deck.cards[len(deck.cards)-1])
		deck.cards = deck.cards[0 : len(deck.cards)-1]

		player2 = append(player2, deck.cards[len(deck.cards)-1])
		deck.cards = deck.cards[0 : len(deck.cards)-1]

		player3 = append(player3, deck.cards[len(deck.cards)-1])
		deck.cards = deck.cards[0 : len(deck.cards)-1]

		player4 = append(player4, deck.cards[len(deck.cards)-1])
		deck.cards = deck.cards[0 : len(deck.cards)-1]
	}

	vira := deck.cards[len(deck.cards)-1]
	deck.cards = deck.cards[0 : len(deck.cards)-1]

	return player1, player2, player3, player4, vira

}
