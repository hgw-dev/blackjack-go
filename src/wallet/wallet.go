package wallet

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
)

type Wallet struct {
	Amount      uint64
	Description string
	CurrentBet  uint64
}

var adjectives map[string][]string
var defaultAmount uint64

func init() {
	// $500
	// We'll do math in cents and convert for display
	defaultAmount = 50000

	adjectives = map[string][]string{
		"color":    {"brown", "black"},
		"material": {"velcro", "duct tape", "leather"},
		"style":    {"lovely", "exquisite", "stylish", "cheap", "tattered"},
	}
}

func New() Wallet {
	walletColor := adjectives["color"][rand.Intn(len(adjectives["color"]))]
	walletMaterial := adjectives["material"][rand.Intn(len(adjectives["material"]))]
	walletStyle := adjectives["style"][rand.Intn(len(adjectives["style"]))]

	description := string(walletStyle + ", " + walletColor + " " + walletMaterial)

	w := Wallet{
		Amount:      defaultAmount,
		Description: description,
		CurrentBet:  uint64(1),
	}

	return w
}

func (w Wallet) GetAmount() string {
	return w.AmountUintToString(w.Amount)
}

func (w Wallet) GetCurrentBet() string {
	return w.AmountUintToString(w.CurrentBet)
}

func (w Wallet) AmountUintToString(amount uint64) string {
	result := strconv.FormatUint(amount, 10)

	result = result[:len(result)-2] + "." + result[len(result)-2:]

	return result
}

func AmountStringToUint(amount string) (uint64, error) {
	idx := int(strings.Index(amount, "."))
	if idx != -1 {
		// 123.4 -> idxDiff = 2
		// 123.45 -> idxDiff = 3 <-- ideal
		// 123.456 -> idxDiff = 4
		idxDiff := len(amount) - idx

		if idxDiff == 2 {
			amount = amount + "0"
		} else if idxDiff > 3 {
			amount = amount[:idx+3]
		}
		amount = amount[:idx] + amount[idx+1:]
	} else {
		amount = amount + "00"
	}

	return strconv.ParseUint(amount, 10, 64)
}

func (w Wallet) PlaceBet(amount uint64) Wallet {
	w.CurrentBet = amount
	return w
}

func (w Wallet) BetWon() Wallet {
	w.Amount += w.CurrentBet
	return w
}

func (w Wallet) BetLost() (Wallet, error) {
	var err error
	if w.Amount <= w.CurrentBet {
		err = errors.New("You ran out of money!")
	}
	w.Amount -= w.CurrentBet
	return w, err
}
