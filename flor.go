package truco

//Flor command
type Flor struct {
}

//Handler flor command
func (flor *Flor) Handler() string {
	return "flor"
}

//Run Flor
func (flor *Flor) Run(board *Board, command string, arg int) bool {
	return true
}

//NewFlor Flor constructor
func NewFlor() Command {
	var flor = new(Flor)

	return flor
}
