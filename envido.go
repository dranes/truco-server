package truco

import "fmt"

//Enviste when two cards has the same suit
type Enviste struct {
	whoAsks    *Player
	whoAnswers *Player
	points     int
}

//Handler how this command is named
func (enviste *Enviste) Handler() string {
	return "envido"
}

//Run enviste
//Siempre se envida al de la izquierda
func (enviste *Enviste) Run(board *Board, command string, arg int) bool {

	switch command {
	case "envido":
		if enviste.whoAsks == nil {
			enviste.whoAsks = board.CurrentPlayer
			enviste.whoAnswers = board.CurrentPlayer.leftPlayer
		} else {
			enviste.whoAsks = board.CurrentPlayer
			enviste.whoAnswers = board.CurrentPlayer.rightPlayer
		}

		board.CurrentPlayer = enviste.whoAnswers
		board.currentPlayerCommand = enviste.whoAnswers
		enviste.points += arg //cantidad de enviste apostado

		//empieza el envido o envida mas piedras

	case "quiero":
		board.CurrentPlayer = board.CurrentPlayerHolder
		fmt.Printf("\nEnvido Total Bet %d\n", enviste.points)
		fmt.Printf("%d", enviste.points)
		return true
	case "no":
		board.CurrentPlayer = board.CurrentPlayerHolder
		fmt.Printf("\nEnvido Total Bet %d\n", enviste.points)

		return true
		//se anula el envido se reparten los puntos
	default:
		//se apuestan mas puntos se espera respuesta

	}

	return false
}

func (enviste *Enviste) resolve() {

}

//NewEnviste enviste constructor
func NewEnviste() Command {
	var enviste = new(Enviste)

	enviste.points = 1 //always the fearless player gets a reward

	return enviste
}
