package main

import (
	"fmt"
	"ds/stack"
)

type node struct {
	val   int
	left  *node
	right *node
}

type binaryTree struct {
	root *node
}

func createBinaryTree(index int, arr []int) *node {
	node := &node{}
	if index >= len(arr) {
		return nil
	}
	node.val = arr[index]
	node.left = createBinaryTree(2*index+1, arr)
	node.right = createBinaryTree(2*index+2, arr)
	return node
}

func PreView(n *node) {
	fmt.Println(n.val)
	if n.left != nil {
		PreView(n.left)
	}

	if n.right != nil {
		PreView(n.right)
	}
}

func MidView(n *node) {
	if n.left != nil {
		MidView(n.left)
	}

	fmt.Println(n.val)

	if n.right != nil {
		MidView(n.right)
	}
}

func PostView(n *node) {

	if n.left != nil {
		PostView(n.left)
	}

	if n.right != nil {
		PostView(n.right)
	}

	fmt.Println(n.val)
}

func MidStackView(n *node, st *stack.Stack) {
	cur := n
	res := []int{}
	for cur != nil || st.Len() != 0 {
		if cur != nil {
			st.Push(cur)
			cur = cur.left
		} else {
			n := st.Pop()
			nn, ok := n.(*node);
			if !ok {
				return
			}
			res = append(res, nn.val)
			cur = nn.right;
		}
	}
	fmt.Println(res)
}

func PreStackView(n *node, st *stack.Stack) {
	cur := n
	res := []int{}
	for cur != nil || st.Len() != 0 {
		if cur != nil {
			st.Push(cur)
			res = append(res, cur.val)
			cur = cur.left
		} else {
			n := st.Pop()
			nn, ok := n.(*node);
			if !ok {
				return
			}
			cur = nn.right;
		}
	}
	fmt.Println(res)
}

func PostStackView(n *node, st *stack.Stack) {
	cur := n
	res := []int{}
	pre := &node{}
	st.Push(cur)
	for cur != nil || st.Len() != 0 {
		for cur != nil {
			st.Push(cur)
			cur = cur.left
		}

		ns := st.Last()
		top := ns.(*node)

		if (top.left == nil && top.right == nil) || (pre == top.left && top.right == nil) || pre == top.right {
			fmt.Println(top.val)
			res = append(res, top.val)
			pre = top
			_ = st.Pop()
		} else {
			cur = top.right
		}
	}
	fmt.Println(res)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	bt := createBinaryTree(0, arr)
	fmt.Printf("%x\n", bt)

	//fmt.Println("preview")
	//PreView(bt)

	//fmt.Println("midview")
	//MidView(bt)

	//fmt.Println("backview")
	//PostView(bt)

	//fmt.Println("midviewSt")
	//midSt := stack.NewStack()
	//MidStackView(bt, midSt)

	//fmt.Println("previewSt")
	//preSt := stack.NewStack()
	//PreStackView(bt, preSt)

	fmt.Println("postiewSt")
	postSt := stack.NewStack()
	PostStackView(bt, postSt)
}
