package tree

type BinaryTree struct {
	Root *node
	Size int
}

type node struct {
	val   int
	left  *node
	right *node
}

func buildBinaryTree(index int, arr []int) *node {
	node := &node{}
	if index >= len(arr) {
		return nil
	}
	node.val = arr[index]
	node.left = buildBinaryTree(2*index+1, arr)
	node.right = buildBinaryTree(2*index+2, arr)
	return node
}

func CreateBinaryTree(arr []int) *BinaryTree {
	return &BinaryTree{
		Root: buildBinaryTree(0, arr),
		Size: 1,
	}
}


