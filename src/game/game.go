package game

import (
	"blackjack/src/hand"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var playerHand hand.Hand
var dealerHand hand.Hand

var notBust bool
var activeGame bool
var playerQuit bool

func init() {
	resetScreen()

	fmt.Println("Welcome to the Blackjack table!")
}

func Start() {
	for playerQuit == false {
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
	hand.ShuffleDeck()

	notBust = true
	activeGame = true

	dealerHand = hand.DealDealer()
	dealerHand.Print()

	playerHand = hand.DealPlayer()
	playerHand.Print()

	playRound()

	if playerQuit == false {
		hand.ResolveHand(playerHand, dealerHand)
	} else {
		fmt.Println("*You run away from the table, leaving your wallet*")
	}
}

func resetScreen() {
	fmt.Println("\033[2J")
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
			return
		}
		prompt()
	}
	fmt.Println()
}

func dealerTurn() {
	fmt.Println("\n-----")
	fmt.Println("Dealer's Turn")

	// reveal face down card
	dealerHand.IsHidden = false
	dealerHand.Print()
	fmt.Println("-----\n")
}
