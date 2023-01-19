package bst

import (
	"testing"
)

// TestBST_Search tests bst.Search
func TestBST_Search(t *testing.T) {
	tree := NewBST()

	tree.Insert('F')
	tree.Insert('B')
	tree.Insert('A')
	tree.Insert('D')
	tree.Insert('C')
	tree.Insert('E')
	tree.Insert('G')
	tree.Insert('I')
	tree.Insert('H')

	result := tree.Search('D')
	expected := &Node{
		Value: 'D',
		Left:  nil,
		Right: nil,
		Level: 3,
	}

	if result.Level != expected.Level {
		t.Errorf("result %v, expected %v", result.Level, expected.Level)
	}
	if result.Value != expected.Value {
		t.Errorf("result %v, expected %v", result.Value, expected.Value)
	}
}
