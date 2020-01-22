package test

import (
	"testing"
	"ds/stack"
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
