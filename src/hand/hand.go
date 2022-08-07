package hand

import (
	"blackjack/src/deck"
	"errors"
	"fmt"
)

type Hand struct {
	Cards    deck.Cards
	Value    int
	IsDealer bool
	IsHidden bool
}

var cardBackUnicode string

func init() {
	cardBackUnicode = "\U0001F0A0"
}

// this is just a lazy way to give game.go access to this deck.go function
func ShuffleDeck() {
	deck.ShuffleDeck()
}

func dealHand(dealer bool) Hand {
	var cards deck.Cards = []deck.Card{}

	h := Hand{
		Cards:    cards,
		IsDealer: dealer,
		IsHidden: dealer,
	}

	h.Cards = h.Cards.DrawNCards(2)
	h.Value = h.GetValue()
	return h
}

func DealPlayer() Hand {
	return dealHand(false)
}

func DealDealer() Hand {
	return dealHand(true)
}

func (h Hand) printDealer(hidden bool) {
	fmt.Print("\t\t[Dealer] ")

	startIndex := 0
	if h.IsHidden {
		startIndex = 1
		fmt.Printf("%s ", cardBackUnicode)
	}

	for _, item := range h.Cards[startIndex:] {
		fmt.Printf("%s ", item.GetCard())
	}

	fmt.Printf(" showing %d\n", h.Value)
}
func (h Hand) printPlayer() {
	fmt.Print("\t\t[Player] ")
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
		h.printDealer(h.IsHidden)
	} else {
		h.printPlayer()
	}
}
func (h Hand) GetValue() int {
	value := 0
	aceCount := 0

	startIndex := 0
	if h.IsHidden == true {
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

	return value
}

func (h Hand) Hit() (Hand, error) {
	h.Cards = h.Cards.DrawNCards(1)
	h.Value = h.GetValue()

	var err error
	if h.Value > 21 {
		err = errors.New("BUST!")
	}
	return h, err
}
