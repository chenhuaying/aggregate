package aggregate

import (
	"github.com/chenhuaying/aggregate/iterator"
	"github.com/chenhuaying/container"
	"github.com/chenhuaying/container/skiplist"
)

type Zmap struct {
	*skiplist.SkipList
}

type ZmapIterator struct {
	currNode  *skiplist.SkipListNode
	aggregate *Zmap
}

func NewZmap() *Zmap {
	return &Zmap{skiplist.NewSkipList()}
}

func (m *Zmap) Iterator() iterator.Iterator {
	return &ZmapIterator{currNode: m.First(), aggregate: m}
}

func (s *Zmap) Find(key container.Comparer) iterator.Iterator {
	node := s.SearchNode(key)
	return &ZmapIterator{currNode: node, aggregate: s}
}

func (s *Zmap) LowerBound(key container.Comparer) iterator.Iterator {
	node := s.LowerBoundNode(key)
	return &ZmapIterator{currNode: node, aggregate: s}
}

func (s *Zmap) LowerBoundFn(key container.Comparer, fn func(x, y container.Comparer) bool) iterator.Iterator {
	node := s.LowerBoundNodeFn(key, fn)
	return &ZmapIterator{currNode: node, aggregate: s}
}

func (s *Zmap) Delete(key container.Comparer) interface{} {
	node := s.DeleteNode(key)
	if node != nil {
		return node.Value()
	} else {
		return nil
	}
}

func (i *ZmapIterator) IteratorDup() iterator.Iterator {
	return &ZmapIterator{currNode: i.currNode, aggregate: i.aggregate}
}

func (i *ZmapIterator) Next() {
	i.currNode = i.currNode.Next()
}

func (i *ZmapIterator) Prev() {
	i.currNode = i.currNode.Prev()
}

func (i *ZmapIterator) First() interface{} {
	if i.currNode != nil {
		return i.currNode.Key()
	} else {
		return nil
	}
}

func (i *ZmapIterator) Second() interface{} {
	if i.currNode != nil {
		return i.currNode.Value()
	} else {
		return nil
	}
}

func (i *ZmapIterator) Begin() bool {
	if i.currNode == i.aggregate.First() {
		return true
	} else {
		return false
	}
}

func (i *ZmapIterator) End() bool {
	if i.currNode == nil {
		return true
	} else {
		return false
	}
}
