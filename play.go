package truco

//Play command
type Play struct {
}

//Handler return command
func (play *Play) Handler() string {
	return "play"
}

//Run execute a command
func (play *Play) Run(board *Board, command string, arg int) bool {
	var cardInPlay *Card
	//var err error

	cardPosition := arg

	switch cardPosition {
	case 1:
		cardInPlay = board.CurrentPlayer.hand.card1
		board.CurrentPlayer.hand.card1 = nil
	case 2:
		cardInPlay = board.CurrentPlayer.hand.card2
		board.CurrentPlayer.hand.card2 = nil
	case 3:
		cardInPlay = board.CurrentPlayer.hand.card3
		board.CurrentPlayer.hand.card3 = nil
	default:
		//err = errors.New("Invalid card")
	}

	if cardInPlay == nil {
		//return false, false, nil, errors.New("Invalid card")
	}

	board.CurrentPlayer = board.CurrentPlayer.rightPlayer

	board.playByTurn[board.currentTurn] = append(board.playByTurn[board.currentTurn], cardInPlay)

	board.lastCardPlayed = cardInPlay

	return true

}

//NewPlay Constructor
func NewPlay() Command {
	var play = new(Play)

	return play
}
