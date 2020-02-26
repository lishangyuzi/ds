package linklist

type LinkList struct {
	header *Node
	len    int
}

func NewLinkList() *LinkList {
	return &LinkList{}
}

func (this *LinkList) Append(node *Node) {
	if this.len == 0 {
		this.header = node
	} else {
		pre := this.header
		for pre.next != nil {
			pre = pre.next
		}
		pre.next = node
	}
	this.len++
}

func (this *LinkList) Pop() interface{} {
	if this.len == 0 {
		return nil
	} else {
		pre := this.header
		for pre.next.next != nil {
			pre = pre.next
		}
		this.len--
		val := pre.next
		pre.next = nil
		return val.data
	}
}

func (this *LinkList) Index(in interface{}) int {
	idx := 0
	pre := this.header
	for pre != nil {
		if pre.data == in {
			return idx
		}
		pre = pre.next
		idx++
	}
	return -1
}

func (this *LinkList) Get(idx int) *Node {
	pre := this.header
	count := 0
	for count < idx {
		pre = pre.next
		count++
	}
	return pre
}

func (this *LinkList) Insert(idx int, node *Node) bool {
	if idx <= 0 || idx > this.len-1 {
		return false
	}

	preNode := this.Get(idx - 1)
	node.next = preNode.next.next
	preNode.next = node

	return true
}

func (this *LinkList) Remove(idx int) bool {
	if this.len <= 0 || idx <= 0 || idx > this.len-1 {
		return false
	}
	preNode := this.Get(idx - 1)
	preNode.next = preNode.next.next

	return true
}

func (this *LinkList) Val() []interface{} {
	pre := this.header
	res := []interface{}{}
	for {
		res = append(res, pre.data)
		if pre.next != nil {
			pre = pre.next
		} else {
			break
		}
	}
	return res
}
