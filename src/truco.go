package main

import "fmt"

func main() {
	deck := NewDeck()

	deck.Shuffle()

	player1, player2, player3, player4, vira := deck.Deal()

	player1Hand := NewHand(player1, vira)
	player2Hand := NewHand(player2, vira)
	player3Hand := NewHand(player3, vira)
	player4Hand := NewHand(player4, vira)

	fmt.Println(player1Hand.card1.face, player1Hand.card2.face, player1Hand.card3.face)
	fmt.Println(player2Hand.card1.face, player2Hand.card2.face, player2Hand.card3.face)
	fmt.Println(player3Hand.card1.face, player3Hand.card2.face, player3Hand.card3.face)
	fmt.Println(player4Hand.card1.face, player4Hand.card2.face, player4Hand.card3.face)

	fmt.Println(vira.face)
}
