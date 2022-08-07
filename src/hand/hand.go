package hand

import (
	"fmt"
	"errors"
)
import "blackjack/src/deck"

type Hand struct {
    Cards     deck.Cards
	Value     int
    IsDealer  bool
	IsHidden  bool
}

var cardBackUnicode string

func init() {
	cardBackUnicode = "\U0001F0A0"
}

func dealHand(dealer bool) Hand{
	var cards deck.Cards = []deck.Card{}

	h := Hand{
		Cards: cards,
		IsDealer: dealer,
		IsHidden: dealer,
	}

	h.Cards = h.Cards.DrawNCards(2)
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
	if h.IsHidden {
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
	if h.IsDealer {
		h.printDealer(true)
	} else {
		h.printPlayer()
	}
}
func (h Hand) getValue() int {
	value := 0
	aceCount := 0

	startIndex := 0
	if h.IsHidden {
		startIndex = 1
	}

    for _, item := range h.Cards[startIndex:] {
		rank := item.Rank
		if rank > 10 {
			rank = 10
		}
		if rank == 1 {
			aceCount += 1
			rank = 11
		}		
		value += rank
    }

	// if we go bust, we want the ace to = 1
	for value > 21 && aceCount > 0 {
		value = value - 10 
		aceCount -= 1
	}

	// it's only blackjack if it's the first two cards dealt
	if len(h.Cards) == 2 && value == 21 && h.IsHidden == false {
		fmt.Println("BLACKJACK")
	}
	
    return value
}

func (h Hand) Hit() (Hand, error) {
	h.Cards = h.Cards.DrawNCards(1)
	h.Value = h.getValue()

	var err error
	if h.Value > 21 {
		err = errors.New("BUST!")
	}
	return h, err
}

func ResolveHand(playerHand Hand, dealerHand Hand) {
	fmt.Printf("Player Value: %d\n", playerHand.getValue())
	fmt.Printf("Dealer Value: %d\n", dealerHand.getValue())
}