package test

import (
	"testing"
	"ds/linklist"
)

func TestLinklist(t *testing.T) {
	ll := linklist.LinkList{}
	ll.Append(linklist.NewNode(1))
	ll.Append(linklist.NewNode(2))
	ll.Append(linklist.NewNode(3))

	p1 := ll.Pop()
	t.Logf("pop item is %v",p1)

	ll.Append(linklist.NewNode(69))

	arr := ll.Val()
	t.Log(arr)
}
