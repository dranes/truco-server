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

	//construct board

	board := NewBoard(player1, player2, player3, player4, &vira)
	var play int
	var err error

	//var cardPlayed *Card
	endOfTurn := false
	endOfGame := false

	board.showGame()

	for endOfGame == false {
		fmt.Printf("Turn: %s> ", board.currentPlayer.name)
		fmt.Scan(&play)
		endOfTurn, endOfGame, _, err = board.Play(play)

		if err != nil {
			fmt.Println(err)
		}

		//board.showGame()
		if endOfTurn {
			board.showGame()
		}
	}
}
