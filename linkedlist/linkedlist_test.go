package linkedlist

import (
	"testing"
)

func TestAppend(t *testing.T) {
	ll := NewLinkedList(10)
	ll.Append(20)
	ll.Append(30)

	if ll.Length != 3 {
		t.Errorf("Expected length 3, got %d", ll.Length)
	}

	if ll.Tail.value != 30 {
		t.Errorf("Expected tail value 30, got %d", ll.Tail.value)
	}

	if ll.Head.value != 10 {
		t.Errorf("Expected head value 10, got %d", ll.Head.value)
	}
}

func TestPrepend(t *testing.T) {
	ll := NewLinkedList(10)
	ll.Prepend(5)
	ll.Prepend(1)

	if ll.Length != 3 {
		t.Errorf("Expected length 3, got %d", ll.Length)
	}

	if ll.Head.value != 1 {
		t.Errorf("Expected head value 1, got %d", ll.Head.value)
	}

	if ll.Tail.value != 10 {
		t.Errorf("Expected tail value 10, got %d", ll.Tail.value)
	}
}

func TestRemoveFirst(t *testing.T) {
	ll := NewLinkedList(10)
	ll.Append(20)
	ll.Append(30)
	removed := ll.RemoveFirst()

	if ll.Length != 2 {
		t.Errorf("Expected length 2, got %d", ll.Length)
	}

	if removed.value != 10 {
		t.Errorf("Expected removed value 10, got %d", removed.value)
	}

	if ll.Head.value != 20 {
		t.Errorf("Expected head value 20, got %d", ll.Head.value)
	}
}

func TestRemoveLast(t *testing.T) {
	ll := NewLinkedList(10)
	ll.Append(20)
	ll.Append(30)
	removed := ll.RemoveLast()

	if ll.Length != 2 {
		t.Errorf("Expected length 2, got %d", ll.Length)
	}

	if removed.value != 30 {
		t.Errorf("Expected removed value 30, got %d", removed.value)
	}

	if ll.Tail.value != 20 {
		t.Errorf("Expected tail value 20, got %d", ll.Tail.value)
	}
}

func TestReverse(t *testing.T) {
	ll := NewLinkedList(10)
	ll.Append(20)
	ll.Append(30)
	ll.Reverse()

	if ll.Head.value != 30 {
		t.Errorf("Expected head value 30, got %d", ll.Head.value)
	}

	if ll.Tail.value != 10 {
		t.Errorf("Expected tail value 10, got %d", ll.Tail.value)
	}
}

func TestInsert(t *testing.T) {
	ll := NewLinkedList(10)
	ll.Append(30)
	ll.Insert(1, 20)

	if ll.Length != 3 {
		t.Errorf("Expected length 3, got %d", ll.Length)
	}

	if ll.Get(1).value != 20 {
		t.Errorf("Expected value 20 at index 1, got %d", ll.Get(1).value)
	}
}

func TestRemove(t *testing.T) {
	ll := NewLinkedList(10)
	ll.Append(20)
	ll.Append(30)
	removed := ll.Remove(1)

	if ll.Length != 2 {
		t.Errorf("Expected length 2, got %d", ll.Length)
	}

	if removed.value != 20 {
		t.Errorf("Expected removed value 20, got %d", removed.value)
	}

	if ll.Get(1).value != 30 {
		t.Errorf("Expected value 30 at index 1, got %d", ll.Get(1).value)
	}
}
