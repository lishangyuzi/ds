package linklist

type Node struct {
	data interface{}
	next *Node
	pre  *Node
}

func NewNode(val interface{}) *Node {
	return &Node{data: val}
}

type Base interface {
	Append(node *Node)
	Pop() interface{}
	Index(in interface{}) int
	Get(idx int) *Node
	Insert(idx int, node *Node) bool
	Remove(idx int) bool
	Val() []interface{}
}
