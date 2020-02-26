package tree

import (
	"ds/stack"
	"fmt"
)

func PreRcView(n *node) {
	if n == nil {
		return
	}
	fmt.Println(n.val)
	PreRcView(n.left)
	PreRcView(n.right)
}

func MidRcView(n *node) {
	if n == nil {
		return
	}
	MidRcView(n.left)
	fmt.Println(n.val)
	MidRcView(n.right)
}

func PostRcView(n *node) {
	if n == nil {
		return
	}
	PostRcView(n.left)
	PostRcView(n.right)
	fmt.Println(n.val)
}

func PreView(n *node) {
	st := stack.NewStack()
	cur := n

	for st.Len() != 0 || cur != nil {

		if cur != nil {
			fmt.Println(cur.val)
			st.Push(cur)
			cur = cur.left
		} else {
			item := st.Pop()
			pop, ok := item.(*node)
			if !ok {
				return
			}
			cur = pop.right
		}

	}
}

func MidView(n *node) {
	st := stack.NewStack()
	cur := n

	for st.Len() != 0 || cur != nil {
		if cur != nil {
			st.Push(cur)
			cur = cur.left
		} else {
			item := st.Pop()
			pop, ok := item.(*node)
			if !ok {
				return
			}
			fmt.Println(pop.val)
			cur = pop.right
		}
	}
}

func PostView(n *node) {
	st := stack.NewStack()
	cur := n
	var pre *node

	for st.Len() != 0 || cur != nil {
		for cur != nil {
			st.Push(cur)
			cur = cur.left
		}

		popItem := st.Last()
		top := popItem.(*node)

		if (top.left == nil && top.right == nil) || (pre == top.left && top.right == nil) || pre == top.right {
			fmt.Println(top.val)
			st.Pop()
			pre = top
		} else {
			cur = top.right
		}
	}

}
