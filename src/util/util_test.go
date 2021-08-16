package util_test

import (
	u "GOFL/src/util"
	"fmt"
	"testing"
)

func TestDetach(t *testing.T) {
	list := u.List{}
	s := []int{1, 2, 3, 4, 5, 6}
	for _, v := range s {
		list = list.Insert(u.Node{Val: v}, list.Len)
	}
	list.Traverse(func(v int) {
		t.Log(fmt.Sprint(v) + "->")
	})

	list = list.Detach(0, &u.Node{})
	list.Traverse(func(v int) {
		t.Log(fmt.Sprint(v) + "->")
	})

	list = list.Detach(1, &u.Node{})
	list.Traverse(func(v int) {
		t.Log(fmt.Sprint(v) + "->")
	})
	list = list.Detach(3, &u.Node{})
	list.Traverse(func(v int) {
		t.Log(fmt.Sprint(v) + "->")
	})
	list = list.Detach(3, &u.Node{})
	list.Traverse(func(v int) {
		t.Log(fmt.Sprint(v) + "->")
	})
	t.Log(list.Len)
}
