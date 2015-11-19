package aggregate

import (
	"github.com/user/aggregate/iterator"
	"github.com/user/skiplist"
)

type Zset struct {
	*skiplist.SkipList
}

type ZsetIterator struct {
	currNode  *skiplist.SkipListNode
	aggregate *Zset
}

func NewZset() *Zset {
	return &Zset{skiplist.NewSkipList()}
}

func (s *Zset) Iterator() iterator.Iterator {
	return &ZsetIterator{currNode: s.First(), aggregate: s}
}

func (s *Zset) Find(key uint32) iterator.Iterator {
	node := s.SearchNode(key)
	return &ZsetIterator{currNode: node, aggregate: s}
}

func (s *Zset) LowerBound(key uint32) iterator.Iterator {
	node := s.LowerBoundNode(key)
	return &ZsetIterator{currNode: node, aggregate: s}
}

func (i *ZsetIterator) Next() {
	i.currNode = i.currNode.Next()
}

func (i *ZsetIterator) Prev() {
	i.currNode = i.currNode.Prev()
}

func (i *ZsetIterator) End() bool {
	if i.currNode == nil {
		return true
	} else {
		return false
	}
}

func (i *ZsetIterator) First() interface{} {
	if i.currNode != nil {
		return i.currNode.Key()
	} else {
		return nil
	}
}

func (i *ZsetIterator) Second() interface{} {
	if i.currNode != nil {
		return i.currNode.Value()
	} else {
		return nil
	}
}
