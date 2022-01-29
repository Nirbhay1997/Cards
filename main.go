package main

import "fmt"

var card string

func main() {

	card = newCard()
	//card := "Ace of Spades"
	fmt.Println(card)

}

func newCard() string {
	return "Five of diamonds"
}
