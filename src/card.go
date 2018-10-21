package main

import "fmt"

//Type of cards
const (
	Swords int = 1
	Sticks int = 2
	Coins  int = 3
	Cups   int = 4
)

//Suits Array with valid suits
var Suits = []int{Swords, Sticks, Coins, Cups}

//Card represent a single card in the game
type Card struct {
	suit      int
	value     int
	gameValue int
	face      string
}

func face(suit int, value int) string {
	if suit == Coins && value == 7 {
		return fmt.Sprintf("%-13s", "7 de oro")
	}

	if suit == Swords && value == 7 {
		return fmt.Sprintf("%-13s", "7 de espadas")
	}

	if suit == Sticks && value == 1 {
		return fmt.Sprintf("%-13s", "Bastillo")
	}

	if suit == Swords && value == 1 {
		return fmt.Sprintf("%-13s", "Espadilla")
	}

	return fmt.Sprintf("%-2d %-9s ", value, suitType(suit))
}

func suitType(suit int) string {

	switch suit {

	case Swords:
		return "Swords"
	case Sticks:
		return "Sticks"
	case Coins:
		return "Coins"
	case Cups:
		return "Cups"

	default:
		return ""
	}
}

//NewCard constructor for Card
func NewCard(suit, value int) Card {
	gameValue := 0

	switch value {
	case 4:
		gameValue = 1
	case 5:
		gameValue = 2
	case 6:
		gameValue = 3
	case 7:
		gameValue = 4
	case 8:
		gameValue = 5
		value = 10
	case 9:
		gameValue = 6
		value = 11
	case 10:
		gameValue = 7
		value = 12
	case 1:
		gameValue = 8
	case 2:
		gameValue = 9
	case 3:
		gameValue = 10
	}

	if suit == Coins && value == 7 { //7 de oro
		gameValue = 11
	}

	if suit == Swords && value == 7 { //7 de espadas
		gameValue = 12
	}

	if suit == Sticks && value == 1 { //bastillo
		gameValue = 13
	}

	if suit == Swords && value == 1 { //espadilla
		gameValue = 14
	}

	card := Card{suit, value, gameValue, face(suit, value)}

	return card
}
