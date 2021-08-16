package util

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	Head *Node
	Len  int
}

//Insert into position. If position > Len, appends
func (li List) Insert(n Node, pos int) List {
	li.Len++
	if li.Head == nil {
		li.Head = &n
		return li
	}
	if pos == 0 {
		n.Next = li.Head
		li.Head = &n
		return li
	}
	insPos := li.Head
	for i := 0; i < pos; i++ {
		if insPos.Next == nil {
			insPos.Next = &n
			return li
		}
		insPos = insPos.Next
	}
	n.Next = insPos.Next
	insPos.Next = &n
	return li
}
func (li List) Traverse(f func(V int)) {
	n := li.Head
	for n != nil {
		f(n.Val)
		n = n.Next
	}
}
func (li List) Detach(pos int, n *Node) List {
	if pos == 0 {
		li.Head = n
		return li
	}
	insPos := li.Head
	for i := 0; i < pos; i++ {
		insPos = insPos.Next
		if insPos == nil {
			return li
		}
	}
	n.Val=insPos.Val
	return li
}
