package helper

// Basic structures

type Node struct {
	value int
	next  *Node
}

type LinkedStack struct {
	head *Node
	tam  int
}

// Public functions

func (ls *LinkedStack) Push(value int) {

}

func (ls *LinkedStack) Pop() (int, error) {
	return -1, nil
}

func (ls *LinkedStack) Peek() (int, error) {
	return -1, nil
}

func (ls *LinkedStack) Empty() bool {
	return ls.tam == 0
}

func (ls *LinkedStack) Size() int {
	return ls.tam
}
