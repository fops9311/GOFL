package main

import (
	dom "Gofl/src/domain"
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	deck := dom.NewDeck()
	deck = deck.Shuffle()
	hand, deck, _ := deck.Pull(5)
	_ = hand
	fmt.Println(time.Since(startTime))
	//fmt.Println(deck)
	//hand, deck, _ := deck.Pull(5)
	//fmt.Println(hand)
	//fmt.Println(deck)
	//hand2, hand, _ := hand.Pull(5)
	//fmt.Println(hand)
	//fmt.Println(hand2)
}
