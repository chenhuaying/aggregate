package aggregate

import (
	"fmt"
	"testing"
)

func TestZsetNew(t *testing.T) {
	z := NewZset()
	if z.IsEmpty() != true {
		t.Error("new must be empty")
	}

	if z.First() != nil {
		t.Error("new zset first must be nil")
	}
}

func TestZsetIterator(t *testing.T) {
	z := NewZset()
	z.Insert(3, "test-3")
	for j := 0; j < 10; j++ {
		key := 10 + j
		value := fmt.Sprintf("test-%d", key)
		z.Insert(uint32(key), value)
	}
	z.Insert(25, "test-25")
	z.Insert(35, "test-35")

	//for i := z.Iterator(); !i.End(); i.Next() {
	//	fmt.Println(i.First(), i.Second())
	//}

	i := z.Find(13)
	// must use type assert, otherwise it will be failed
	if i.First().(uint32) != 13 || i.Second().(string) != "test-13" {
		t.Error("zset find error node!", i.First(), i.Second())
	}

	dest := []uint32{13, 14, 15, 16, 17, 18, 19, 25, 35}
	for j, c := i, 0; !j.End(); j.Next() {
		if i.First().(uint32) != dest[c] {
			t.Error("next iterator error")
		}
		c += 1
	}

	dest = []uint32{13, 12, 11, 10, 3}
	for j, c := z.Find(13), 0; !j.End(); j.Prev() {
		if j.First().(uint32) != dest[c] {
			t.Error("next iterator error")
		}
		c += 1
	}

	dest = []uint32{25, 35}
	for j, c := z.LowerBound(20), 0; !j.End(); j.Next() {
		if j.First().(uint32) != dest[c] {
			t.Error("prev iterator error")
		}
		c++
	}

	i = z.LowerBound(40)
	if !i.End() {
		t.Error("find 40 node?")
	}
}
