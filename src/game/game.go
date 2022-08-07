package game

import (
	"fmt"
    "bufio"
    "os"
    "strings"
)
import "blackjack/src/hand"

var playerHand hand.Hand
var dealerHand hand.Hand

var notBust bool
var activeGame bool

func init(){
    notBust = true
    activeGame = true

	fmt.Println("Welcome to the Blackjack table!")
	fmt.Println("Shuffling...\n")
}

func Start(){
    dealerHand = hand.DealDealer()
    dealerHand.Print()

    playerHand = hand.DealPlayer()
    playerHand.Print()

    playHand()

	hand.ResolveHand(playerHand, dealerHand)
}

func prompt() {
    fmt.Print("> ")
}

func stand() {
    fmt.Println("\"I'll stand...\"\n")
    activeGame = false
}

func hit() {
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
 
func playHand() {
    commands := map[string]interface{}{
        "help": displayHelp,
        "stand": stand,
        "s": stand,
        "hit": hit,
        "h": hit,
    }
    reader := bufio.NewScanner(os.Stdin)
    prompt()
    for reader.Scan() {
        text := strings.ToLower(strings.TrimSpace(reader.Text()))
        if command, exists := commands[text]; exists {
            command.(func())()
        } else if (
            strings.EqualFold("exit", text) || 
            strings.EqualFold("e", text) || 
            strings.EqualFold("quit", text) || 
            strings.EqualFold("q", text) ){
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