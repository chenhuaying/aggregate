package aggregate

import (
	"github.com/user/aggregate/iterator"
	"github.com/user/skiplist"
)

type zset struct {
	*skiplist.SkipList
}

type SkipListIterator struct {
	currNode  *skiplist.SkipListNode
	aggregate *zset
}

func NewZset() *zset {
	return &zset{skiplist.NewSkipList()}
}

func (s *zset) Iterator() iterator.Iterator {
	return &SkipListIterator{currNode: s.First(), aggregate: s}
}

func (s *zset) Find(key uint32) iterator.Iterator {
	node := s.SearchNode(key)
	return &SkipListIterator{currNode: node, aggregate: s}
}

func (s *zset) LowerBound(key uint32) iterator.Iterator {
	node := s.LowerBoundNode(key)
	return &SkipListIterator{currNode: node, aggregate: s}
}

func (i *SkipListIterator) Next() {
	i.currNode = i.currNode.Next()
}

func (i *SkipListIterator) Prev() {
	i.currNode = i.currNode.Prev()
}

func (i *SkipListIterator) End() bool {
	if i.currNode == nil {
		return true
	} else {
		return false
	}
}

func (i *SkipListIterator) First() interface{} {
	if i.currNode != nil {
		return i.currNode.Key()
	} else {
		return nil
	}
}

func (i *SkipListIterator) Second() interface{} {
	if i.currNode != nil {
		return i.currNode.Value()
	} else {
		return nil
	}
}
