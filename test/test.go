package test

import "testing"

import (
	"ds/stack"
	"ds/linklist"
	"fmt"
)

func TestStack(t *testing.T) {
	st := stack.NewStack()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)
	st.Push(5)
	t.Log(st)
	t.Log(st.Pop())
	t.Log(st)
}

func TestLinklist(t *testing.T) {
	ll := linklist.LinkList{}
	ll.Append(linklist.NewNode(1))
	ll.Append(linklist.NewNode(2))
	ll.Append(linklist.NewNode(3))

	ll.Val()

	p1 := ll.Pop()
	fmt.Printf("弹出的元素是%v",p1)

	ll.Val()

	ll.Append(linklist.NewNode(69))

	ll.Val()
	//t.Log(ll.Val())
	//t.Log(st.Pop())
	//t.Log(st)
}
