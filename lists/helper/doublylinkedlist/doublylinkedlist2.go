package helper

type Node struct {
	next *Node
	val  int
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	tam  int
}
