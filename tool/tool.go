package main

import (
	_ "bufio"
	"fmt"
	"os"
	"text/template"
)

type param struct {
	Name string
}

func main() {
	const st = `package aggregate

import (
	"github.com/user/aggregate/iterator"
	"github.com/user/container"
	"github.com/user/container/skiplist"
)

type {{.Name}} struct {
	*skiplist.SkipList
	// your own field here
}

type {{.Name}}Iterator struct {
	currNode  *skiplist.SkipListNode
	aggregate *{{.Name}}
	// your own field here
}

func New{{.Name}}() *{{.Name}} {
	return &{{.Name}}{skiplist.NewSkipList()}
}

func (m *{{.Name}}) Iterator() iterator.Iterator {
	return &{{.Name}}Iterator{currNode: m.First(), aggregate: m}
}

func (s *{{.Name}}) Find(key container.Comparer) iterator.Iterator {
	node := s.SearchNode(key)
	return &{{.Name}}Iterator{currNode: node, aggregate: s}
}

func (s *{{.Name}}) LowerBound(key container.Comparer) iterator.Iterator {
	node := s.LowerBoundNode(key)
	return &{{.Name}}Iterator{currNode: node, aggregate: s}
}

func (i *{{.Name}}Iterator) Next() {
	i.currNode = i.currNode.Next()
}

func (i *{{.Name}}Iterator) Prev() {
	i.currNode = i.currNode.Prev()
}

func (i *{{.Name}}Iterator) First() interface{} {
	if i.currNode != nil {
		return i.currNode.Key()
	} else {
		return nil
	}
}

func (i *{{.Name}}Iterator) Second() interface{} {
	if i.currNode != nil {
		return i.currNode.Value()
	} else {
		return nil
	}
}

func (i *{{.Name}}Iterator) Begin() bool {
	if i.currNode == i.aggregate.First() {
		return true
	} else {
		return false
	}
}

func (i *{{.Name}}Iterator) End() bool {
	if i.currNode == nil {
		return true
	} else {
		return false
	}
}
`
	g := param{
		Name: "Zmap",
	}

	outFile, err := os.Create("text.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	//output := bufio.NewWriter(outFile)
	t := template.Must(template.New("aggregate").Parse(st))
	if err := t.Execute(outFile, g); err != nil {
		fmt.Println(err)
	}
}
