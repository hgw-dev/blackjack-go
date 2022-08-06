module blackjack/main

go 1.13

require (
	blackjack/src/deck v0.0.0-00010101000000-000000000000 // indirect
	blackjack/src/game v0.0.0
)

replace blackjack/src/game => ./game

replace blackjack/src/deck => ./deck

replace blackjack/src/hand => ./hand
