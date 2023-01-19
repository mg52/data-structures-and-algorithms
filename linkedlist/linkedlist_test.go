package linkedlist

import (
	"testing"
)

// TestList_ToSlice tests linkedlist.ToSlice
func TestList_ToSlice(t *testing.T) {
	sl := List{}
	sl.Insert(1)
	sl.Insert(2)
	sl.Insert(3)
	sl.Insert(4)
	sl.Insert(5)

	result := sl.ToSlice()

	var expected []interface{}
	expected = append(expected, 1)
	expected = append(expected, 2)
	expected = append(expected, 3)
	expected = append(expected, 4)
	expected = append(expected, 5)

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("result %v, expected %v", result[i], expected[i])
		}
	}
}
