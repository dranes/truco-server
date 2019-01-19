package truco

import (
	"fmt"
)

//Command interface for commadns
type Command interface {
	Handler() string
	Run(*Board, string, int) bool
}

//Board the game itself
type Board struct {
	firstTurn            *Player
	CurrentPlayer        *Player
	CurrentPlayerHolder  *Player
	currentPlayerCommand *Player
	currentTurn          int
	team1Score           int
	team2Score           int
	Team1WinByTurn       int
	Team2WinByTurn       int
	team1                []*Player
	team2                []*Player
	playByTurn           [][]*Card
	winnerByTurn         *Player
	player1              *Player
	player2              *Player
	player3              *Player
	player4              *Player
	vira                 *Card
	activeCommand        string
	voice                *Player
	envistePointsBet     int
	enviste              *Enviste
	commands             map[string]Command
	commandLocked        Command
	lastCardPlayed       *Card
	resolved             bool
}

//Play put a card in game
func (board *Board) Play(command string, arg int) (bool, bool, *Card, bool, error) { //cardPosition should be command

	var err error
	var handler Command
	endOfTurn := false
	endOfGame := false

	if board.resolved == false {
		handler = board.commands[board.activeCommand]
	} else {
		handler = board.commands[command]
		board.activeCommand = command
		board.CurrentPlayerHolder = board.CurrentPlayer
	}

	/*
		if key == false {
			err = errors.New("invalid command")
		}
	*/

	board.resolved = handler.Run(board, command, arg)

	if len(board.playByTurn[board.currentTurn]) == 4 { //end of turn
		board.endOfTurn()
		endOfTurn = true
	}

	if board.currentTurn == 3 { //end of game
		endOfGame = true
	}

	return endOfTurn, endOfGame, board.lastCardPlayed, board.resolved, err
}

func (board *Board) endOfTurn() {
	var highestCard = board.playByTurn[board.currentTurn][0]
	for _, card := range board.playByTurn[board.currentTurn] {
		if card.gameValue > highestCard.gameValue {
			highestCard = card
		}
	}

	board.CurrentPlayer = highestCard.belongsTo.belongsTo

	if board.CurrentPlayer == board.team1[0] || board.CurrentPlayer == board.team1[1] {
		board.Team1WinByTurn++
	} else {
		board.Team2WinByTurn++
	}
	board.currentTurn++

}

/*
func (board *Board) resolveCommand(command string, arg int) {
	lock := board.commandLocked

	if lock == nil {
		key, lock := board.commands[command]

		if key == nil {
			return errors.New("invalid command")
		}
	}

	lock.Run(board, command, arg)

	board.commandLocked = lock
}
*/

//ShowGame show status of cards
func (board *Board) ShowGame() {
	fmt.Println("----------------------------------------------")
	var pointer *Player

	for pointer != board.CurrentPlayer {
		if pointer == nil {
			pointer = board.CurrentPlayer
		}
		fmt.Printf("%s --> %s \n", pointer.Name, pointer.showHand())
		pointer = pointer.rightPlayer
	}

	//fmt.Println(board.vira.face)
}

//NewBoard constructor
func NewBoard(player1 *Player, player2 *Player, player3 *Player, player4 *Player, vira *Card) *Board {
	board := new(Board)
	board.currentTurn = 0
	board.team1 = make([]*Player, 2)
	board.team2 = make([]*Player, 2)
	board.team1[0] = player1
	board.team1[1] = player3
	board.team2[0] = player2
	board.team2[1] = player4
	player1.team = board.team1
	player3.team = board.team1
	player2.team = board.team2
	player4.team = board.team2

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

	board.CurrentPlayer = player1
	board.firstTurn = player1
	board.CurrentPlayerHolder = player1

	player1.rightPlayer = player2
	player2.rightPlayer = player3
	player3.rightPlayer = player4
	player4.rightPlayer = player1

	player1.leftPlayer = player4
	player4.leftPlayer = player3
	player3.leftPlayer = player2
	player2.leftPlayer = player1

	//register commands

	enviste := NewEnviste()
	truco := NewTruco()
	flor := NewFlor()
	play := NewPlay()

	board.commands = make(map[string]Command)

	board.commands[enviste.Handler()] = enviste
	board.commands[truco.Handler()] = enviste
	board.commands[flor.Handler()] = enviste
	board.commands[play.Handler()] = play

	board.resolved = true

	return board
}
