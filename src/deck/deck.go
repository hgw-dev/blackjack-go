package deck

import (
	"fmt"
	"math/rand"
	"time"
	"errors"
	"strings"
)

var rankNameMap map[int]string

var suitUnicodeMap map[string]string

var deck []Card

type Card struct {  
    Suit  string
	Rank  int
}

// Unicode for Cards
// https://en.wikipedia.org/wiki/Playing_cards_in_Unicode#Playing_Cards_(block)
func (c Card) Print() {
	rankName := rankToName(c.Rank)
	rankValue := c.Rank
	if (rankValue > 11){
		rankValue += 1
	}
	
	// base value -> rune -> int
	uniCodePoint := int([]rune(suitUnicodeMap[c.Suit])[0])
	// base value as int + rank value
	offsetUniCodePoint := string(uniCodePoint + rankValue)

	fmt.Printf(
		"[%s ] %s of %s\n", 
		offsetUniCodePoint,
		strings.Title(rankName), 
		strings.Title(c.Suit),
	)
}

func rankToName(index int) string {
	return rankNameMap[index]
}

func init(){
	rankNameMap = map[int]string{
		1: "ace", 
		2: "two", 3: "three", 4: "four", 
		5: "five", 6: "six", 7: "seven", 
		8: "eight", 9: "nine", 10: "ten", 
		11: "jack", 12: "queen", 13: "king",
	}
	suits := []string {
		"clubs", 
		"spades", 
		"hearts", 
		"diamonds",
	}
	// base unicode values, we add the rank to get the card
	suitUnicodeMap = map[string]string {
		"clubs": "\U0001F0D0",
		"spades": "\U0001F0A0",
		"hearts": "\U0001F0B0",
		"diamonds": "\U0001F0C0",
	}

	// fill deck
	for s := range suits {
		for r := 1; r <= 13; r++ {
			c := Card{
				Rank: r,
				Suit: suits[s],
			}
			deck = append(deck, c)
		}
	}
	
	// shuffle
	rand.Seed(time.Now().UnixNano())
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
}

func Draw() (first Card){
	var err error
	if len(deck) > 1{
		first = deck[0]
		deck = deck[1:]
	} else {
		err = errors.New("Out of cards!")
	}
		
	if err == nil {
		first.Print()
	} else {
		fmt.Println(err)
	}
	return first
}