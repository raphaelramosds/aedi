package helper

import "fmt"

// Base structures

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

// Internal functions

func initNode(val int) *Node {
	newNode := new(Node)
	newNode.val = val
	return newNode
}

func (dll *DoublyLinkedList) getAll() {
	fmt.Printf("[")
	for i := 0; i < dll.tam; i++ {
		fmt.Printf("%d ", dll.Get(i))
	}
	fmt.Printf("]\n")
}

// Public functions

func (dll *DoublyLinkedList) Init() {
	dll.head = nil
	dll.tam = 0
	dll.tail = nil
}

func (dll *DoublyLinkedList) Add(value int) {

	newNode := initNode(value)

	if dll.tam == 0 {
		dll.head = newNode
		dll.tail = newNode
		newNode.prev = nil
	} else {
		aux := dll.tail
		aux.next = newNode
		dll.tail = newNode
		newNode.prev = aux
	}

	dll.tam++

}

func (dll *DoublyLinkedList) AddOnIndex(value int, index int) {

	if index < 0 || index >= dll.tam {
		return
	}

	newNode := initNode(value)
	aux := dll.head

	if index == 0 {

		dll.head = newNode
		newNode.prev = nil
		newNode.next = aux
		aux.prev = newNode

	} else {

		for i := 0; aux != nil; i++ {
			if i == index-1 {

				right := aux.next

				aux.next = newNode
				newNode.next = right
				right.prev = newNode
				newNode.prev = aux

				break
			}
			aux = aux.next
		}
	}

	dll.tam++
}

func (dll *DoublyLinkedList) Remove() {

	if dll.tam == 0 {
		return
	}

	last := dll.tail
	secondLast := last.prev

	dll.tail = secondLast
	last.prev = nil
	secondLast.next = nil

	dll.tam--

}

func (dll *DoublyLinkedList) RemoveOnIndex(index int) {

	if index < 0 || index >= dll.tam {
		return
	}

	aux := dll.head

	if index == dll.tam-1 {

		last := dll.tail
		secondLast := last.prev

		dll.tail = secondLast
		last.prev = nil
		secondLast.next = nil
	}

	if index == 0 {

		dll.head = aux.next
		aux.next.prev = nil

	} else {
		for i := 0; aux != nil; i++ {
			if i == index {

				left := aux.prev
				right := aux.next

				left.next = right
				right.prev = left

				break
			}
			aux = aux.next
		}
	}

	dll.tam--
}

func (dll *DoublyLinkedList) Get(value int) int {

	if value < 0 || value >= dll.tam {
		return -1
	}

	aux := dll.head

	for i := 0; aux != nil; i++ {
		if i == value {
			return aux.val
		}
		aux = aux.next
	}

	return 0
}

func (dll *DoublyLinkedList) Set(value int, index int) {
	if index < 0 || index >= dll.tam {
		return
	}

	aux := dll.head

	for i := 0; aux != nil; i++ {
		if i == index {
			aux.val = value
			break
		}
		aux = aux.next
	}
}

func (dll *DoublyLinkedList) Size() int {
	return dll.tam
}
