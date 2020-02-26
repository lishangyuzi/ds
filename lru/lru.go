package Lru

import (
	"ds/linklist"
)

type NodeVal struct {
	Key string
	Val string
}

type Lru struct {
	cap  int
	maps map[string]interface{}
	Ll   linklist.Base
}

func NewLru() *Lru {
	return &Lru{
		maps: make(map[string]interface{}),
		Ll:   linklist.NewLinkList(),
		cap:  5,
	}
}

func (this *Lru) Append(key string, val string) {
	length := len(this.maps)
	if length+1 > this.cap {
		last := this.Ll.Pop()
		lastItem := last.(NodeVal)
		delete(this.maps, lastItem.Key)
	}

	n := linklist.NewNode(NodeVal{key, val})
	this.maps[key] = val
	this.Ll.Append(n)
}

func (this *Lru) Get(key string) interface{} {
	val := this.maps[key]
	if val == nil {
		return nil
	}
	nv := NodeVal{key, val.(string)}
	n := this.Ll.Index(nv)
	this.Ll.Remove(n)
	this.Ll.Append(linklist.NewNode(nv))
	return val
}
