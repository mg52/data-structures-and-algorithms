package bst

import (
	"fmt"
)

type BST struct {
	Root *Node
}

func NewBST() *BST {
	return &BST{
		Root: nil,
	}
}

type Node struct {
	Value byte
	Left  *Node
	Right *Node
	Level int
}

func (t *BST) Search(data byte) *Node {
	retVal := t.Root.Search(data)
	return retVal
}

func (n *Node) Search(data byte) *Node {
	if n.Value == data {
		return n
	} else if data > n.Value && n.Right != nil {
		return n.Right.Search(data)
	} else if data < n.Value && n.Left != nil {
		return n.Left.Search(data)
	}
	return nil
}

func (t *BST) Insert(data byte) {
	if t.Root == nil {
		t.Root = &Node{Value: data, Level: 1}
	} else {
		t.Root.Insert(data, 2)
	}
}

func (n *Node) Insert(data byte, currentLevel int) {
	if data <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: data, Level: currentLevel}
		} else {
			n.Left.Insert(data, currentLevel+1)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: data, Level: currentLevel}
		} else {
			n.Right.Insert(data, currentLevel+1)
		}
	}
}

func PrintPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%c%d ", n.Value, n.Level)
		PrintPreOrder(n.Left)
		PrintPreOrder(n.Right)
	}
}

func PrintPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		PrintPostOrder(n.Left)
		PrintPostOrder(n.Right)
		fmt.Printf("%c%d ", n.Value, n.Level)
	}
}

func PrintInOrder(n *Node) {
	if n == nil {
		return
	} else {
		PrintInOrder(n.Left)
		fmt.Printf("%c%d ", n.Value, n.Level)
		PrintInOrder(n.Right)
	}
}
