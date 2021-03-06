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
	li.Len--
	if pos == 0 {
		n = li.Head
		li.Head = li.Head.Next
		return li
	}
	insPos := li.Head
	for i := 0; i < pos-1; i++ {
		if insPos.Next.Next == nil {
			break
		}
		insPos = insPos.Next
	}
	n.Val = insPos.Next.Val
	insPos.Next = insPos.Next.Next
	return li
}
