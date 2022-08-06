package hand

import (
	"fmt"
)
import "blackjack/src/deck"

type Hand struct {
    Cards   []deck.Card
	Value   int
    Dealer  bool
	Hidden  bool
}

var cardBackUnicode string

func init() {
	cardBackUnicode = "\U0001F0A0"
}

func dealHand(dealer bool) Hand{
	h := Hand{
		Cards: deck.DrawNFromDeck(2),
		Dealer: dealer,
		Hidden: dealer,
	}
	h.Value = h.getValue()
	return h
}

func DealPlayer() Hand{
	return dealHand(false)
}

func DealDealer() Hand{
	return dealHand(true)
}
func (h Hand) printDealer(hidden bool) {
	owner := "Dealer"
	fmt.Printf("%s ", owner)

	startIndex := 0
	if h.Hidden {
		startIndex = 1
		fmt.Printf("%s ", cardBackUnicode)
	}
	
    for _, item := range h.Cards[startIndex:] {
		fmt.Printf("%s ", item.GetCard())
	}

	fmt.Printf(" showing %d\n", h.Value)
	fmt.Printf("\n")
}
func (h Hand) printPlayer() {
	fmt.Printf("Player ")
    for _, item := range h.Cards {
		fmt.Printf(
			"%s ", 
			item.GetCard(),
		)
	}
	fmt.Printf(" showing %d\n", h.Value)
}
func (h Hand) Print() {
	if h.Dealer {
		h.printDealer(true)
	} else {
		h.printPlayer()
	}
}
func (h Hand) getValue() int {
	value := 0
	hasAce := false

	startIndex := 0
	if h.Hidden {
		startIndex = 1
	}
    for _, item := range h.Cards[startIndex:] {
		rank := item.Rank
		if rank > 10 {
			rank = 10
		}
		if rank == 1 {
			hasAce = true
			rank = 11
		}		
		value += rank
    }

	// if we go bust, we want the ace to = 1
	if value > 21 && hasAce {
		value -= 10
	}
	if value == 21 && h.Hidden == false {
		fmt.Println("BLACKJACK")
	}
	
    return value
}
