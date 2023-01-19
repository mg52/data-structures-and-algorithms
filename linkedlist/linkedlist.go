package linkedlist

import (
	"fmt"
)

type Node struct {
	info interface{}
	next *Node
}

type List struct {
	Head *Node
}

func (l *List) Insert(d interface{}) {
	aNode := &Node{info: d, next: nil}
	if l.Head == nil {
		l.Head = aNode
	} else {
		p := l.Head
		for p.next != nil {
			p = p.next
		}
		p.next = aNode
	}
}

func Print(n *Node) {
	if n != nil {
		fmt.Printf("%v", n.info)
	}
	if n.next != nil {
		fmt.Print(" -> ")
		Print(n.next)
	} else {
		fmt.Println("")
	}
}

func (l *List) PrintList() {
	Print(l.Head)
}

func ToSlice(n *Node, arr []interface{}) []interface{} {
	if n != nil {
		arr = append(arr, n.info)
	}
	if n.next != nil {
		return ToSlice(n.next, arr)
	} else {
		return arr
	}
}

func (l *List) ToSlice() []interface{} {
	var retVal []interface{}
	retVal = ToSlice(l.Head, retVal)
	return retVal
}
