// Package linkedlist provides a basic implementation of a singly linked list.
// It supports common operations such as appending, prepending, removing nodes, and reversing the list.
package linkedlist

import (
	"fmt"
)

// Node represents an element in the linked list.
// It contains a value and a pointer to the next node in the list.
type Node struct {
	value int
	next  *Node
}

// LinkedList represents a singly linked list.
// It contains a pointer to the head node, the tail node, and the length of the list.
type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

// PrintLinkedList prints all the values in the linked list in order.
// It starts from the head node and iterates through the list until the end.
func (ll *LinkedList) PrintLinkedList() {
	temp := ll.Head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

// Append adds a new node with the specified value to the end of the linked list.
// If the list is empty, the new node becomes both the head and the tail.
func (ll *LinkedList) Append(value int) {
	newNode := NewNode(value)
	if ll.Length == 0 {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.next = newNode
		ll.Tail = newNode
	}
	ll.Length++
}

// RemoveLast removes the last node in the linked list and returns it.
// If the list is empty, it returns nil. If there's only one node, the list becomes empty.
func (ll *LinkedList) RemoveLast() *Node {
	if ll.Length == 0 {
		return nil
	}

	temp, pre := ll.Head, ll.Head

	for temp != nil {
		pre = temp
		temp = temp.next
	}
	ll.Tail = pre
	ll.Tail.next = nil
	ll.Length--
	if ll.Length == 0 {
		ll.Tail = nil
		ll.Head = nil
	}
	return temp
}

// NewNode creates and returns a new node with the specified value.
func NewNode(value int) *Node {
	return &Node{value: value}
}

// NewLinkedList creates and returns a new linked list initialized with a single node.
func NewLinkedList(value int) *LinkedList {
	node := NewNode(value)
	return &LinkedList{
		Head:   node,
		Tail:   node,
		Length: 1,
	}
}

// Prepend adds a new node with the specified value to the beginning of the linked list.
// If the list is empty, the new node becomes both the head and the tail.
func (ll *LinkedList) Prepend(value int) {
	newNode := NewNode(value)
	if ll.Length == 0 {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		newNode.next = ll.Head
		ll.Head = newNode
	}
	ll.Length++
}

// RemoveFirst removes and returns the first node in the linked list.
// If the list is empty, it returns nil.
func (ll *LinkedList) RemoveFirst() *Node {
	if ll.Length == 0 {
		return nil
	}

	temp := ll.Head
	ll.Head = ll.Head.next
	temp.next = nil
	ll.Length--

	if ll.Length == 0 {
		ll.Tail = nil
	}

	return temp
}

// Get returns the node at the specified index in the linked list.
// If the index is out of bounds, it returns nil.
func (ll *LinkedList) Get(index int) *Node {
	if index < 0 || index >= ll.Length {
		return nil
	}

	temp := ll.Head
	for i := 0; i < index; i++ {
		temp = temp.next
	}
	return temp
}

// Set updates the value of the node at the specified index.
// It returns true if the node is successfully updated, and false otherwise.
func (ll *LinkedList) Set(index int, value int) bool {
	node := ll.Get(index)
	if node != nil {
		node.value = value
		return true
	}
	return false
}

// Insert inserts a new node with the specified value at the given index in the linked list.
// It returns true if the insertion is successful, and false otherwise.
func (ll *LinkedList) Insert(index int, value int) bool {
	if index < 0 || index >= ll.Length {
		return false
	}
	if index == 0 {
		ll.Prepend(value)
		return true
	}
	if index == ll.Length {
		ll.Append(value)
		return true
	}
	newNode := NewNode(value)
	temp := ll.Get(index - 1)
	newNode.next = temp.next
	temp.next = newNode
	ll.Length++
	return true
}

// Remove removes and returns the node at the specified index in the linked list.
// If the index is out of bounds, it returns nil.
func (ll *LinkedList) Remove(index int) *Node {
	if index < 0 || index >= ll.Length {
		return nil
	}

	if index == 0 {
		return ll.RemoveFirst()
	}

	if index == ll.Length-1 {
		return ll.RemoveLast()
	}

	previous := ll.Get(index - 1)
	temp := previous.next

	previous.next = temp.next
	temp.next = nil
	ll.Length--
	return temp
}

// Reverse reverses the linked list in place.
func (ll *LinkedList) Reverse() {
	temp := ll.Head
	ll.Head = ll.Tail
	ll.Tail = temp

	var before *Node = nil
	var after = temp.next

	for i := 0; i < ll.Length; i++ {
		after = temp.next
		temp.next = before
		before = temp
		temp = after
	}
}
