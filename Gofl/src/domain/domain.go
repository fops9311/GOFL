package domain

import (
	u "Gofl/src/util"
	"errors"
	"math/rand"
	"time"
)

type Suit int

type Face int

const (
	Ace Face = iota
	King
	Quin
	Jack
	Ten
	Nine
	Eight
	Seven
	Six
)
const (
	Harts Suit = iota
	Diamonds
	Spades
	Clubs
)

var Suits [4]Suit = [4]Suit{Harts, Diamonds, Spades, Clubs}
var Faces [9]Face = [9]Face{Ace, King, Quin, Jack, Ten, Nine, Eight, Seven, Six}

type Card struct {
	S Suit
	F Face
}

type Deck []Card

func (d Deck) Pull(num int) (pulled Deck, remain Deck, err error) {
	if num > len(d) {
		return []Card{}, d, errors.New("Not enough cards")
	}
	pulled, d = d[:num], d[num:]
	return pulled, d, nil
}
func (d Deck) Push(pushed Deck) Deck {
	d = append(pushed, d...)
	return d
}
func (d Deck) Shuffle() Deck {
	result := make(Deck, 0, len(d))
	list := u.List{}
	rand.Seed(time.Now().UnixNano())
	var rInt int
	for i := 0; i < len(d); i++ {
		if list.Len > 0 {
			rInt = rand.Intn(list.Len)
		}
		list = list.Insert(u.Node{Val: i}, rInt)

	}
	list.Traverse(func(V int) {
		result = append(result, d[V])
	})
	return result
}
func NewDeck() Deck {
	d := make(Deck, 0, len(Suits)*len(Faces)+10)
	for _, s := range Suits {
		for _, f := range Faces {
			d = append(d, Card{S: s, F: f})
		}
	}
	return d
}
