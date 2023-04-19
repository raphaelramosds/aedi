package helper

import (
	"errors"
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

type LinkedQueue struct {
	head *Node
	rear *Node
	tam  int
}

func (lq *LinkedQueue) Display() {
	curr := lq.head
	fmt.Printf("Peek: %d, Size: %d [", lq.Peek(), lq.Size())
	for curr != nil {
		fmt.Printf("%d ", curr.val)
		curr = curr.next
	}
	fmt.Printf("]\n")
}

func (lq *LinkedQueue) Enqueue(value int) {
	newNode := Node{value, nil}
	if lq.tam == 0 {
		lq.head = &newNode
	} else {
		curr := lq.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = &newNode
	}
	lq.rear = &newNode
	lq.tam++
}

func (lq *LinkedQueue) Dequeue() error {
	if lq.tam == 0 {
		return errors.New("Empty queue")
	}
	temp := lq.head
	lq.head = temp.next
	lq.tam--
	return nil
}

func (lq *LinkedQueue) Size() int {
	return lq.tam
}

func (lq *LinkedQueue) Peek() int {
	if lq.tam == 0 {
		return -1
	}
	return lq.head.val
}

func (lq *LinkedQueue) IsEmpty() bool {
	return lq.tam == 0
}
