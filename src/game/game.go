package game

import (
	"fmt"
    // "bufio"
    // "os"
    // "strings"
)
import "blackjack/src/hand"

func init(){
	fmt.Println("Welcome to the Blackjack table!")
	fmt.Println("Shuffling...\n")
}

func Start(){
    d := hand.DealDealer()
    d.Print()

    p := hand.DealPlayer()
    p.Print()

    // deal()
}

// func prompt() {
//     fmt.Print("> ")
// }

// func printUnknown(text string) {
//     fmt.Println(text, ": command not found")
// }
 
// func displayHelp() {
//     fmt.Println("help    - Show available commands")
// }
 
// func handleInvalidCmd(text string) {
//     defer printUnknown(text)
// }
 
// func deal() {
//     commands := map[string]interface{}{
//         "help": displayHelp,
//     }
//     reader := bufio.NewScanner(os.Stdin)
//     prompt()
//     for reader.Scan() {
//         text := strings.ToLower(strings.TrimSpace(reader.Text()))
//         if command, exists := commands[text]; exists {
//             command.(func())()
//         } else if strings.EqualFold("exit", text) {
//             return
//         } else if text != "" {
//             handleInvalidCmd(text)
//         }
//         prompt()
//     }
//     fmt.Println()
// }