package linklist

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

func NewNode(val interface{}) *Node {
	return &Node{data: val}
}

type LinkList struct {
	header *Node
	len    int
}

func NewLinkList() LinkList {
	return LinkList{}
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
	this.len += 1
}

func (this *LinkList) Pop() interface{}{
	if this.len==0 {
		return nil
	}else{
		pre := this.header
		for pre.next.next!=nil  {
			pre = pre.next
		}
		this.len--
		val := pre.next
		pre.next = nil
		return val.data
	}
}

func (this *LinkList) Index(in interface{}) int{
	idx := 0
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


func (this *LinkList) Insert(idx int, node *Node) bool {
	if this.len == 0 {
		return false
	} else if idx <= 0 {
		this.header = this.header.next
		return true
	}
	var num int
	pre := this.header
	for {
		if pre.next != nil {
			if num+1 == idx {
				tmp := pre.next
				pre.next = node
				pre.next.next = tmp
				break
			}
			pre = pre.next
			num++
		} else {
			break
		}
	}
	return false
}

func (this *LinkList) Remove(idx int) bool {
	if this.len <= 0 {
		return false
	} else if idx <= 0 {
		this.header = this.header.next
		return true
	}
	var num int
	pre := this.header
	for {
		if pre.next != nil {
			if num+1 == idx {
				pre.next = pre.next.next
				return true
			}
			pre = pre.next
			num++
		} else {
			break
		}
	}
	return false
}

func (this *LinkList) Val() {
	pre := this.header
	for {
		fmt.Println(pre.data)
		if pre.next != nil {
			pre = pre.next
		} else {
			return
		}
	}
}
