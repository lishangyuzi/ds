package main

import (
	"ds/linklist"
	"fmt"
)

type NodeVal struct {
	key string
	val string
}

type Lru struct {
	cap  int
	maps map[string]interface{}
	ll   linklist.LinkList
}

func NewLru() *Lru{
	return &Lru{
		maps: make(map[string]interface{}),
		ll: linklist.NewLinkList(),
	}
}

func (this *Lru) append(key string,val string){
	//len := len(this.maps)
	// todo

	n := linklist.NewNode(NodeVal{key,val})
	this.maps[key] = val
	//this.cap++
	this.ll.Append(n)
}

func(this *Lru) get(key string) interface{}{
  	val := this.maps[key]
  	nv := NodeVal{key,val.(string)}
  	n:= this.ll.Index(nv)
  	this.ll.Remove(n)
  	this.ll.Append(linklist.NewNode(nv))
  	return val
}

func main() {
	lru := NewLru()
	lru.append("key1","cool1")
	lru.append("key2","cool2")
	lru.append("key3","cool3")
	lru.append("key4","cool4")
	lru.append("key5","cool5")
	//lru.ll.Val()
	lru.get("key3")
	//lru.ll.Index("key2")
	lru.ll.Val()

	fmt.Println(lru)
}
