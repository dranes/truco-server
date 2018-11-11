package truco

//Truco command
type Truco struct {
}

//Handler truco constructor
func (truco *Truco) Handler() string {
	return "truco"
}

//Run run truco
func (truco *Truco) Run(board *Board, command string, arg int) bool {
	return true
}

//NewTruco truco constructor
func NewTruco() Command {
	var truco = new(Truco)

	return truco
}
