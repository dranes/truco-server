package main

import "fmt"

//Player struct
type Player struct {
	name        string
	hand        *Hand
	teammate    *Player
	rightPlayer *Player
	leftPlayer  *Player
}

//NewPlayer player Constructor
func NewPlayer(name string, hand *Hand) *Player {
	player := new(Player)
	player.name = name
	player.hand = hand

	return player
}

func (player Player) showHand() string {
	return fmt.Sprintf("%s %s %s", player.hand.card1.face, player.hand.card2.face, player.hand.card3.face)
}
