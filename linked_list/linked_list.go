package linkedList

import (
	"errors"
	"fmt"
)

// LinkedList the linked list struct
type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

// Node the Node struct
type Node struct {
	data interface{}
	prev *Node
	next *Node
}

func newNode(value interface{}) *Node {
	return &Node{data: value}
}

func newList() *LinkedList {
	list := new(LinkedList)
	list.length = 0
	return list
}

func (list *LinkedList) size() int {
	return list.length
}

func (list *LinkedList) appendNode(node *Node) {
	if list.size() == 0 {
		list.head = node
		list.tail = node
	} else {
		exTail := list.tail
		exTail.next = node
		node.prev = exTail
		list.tail = node
	}

	list.length++
}

func (list *LinkedList) appendValue(value *interface{}) {
	node := newNode(value)
	list.appendNode(node)
}

func (list *LinkedList) printList() {
	if list.size() == 0 {
		fmt.Println(errors.New("List is empty"))
	} else {
		head := list.head
		for head != nil {
			fmt.Println(head.data)
		}
	}
}
