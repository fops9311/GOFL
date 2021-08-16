package domain

type Game interface {
	InitGame() Game
	AddPlayer(p Player, playerId string) Game
	StartGame() Game
	ApplyRules(act Action, playerId string) Game
	EndGame() Game
}

type Action interface{}

type TransferAction struct {
	from   string
	to     string
	cardId int
}

type FoolGame struct {
	players map[string]Player
	decks   map[string]Deck
}

func (fg FoolGame) InitGame() Game {
	fg.players = make(map[string]Player)
	fg.decks = make(map[string]Deck)
	return fg
}
func (fg FoolGame) AddPlayer(p Player, playerId string) Game {
	fg.players[playerId] = p
	fg.decks[playerId] = make(Deck, 0)
	return fg
}
func (fg FoolGame) StartGame() Game {
	return fg
}
func (fg FoolGame) EndGame() Game {
	return fg
}
func (fg FoolGame) ApplyRules(act Action, playerId string) Game {
	switch action := act.(type) {
	case TransferAction:
		fg.decks[action.to] = append(fg.decks[action.to], fg.decks[action.from][action.cardId])
		fg.decks[action.from] = append(fg.decks[action.from][:action.cardId], fg.decks[action.from][action.cardId+1:]...)
	}
	return fg
}
