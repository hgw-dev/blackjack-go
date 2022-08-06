module blackjack/main

go 1.13

require blackjack/game v0.0.0

replace blackjack/game => ./game

replace blackjack/deck => ./deck
