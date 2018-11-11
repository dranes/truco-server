package main

import (
	"fmt"

	"github.com/dranes/truco"
)

func main() {
	deck := truco.NewDeck()

	deck.Shuffle()

	hand1, hand2, hand3, hand4, vira := deck.Deal()

	player1Hand := truco.NewHand(hand1, vira)
	player2Hand := truco.NewHand(hand2, vira)
	player3Hand := truco.NewHand(hand3, vira)
	player4Hand := truco.NewHand(hand4, vira)

	player1 := truco.NewPlayer("Player 1", player1Hand)
	player2 := truco.NewPlayer("Player 2", player2Hand)
	player3 := truco.NewPlayer("Player 3", player3Hand)
	player4 := truco.NewPlayer("Player 4", player4Hand)

	//construct board

	board := truco.NewBoard(player1, player2, player3, player4, &vira)
	//var play int
	var err error
	var card *truco.Card
	var command string
	var arg int

	//var cardPlayed *Card
	endOfTurn := false
	endOfGame := false

	board.ShowGame()

	for endOfGame == false {
		fmt.Printf("Turn: %s> ", board.CurrentPlayer.Name)
		fmt.Scanf("%s %d", &command, &arg)

		endOfTurn, endOfGame, card, err = board.Play(command, arg)

		if err == nil {
			fmt.Println(card.Face)
		}

		if err != nil {
			fmt.Println(err)
		}

		//board.showGame()
		if endOfTurn {
			board.ShowGame()
		}

		if endOfGame {
			fmt.Printf("----------- END OF GAME -----------\n")
			fmt.Printf("Team1 %d pts Team2 %d pts\n", board.Team1WinByTurn, board.Team2WinByTurn)
		}
	}
}
