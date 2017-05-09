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

// NewNode create new node with given value
func NewNode(value interface{}) *Node {
	return &Node{data: value}
}

// NewList new linkedlist
func NewList() *LinkedList {
	list := new(LinkedList)
	list.length = 0
	return list
}

// Size return length of list
func (list *LinkedList) Size() int {
	return list.length
}

// AppendNode addend a node at the end
func (list *LinkedList) AppendNode(node *Node) {
	if list.Size() == 0 {
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

// AppendValue addend a node at the end with given value
func (list *LinkedList) AppendValue(value interface{}) {
	node := NewNode(value)
	list.AppendNode(node)
}

// PrintList print the list
func (list *LinkedList) PrintList() {
	if list.Size() == 0 {
		fmt.Println(errors.New("List is empty"))
	} else {
		head := list.head
		for head != nil {
			fmt.Println(head.data)
			head = head.next
		}
	}
}

// Prepend add a value at the beginning
func (l *LinkedList) Prepend(value interface{}) {

}

// Add add at given index
func (l *LinkedList) Add(value interface{}, index int) error {
	return nil
}

// RemoveValue remove value
func (l *LinkedList) RemoveValue(value interface{}) error {
	return nil
}

// RemoveNode remove node at index
func (l *LinkedList) RemoveNode(index int) error {
	return nil
}

// Get get node at index
func (l *LinkedList) Get(index int) (*Node, error) {
	return nil, nil
}

// Find find if node exists
func (l *LinkedList) Find(node *Node) (int, error) {
	return 0, nil
}

// Clear clear list
func (l *LinkedList) Clear() {

}

// Concat add 2 lists
func (l *LinkedList) Concat(k *LinkedList) {

}

// Map map nodes
func (list *LinkedList) Map(f func(node *Node)) {

}

// Each iterate through each node
func (list *LinkedList) Each(f func(node Node)) {

}
