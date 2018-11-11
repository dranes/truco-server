package truco

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
func (enviste *Enviste) Run(board *Board, command string, arg int) bool {

	switch command {
	case "envido":
		//empieza el envido
	case "no":
		//se anula el envido se reparten los puntos
	default:
		//se apuestan mas puntos se espera respuesta

	}

	return true
}

//NewEnviste enviste constructor
func NewEnviste() Command {
	var enviste Command = new(Enviste)

	return enviste
}
