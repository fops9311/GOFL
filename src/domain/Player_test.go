package domain_test

import (
	dom "GOFL/src/domain"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	p1 := dom.NewPlayer()

	if p1 == nil {
		t.Fatalf("p1 should not be nil: %v", p1)
	}
}
func TestPlayerNewAction(t *testing.T) {
	p1 := dom.NewPlayer()
	actionChan := p1.ActionChannel()
	go func() {
		p1.NewAction(2)
	}()
	action := <-*actionChan
	if action.(int) != 2 {
		t.Fatalf("action.Type should == 2: %v", action.(int))
	}
}
