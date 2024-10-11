# LinkedList Package 

This is a simple implementation of a singly linked list in Go. The package provides basic operations like appending, prepending, removing nodes, and reversing the list.

## Data Structures

### `Node`
Represents an individual element in the linked list.

```go
type Node struct {
    value int
    next  *Node
}
```

### `LinkedList`
The main structure that contains the linked list and supports various list operations.

```go
type LinkedList struct {
    Head   *Node
    Tail   *Node
    Length int
}
```

## Key Methods

### `Append(value int)`
Adds a new node with the given value to the end of the linked list.

**Pseudo-code:**
1. Create a new node with the given value.
2. If the list is empty, set `Head` and `Tail` to the new node.
3. Otherwise, link the current `Tail` to the new node and update `Tail`.
4. Increment the list length.

Example:
```go
ll.Append(10)
```

### `Prepend(value int)`
Adds a new node with the given value to the beginning of the linked list.

**Pseudo-code:**
1. Create a new node with the given value.
2. If the list is empty, set `Head` and `Tail` to the new node.
3. Otherwise, link the new node to the current `Head` and update `Head`.
4. Increment the list length.

Example:
```go
ll.Prepend(5)
```

### `RemoveFirst() *Node`
Removes and returns the first node in the list.

**Pseudo-code:**
1. If the list is empty, return `nil`.
2. Store the current `Head` in a temporary variable.
3. Set `Head` to the next node.
4. If the list becomes empty, set `Tail` to `nil`.
5. Return the removed node.

Example:
```go
removedNode := ll.RemoveFirst()
```

### `RemoveLast() *Node`
Removes and returns the last node in the list.

**Pseudo-code:**
1. If the list is empty, return `nil`.
2. Traverse the list to find the node before the current `Tail`.
3. Set `Tail` to this node and update its `next` pointer to `nil`.
4. If the list becomes empty, set `Head` to `nil`.
5. Return the removed node.

Example:
```go
removedNode := ll.RemoveLast()
```

### `Reverse()`
Reverses the entire linked list.

**Pseudo-code:**
1. Swap `Head` and `Tail`.
2. Traverse the list and reverse the pointers of each node.
3. Continue until the entire list is reversed.

Example:
```go
ll.Reverse()
```

### `PrintLinkedList()`
Prints all values in the linked list.

**Pseudo-code:**
1. Start from `Head`.
2. Traverse the list, printing the value of each node.
3. Stop when reaching the end of the list.

Example:
```go
ll.PrintLinkedList()
```

## Example Usage

```go
ll := linkedlist.NewLinkedList(10)
ll.Append(20)
ll.Prepend(5)
ll.PrintLinkedList()  // Output: 5 10 20
```
