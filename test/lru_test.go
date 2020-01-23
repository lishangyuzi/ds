package test

import (
	"ds/Lru"
	"testing"
)

func TestLru(t *testing.T) {
	lru := Lru.NewLru()
	lru.Append("key1", "cool1")
	lru.Append("key2", "cool2")
	lru.Append("key3", "cool3")
	lru.Append("key4", "cool4")
	lru.Append("key5", "cool5")
	lru.Append("key6", "cool6")

	lru.Get("key3")
	lru.Get("key11")

	nn3 := Lru.NodeVal{Key: "key3", Val: "cool3"}
	nn6 := Lru.NodeVal{Key: "key6", Val: "cool6"}
	idx3 := lru.Ll.Index(nn3)
	idx6 := lru.Ll.Index(nn6)
	if idx3 != 4 && idx6 != 3{
		t.Error("nlu error")
	}
}
