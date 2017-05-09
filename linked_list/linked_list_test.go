package linkedList

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := NewList()
	list.AppendValue(1)
	list.AppendValue("nomnom")
	if list.Size() < 2 {
		t.Error()
	}
	list.PrintList()
}
