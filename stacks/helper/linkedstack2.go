package helper

import "errors"

// Basic structures

type Node struct {
	value int
	next  *Node
}

type LinkedStack struct {
	head *Node
	size int
}

// Public functions

func (ls *LinkedStack) Push(value int) {
	newNode := Node{value, nil}
	curr := ls.head
	if curr != nil {
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = &newNode
	} else {
		ls.head = &newNode
	}
	ls.size++
}

func (ls *LinkedStack) Pop() (int, error) {
	if ls.size == 0 {
		return -1, errors.New("stack is empty")
	} else {
		curr := ls.head
		for i := 0; i < ls.size-1; i++ {
			curr = curr.next
		}
		curr.next = nil
		ls.size--
		return curr.value, nil
	}
}

func (ls *LinkedStack) Peek() (int, error) {
	// TODO peek is the at the beginning
}

func (ls *LinkedStack) Empty() bool {
	return ls.size == 0
}

func (ls *LinkedStack) Size() int {
	return ls.size
}
