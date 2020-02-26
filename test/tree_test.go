package test

import (
	"ds/tree"
	"fmt"
	"testing"
)

func TestCreateBinaryTree(t *testing.T) {
	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	newT := tree.CreateBinaryTree(source)
	// fmt.Println(newT.Root)
	// tree.PreRcView(newT.Root)
	// fmt.Println("-----------------")
	// tree.MidRcView(newT.Root)
	// fmt.Println("-----------------")
	// tree.PostRcView(newT.Root)

	fmt.Println("--------no recursion--------")
	// tree.PreView(newT.Root)
	tree.PostView(newT.Root)
}
