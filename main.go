package main

import (
	"fmt"
	"github.com/heathernfran/blackjack/blackjack"
	)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println("Winnings:", winnings)
}