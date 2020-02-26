package linklist

type DuLinklist struct{
	header *Node
	tail *Node
	len int
}

func NewDuLink() *DuLinklist{
	return &DuLinklist{}
}

func (this *DuLinklist) Append(node *Node){
	if this.len == 0{
		this.header , this.tail = node , node
	}
	this.tail.next = node
	node.pre  = this.tail
	this.tail = node
	this.len ++
}

func (this *DuLinklist) Pop() (item interface{}){
	item = this.tail.data
	this.tail = this.tail.pre
	this.tail.next = nil
	this.len --
	return 
}

func (this *DuLinklist) Index( in interface{}) int{
	idx  :=  0
	pre := this.header
	for pre != nil {
		if pre.data == in {
			return idx
		}
		pre = pre.next
		idx ++
	}
	return -1 
}

func (this *DuLinklist) Get(idx int) *Node{
	pre := this.header
	count := 0
	for count < idx {
		pre = pre.next
		count++
	}
	return pre
}

func (this *DuLinklist) Insert(idx int, node *Node) bool {
	if idx <= 0 || idx > this.len-1 {
		return false
	}

	preNode := this.Get(idx - 1)
	node.next = preNode.next.next
	preNode.next = node

	return true
}

func (this *DuLinklist) Remove(idx int) bool {
	if this.len <= 0 || idx <= 0 || idx > this.len-1 {
		return false
	}
	preNode := this.Get(idx - 1)
	preNode.next = preNode.next.next

	return true
}

func (this *DuLinklist) Val() []interface{} {
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


