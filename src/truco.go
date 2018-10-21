package main

import "fmt"

func main() {
	deck := NewDeck()

	deck.Shuffle()

	hand1, hand2, hand3, hand4, vira := deck.Deal()

	player1Hand := NewHand(hand1, vira)
	player2Hand := NewHand(hand2, vira)
	player3Hand := NewHand(hand3, vira)
	player4Hand := NewHand(hand4, vira)

	player1 := NewPlayer("Player 1", player1Hand)
	player2 := NewPlayer("Player 2", player2Hand)
	player3 := NewPlayer("Player 3", player3Hand)
	player4 := NewPlayer("Player 4", player4Hand)

	fmt.Println(player1.showHand())
	fmt.Println(player2.showHand())
	fmt.Println(player3.showHand())
	fmt.Println(player4.showHand())

	fmt.Println(vira.face)
}
