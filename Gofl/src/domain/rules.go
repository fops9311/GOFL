package domain

type FoolGame struct {
	players []Player
	deck    Deck
	hands   []Deck
	table   Deck
	out     Deck
}

//AddPlayer addes player to a game
func (fr FoolGame) AddPlayer(p Player) FoolGame {
	if fr.players == nil {
		fr.players = make([]Player, 0, 4)
	}
	fr.players = append(fr.players, p)
	return fr
}

//StartGame starts game
func (fr FoolGame) StartGame() FoolGame {
	return fr
}

type Player struct {
	actions *chan Action
}

func (p Player) ActionChannel() *chan Action {
	return p.actions
}
func (p Player) NewAction(a Action) {
	*p.actions <- a
}
func NewPlayer() *Player {
	ach := make(chan Action)
	p := &Player{actions: &ach}
	return p
}

type Action struct {
	Type int
}
