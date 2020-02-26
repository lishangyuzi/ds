package test

import (
	"ds/linklist"
	"testing"
)

func TestLinklist(t *testing.T) {
	var ll linklist.Base
	ll = linklist.NewLinkList()
	ll.Append(linklist.NewNode(1))
	ll.Append(linklist.NewNode(2))
	ll.Append(linklist.NewNode(3))

	p1 := ll.Pop()
	t.Logf("pop item is %v", p1)

	ll.Append(linklist.NewNode(69))

	ll.Insert(1, linklist.NewNode(888))
	ll.Remove(2)
	arr := ll.Val()
	t.Log(arr)
}

func TestDuLinklist(t *testing.T) {
	var ll linklist.Base
	ll = linklist.NewDuLink()
	ll.Append(linklist.NewNode(1))
	ll.Append(linklist.NewNode(2))
	ll.Append(linklist.NewNode(3))

	p1 := ll.Pop()
	t.Logf("pop item is %v", p1)

	ll.Append(linklist.NewNode(69))

	ll.Insert(1, linklist.NewNode(888))
	ll.Remove(2)
	arr := ll.Val()
	t.Log(arr)
}
