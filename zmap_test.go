package aggregate

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/chenhuaying/container"
)

type KeyUint uint32

func (k KeyUint) Less(x container.Comparer) bool {
	return k < x.(KeyUint)
}

func lessFn(x, y container.Comparer) bool {
	return x.(KeyUint) < y.(KeyUint)
}

func TestZmapNew(t *testing.T) {
	z := NewZmap()
	if z.IsEmpty() != true {
		t.Error("new must be empty")
	}

	if z.First() != nil {
		t.Error("new Zmap first must be nil")
	}
}

func TestZmapIteratorBeginAndEnd(t *testing.T) {
	z := NewZmap()
	z.Insert(KeyUint(3), "test-3")
	for j := 0; j < 10; j++ {
		key := 10 + j
		value := fmt.Sprintf("test-%d", key)
		z.Insert(KeyUint(key), value)
	}
	z.Insert(KeyUint(25), "test-25")
	z.Insert(KeyUint(35), "test-35")

	i := z.Find(KeyUint(3))
	fmt.Printf("is Begin %v, %d %v\n", i.Begin(), i.First(), i.Second())
	if !i.Begin() {
		t.Error("3 is the begin, why this case failed?")
	}

	i = z.Find(KeyUint(21))
	if !i.End() {
		t.Error("21 is not in the Zmap, why iterator not at end?")
	}
}

func TestZmapIterator(t *testing.T) {
	z := NewZmap()
	z.Insert(KeyUint(3), "test-3")
	for j := 0; j < 10; j++ {
		key := 10 + j
		value := fmt.Sprintf("test-%d", key)
		z.Insert(KeyUint(key), value)
	}
	z.Insert(KeyUint(25), "test-25")
	z.Insert(KeyUint(35), "test-35")

	//for i := z.Iterator(); !i.End(); i.Next() {
	//	fmt.Println(i.First(), i.Second())
	//}

	i := z.Find(KeyUint(13))
	// must use type assert, otherwise it will be failed
	if i.First().(KeyUint) != 13 || i.Second().(string) != "test-13" {
		t.Error("Zmap find error node!", i.First(), i.Second())
	}

	dest := []uint32{13, 14, 15, 16, 17, 18, 19, 25, 35}
	for j, c := i, 0; !j.End(); j.Next() {
		if uint32(i.First().(KeyUint)) != dest[c] {
			t.Error("next iterator error")
		}
		c += 1
	}

	dest = []uint32{13, 12, 11, 10, 3}
	for j, c := z.Find(KeyUint(13)), 0; !j.End(); j.Prev() {
		if uint32(j.First().(KeyUint)) != dest[c] {
			t.Error("next iterator error")
		}
		c += 1
	}

	dest = []uint32{25, 35}
	for j, c := z.LowerBound(KeyUint(20)), 0; !j.End(); j.Next() {
		if uint32(j.First().(KeyUint)) != dest[c] {
			t.Error("prev iterator error")
		}
		c++
	}

	i = z.LowerBound(KeyUint(40))
	if !i.End() {
		t.Error("find 40 node?")
	}
}

func TestZmapIteratorAssignment(t *testing.T) {
	z := NewZmap()
	for j := 0; j < 10; j++ {
		key := 10 + j
		value := fmt.Sprintf("test-%d", key)
		z.Insert(KeyUint(key), value)
	}

	itr := z.LowerBound(KeyUint(5))

	itr2 := (itr).(*ZmapIterator).IteratorDup()
	if itr2.(*ZmapIterator).aggregate != itr.(*ZmapIterator).aggregate ||
		itr2.(*ZmapIterator).currNode != itr.(*ZmapIterator).currNode {
		t.Error("copy iterator of zmap failed")
	}

	if reflect.ValueOf(itr2) == reflect.ValueOf(itr) {
		t.Error("not dup iterator of zmap")
	}
}

func TestZmapLowerBoundFn(t *testing.T) {
	z := NewZmap()
	z.Insert(KeyUint(3), "test-3")
	for j := 0; j < 10; j++ {
		key := 10 + j
		value := fmt.Sprintf("test-%d", key)
		z.Insert(KeyUint(key), value)
	}
	z.Insert(KeyUint(25), "test-25")
	z.Insert(KeyUint(35), "test-35")

	dest := []uint32{25, 35}
	for j, c := z.LowerBound(KeyUint(20)), 0; !j.End(); j.Next() {
		if uint32(j.First().(KeyUint)) != dest[c] {
			t.Error("prev iterator error")
		}
		c++
	}
}
