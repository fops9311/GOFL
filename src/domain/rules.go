package domain

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
