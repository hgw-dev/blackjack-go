package game

import (
	"fmt"
	"blackjack/deck"
)

func init(){
	fmt.Println("Welcome to the Blackjack table!")
	fmt.Println("Shuffling...\n\n")
}

func Start(){
	deck.Draw()
}