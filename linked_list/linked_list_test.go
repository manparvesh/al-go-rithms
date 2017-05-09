package linkedList

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := newList()
	list.appendValue(interface{1})
	list.appendValue(interface{2})
	list.printList()
}
