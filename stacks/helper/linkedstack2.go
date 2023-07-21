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

// TODO insert af beginning since it's a linked stack
func (ls *LinkedStack) Push(value int) {
}

// TODO pop at beginning since it's a linked stack
func (ls *LinkedStack) Pop() (int, error) {
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
