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
	for cur != nil {
		st.Push(cur)
		cur = cur.left
	}
	for st.Len() > 0 {
		n := st.Pop()
		nn, ok := n.(*node);
		if !ok {
			return
		}
		fmt.Println(nn.val)

		if nn.right != nil {
			MidStackView(nn.right, st)
		}
	}

}

func PreStackView(n *node, st *stack.Stack) {

	cur := n
	for cur != nil {
		st.Push(cur)
		fmt.Println(cur.val)
		cur = cur.left
	}
	for st.Len() > 0 {
		n := st.Pop()
		nn, ok := n.(*node);
		if !ok {
			return
		}

		if nn.right != nil {
			PreStackView(nn.right, st)
		}
	}

}

func PostStackView(n *node, st *stack.Stack) {

	cur := n
	for cur != nil {
		st.Push(cur)
		cur = cur.left
	}
	for st.Len() > 0 {
		n := st.Pop()
		nn, ok := n.(*node);
		if !ok {
			return
		}

		if nn.right != nil {

			PostStackView(nn.right, st)
		}
			fmt.Println(nn.val)



	}

}


func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	bt := createBinaryTree(0, arr)
	fmt.Printf("%x\n", bt)
	fmt.Println("preview")
	PreView(bt)

	fmt.Println("midview")
	MidView(bt)

	fmt.Println("backview")
	PostView(bt)

	fmt.Println("midviewSt")
	midSt := stack.NewStack()
	MidStackView(bt, midSt)

	fmt.Println("previewSt")
	preSt := stack.NewStack()
	PreStackView(bt, preSt)

	fmt.Println("postiewSt")
	postSt := stack.NewStack()
	PostStackView(bt, postSt)
}
