package domain

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
