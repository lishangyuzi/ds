package stack

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	return new(Stack)
}

func (this *Stack) Len() int {
	return len(this.data)
}

func (this *Stack) Push(node interface{}) {
	this.data = append(this.data, node)
}

func (this *Stack) Pop() interface{} {
	len := len(this.data)
	if len == 0 {
		return nil
	}
	node := this.data[len-1]
	this.data = this.data[0 : len-1]
	return node
}

