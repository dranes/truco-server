package main

import "fmt"

//Board the game itself
type Board struct {
	firstTurn     *Player
	currentPlayer *Player
	currentTurn   int
	team1Score    int
	team2Score    int
	team1         []*Player
	team2         []*Player
	playByTurn    [][]*Card
	winnerByTurn  *Player
	player1       *Player
	player2       *Player
	player3       *Player
	player4       *Player
	vira          *Card
}

//Play put a card in game
func (board *Board) Play(cardPosition int) (bool, bool, *Card) {
	var cardInPlay *Card

	endOfTurn := false
	endOfGame := false

	switch cardPosition {
	case 1:
		cardInPlay = board.currentPlayer.hand.card1
		board.currentPlayer.hand.card1 = nil
	case 2:
		cardInPlay = board.currentPlayer.hand.card2
		board.currentPlayer.hand.card2 = nil
	case 3:
		cardInPlay = board.currentPlayer.hand.card3
		board.currentPlayer.hand.card3 = nil
	default:
		panic("Invalid card")
	}

	board.currentPlayer = board.currentPlayer.rightPlayer

	board.playByTurn[board.currentTurn] = append(board.playByTurn[board.currentTurn], cardInPlay)

	if len(board.playByTurn[board.currentTurn]) == 4 { //end of turn
		board.endOfTurn()
		endOfTurn = true
	}

	if board.currentTurn == 3 { //end of game
		endOfGame = true
	}

	return endOfTurn, endOfGame, cardInPlay
}

func (board *Board) endOfTurn() {
	var highestCard = board.playByTurn[board.currentTurn][0]
	for _, card := range board.playByTurn[board.currentTurn] {
		if card.gameValue > highestCard.gameValue {
			highestCard = card
		}
	}

	board.currentPlayer = highestCard.belongsTo.belongsTo
	board.currentTurn++

}
func (board *Board) showGame() {
	fmt.Println("----------------------------------------------")
	fmt.Println(board.player1.showHand())
	fmt.Println(board.player2.showHand())
	fmt.Println(board.player3.showHand())
	fmt.Println(board.player4.showHand())

	//fmt.Println(board.vira.face)
}

//NewBoard constructor
func NewBoard(player1 *Player, player2 *Player, player3 *Player, player4 *Player, vira *Card) *Board {
	board := new(Board)
	board.currentTurn = 0

	//Initialize array
	board.playByTurn = make([][]*Card, 3)
	for i := range board.playByTurn {
		board.playByTurn[i] = make([]*Card, 0)
	}

	board.player1 = player1
	board.player2 = player2
	board.player3 = player3
	board.player4 = player4
	board.vira = vira

	board.currentPlayer = player1
	board.firstTurn = player1

	player1.rightPlayer = player2
	player2.rightPlayer = player3
	player3.rightPlayer = player4
	player4.rightPlayer = player1

	player1.leftPlayer = player4
	player4.leftPlayer = player3
	player3.leftPlayer = player2
	player2.leftPlayer = player1

	return board
}
