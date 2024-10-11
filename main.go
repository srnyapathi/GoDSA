package main

import (
	"fmt"

	"github.com/srnyapathi/GoDSA/linkedlist"
)

func main() {
	// Initialize the linked list with one element
	ll := linkedlist.NewLinkedList(1)
	fmt.Println("Initial linked list:")
	ll.PrintLinkedList()

	// Append new elements
	ll.Append(2)
	fmt.Println("Linked list after appending 2:")
	ll.PrintLinkedList()

	ll.Append(3)
	fmt.Println("Linked list after appending 3:")
	ll.PrintLinkedList()

	// Prepend an element
	ll.Prepend(9)
	fmt.Println("Linked list after prepending 9:")
	ll.PrintLinkedList()

	// Remove the first element
	removedFirst := ll.RemoveFirst()
	fmt.Printf("Removed first element: %d\n", removedFirst.Value)
	ll.PrintLinkedList()

	// Remove the last element
	removedLast := ll.RemoveLast()
	fmt.Printf("Removed last element: %d\n", removedLast.Value)
	ll.PrintLinkedList()

	// Get element at index 1
	node := ll.Get(1)
	if node != nil {
		fmt.Printf("Element at index 1: %d\n", node.Value)
	}

	// Set value of the element at index 1
	ll.Set(1, 99)
	fmt.Println("Linked list after setting value at index 1 to 99:")
	ll.PrintLinkedList()

	// Insert an element at index 1
	ll.Insert(1, 50)
	fmt.Println("Linked list after inserting 50 at index 1:")
	ll.PrintLinkedList()

	// Remove element at index 2
	removed := ll.Remove(2)
	if removed != nil {
		fmt.Printf("Removed element at index 2: %d\n", removed.Value)
	}
	ll.PrintLinkedList()

	// Reverse the linked list
	ll.Reverse()
	fmt.Println("Reversed linked list:")
	ll.PrintLinkedList()
}
