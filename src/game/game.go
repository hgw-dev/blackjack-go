package game

import (
	"blackjack/src/hand"
	"blackjack/src/wallet"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var playerWallet wallet.Wallet

var playerHand hand.Hand
var dealerHand hand.Hand

var notBust bool
var activeGame bool
var playerQuit bool
var emptyWallet bool

func init() {
	playerWallet = wallet.New()
	resetScreen()

	fmt.Printf(
		"*You approach a table holding a %s wallet with %s doubloons*\n",
		playerWallet.Description, playerWallet.GetAmount(),
	)
	fmt.Println("\"Welcome to the Blackjack table!\"")
}

func Start() {
	for playerQuit == false && emptyWallet == false {
		for {
			reader := bufio.NewScanner(os.Stdin)
			fmt.Println("\nWould you like to play a hand?")
			prompt()

			reader.Scan()
			text := strings.ToLower(strings.TrimSpace(reader.Text()))
			if text == "yes" || text == "y" {
				resetScreen()
				fmt.Println("\nShuffling...\n")
				playHand()
				break
			} else {
				return
			}
		}
	}
}

func playHand() {
	placeBet()

	hand.ShuffleDeck()

	notBust = true
	activeGame = true
	emptyWallet = false

	dealerHand = hand.DealDealer()
	dealerHand.Print()

	playerHand = hand.DealPlayer()
	playerHand.Print()

	if playerHand.Value == 21 {
		fmt.Println("BLACKJACK")
		activeGame = false
		resolveHand()
		return
	}

	playRound()

	if playerQuit == true {
		fmt.Printf("*You run away from the table, leaving your %s wallet and your %s doubloons*",
			playerWallet.Description, playerWallet.GetAmount(),
		)
	} else {
		resolveHand()

		if emptyWallet == true {
			fmt.Printf("*You pretend to dig in your wallet for more doubloons, before meakly walking away*")
		}
	}
}

func placeBet() {
	promptForBet := true
	var amount uint64 = 0
	for promptForBet {
		reader := bufio.NewScanner(os.Stdin)
		fmt.Println("\nPlace your bet")

		fmt.Print(playerWallet.GetAmount() + " ")
		prompt()

		reader.Scan()

		text := strings.ToLower(strings.TrimSpace(reader.Text()))
		var err error
		amount, err = wallet.AmountStringToUint(text)

		if err != nil {
			fmt.Println("\nInvalid bet! Try again")
		} else {
			promptForBet = false
		}
	}
	playerWallet = playerWallet.PlaceBet(uint64(amount))
}

func resolveHand() {
	dealerHand.IsHidden = false
	dealerHand.Value = dealerHand.GetValue()

	dealerHand.Print()
	playerHand.Print()
	fmt.Println("")

	dealerBust := dealerHand.Value > 21
	playerBust := playerHand.Value > 21
	playerBeatDealer := playerHand.Value > dealerHand.Value

	if playerBust == false && ((playerBeatDealer && dealerBust == false) || dealerBust) {
		fmt.Printf("<<< YOU WIN %s doubloons! >>>\n", playerWallet.GetCurrentBet())
		fmt.Println("*Dopamine surges throughout your brain. You enjoy gambling.*")
		fmt.Println("-----")

		playerWallet = playerWallet.BetWon()
	} else {
		fmt.Printf("<<< YOU LOSE %s doubloons! LOSER! >>>\n", playerWallet.GetCurrentBet())
		fmt.Println("*You have an overwhelming urge to play another hand*")
		fmt.Println("-----")

		var emptyWalletError error
		playerWallet, emptyWalletError = playerWallet.BetLost()
		if emptyWalletError != nil {
			fmt.Println(emptyWalletError)
			emptyWallet = true
		}
	}
}

func resetScreen() {
	fmt.Println("\n\n\n")
	// fmt.Println("\033[2J")
}

func prompt() {
	fmt.Print("> ")
}

func stand() {
	resetScreen()
	fmt.Println("\"I'll stand...\"")
	activeGame = false
}

func hit() {
	resetScreen()
	fmt.Println("\"Hit me!\"\n")
	var bustError error
	playerHand, bustError = playerHand.Hit()

	dealerHand.Print()
	playerHand.Print()
	if bustError != nil {
		fmt.Println(bustError)
		notBust = false
		activeGame = false
	} else if playerHand.Value == 21 {
		activeGame = false
	}
}

func printUnknown(text string) {
	if text != "" {
		fmt.Println(text, ": command not found\n")
	}
	fmt.Println("Available options are \"hit\" and \"stand\".")
	fmt.Println("Use \"help\" for additional details")
	fmt.Println("Use \"quit\" to end the game.")
}

func displayHelp() {
	resetScreen()
	fmt.Println("The Goal:")
	fmt.Println("\tBeat the dealer")
	fmt.Println("\tby getting a score as close to 21 as possible,")
	fmt.Println("\twithout going over.\n")
	fmt.Println("hit (h)    - Ask the dealer for another card")
	fmt.Println("stand (s)  - Stay with your hand")
	fmt.Println("\nhelp       - Show this message again")
	fmt.Println("quit (q)   - End the game")
}

func handleInvalidCmd(text string) {
	defer printUnknown(text)
}

func playRound() {
	playerTurn()

	if playerQuit == false {
		dealerTurn()
	}
}

func playerTurn() {
	commands := map[string]interface{}{
		"help":  displayHelp,
		"stand": stand,
		"s":     stand,
		"hit":   hit,
		"h":     hit,
	}
	fmt.Println("\n<Players's Turn>")

	reader := bufio.NewScanner(os.Stdin)
	prompt()
	for {
		reader.Scan()
		text := strings.ToLower(strings.TrimSpace(reader.Text()))
		if command, exists := commands[text]; exists {
			command.(func())()
		} else if strings.EqualFold("exit", text) ||
			strings.EqualFold("e", text) ||
			strings.EqualFold("quit", text) ||
			strings.EqualFold("q", text) {
			playerQuit = true
			return
		} else {
			handleInvalidCmd(text)
		}
		if activeGame == false {
			fmt.Println("</Players's Turn>")
			return
		}
		prompt()
	}
	fmt.Println()
}

func dealerTurn() {
	fmt.Println("\n<Dealer's Turn>")

	// reveal face down card
	dealerHand.IsHidden = false
	dealerActiveGame := true

	dealerHand.Value = dealerHand.GetValue()
	dealerHand.Print()
	fmt.Println("")

	for dealerActiveGame {
		dealerValue := dealerHand.Value

		if dealerValue >= 17 {
			dealerActiveGame = false
		} else {
			dealerHand, _ = dealerHand.Hit()

			dealerHand.Print()
			if dealerHand.Value >= 21 {
				dealerActiveGame = false
			}
		}
	}
	fmt.Println("</Dealer's Turn>\n")
}
