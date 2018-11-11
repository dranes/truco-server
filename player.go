package truco

import "fmt"

//Player struct
type Player struct {
	Name        string
	hand        *Hand
	teammate    *Player
	rightPlayer *Player
	leftPlayer  *Player
	team        []*Player
}

//NewPlayer player Constructor
func NewPlayer(name string, hand *Hand) *Player {
	player := new(Player)
	hand.belongsTo = player
	player.Name = name
	player.hand = hand

	return player
}

func (player Player) showHand() string {
	card1 := ""
	card2 := ""
	card3 := ""

	if player.hand.card1 != nil {
		card1 = player.hand.card1.Face
	}
	if player.hand.card2 != nil {
		card2 = player.hand.card2.Face
	}
	if player.hand.card3 != nil {
		card3 = player.hand.card3.Face
	}

	return fmt.Sprintf("%-13s %-13s %-13s %-2d %-2d", card1, card2, card3, player.hand.envisteValue, player.hand.florValue)
}

//PlayCard put a card in game
func (player Player) PlayCard() {

}