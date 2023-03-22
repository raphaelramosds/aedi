package helper

import "fmt"

// Basic structures

type No struct {
	value int
	next  *No
}

type LinkedStack struct {
	head *No
	tam  int
}

// Internal functions

func (ls *LinkedStack) getAll() {
	aux := ls.head

	fmt.Printf("[")

	for aux != nil {
		fmt.Printf("%d", aux.value)
		if aux.next != nil {
			fmt.Printf(" ")
		}
		aux = aux.next
	}

	fmt.Printf("]\n")
}

func initNode(value int) *No {

	newNode := new(No)
	newNode.value = value
	newNode.next = nil

	return newNode
}

// Public functions

func (ls *LinkedStack) Init() {
	ls.tam = 0
	ls.head = nil
}

func (ls *LinkedStack) Push(value int) {

	newNode := initNode(value)

	if ls.tam == 0 {
		ls.head = newNode
	} else {
		aux := ls.head
		for aux != nil {
			if aux.next == nil {
				aux.next = newNode
				break
			}
			aux = aux.next
		}
	}

	ls.tam++
}

func (ls *LinkedStack) IsEmpty() bool {
	return ls.tam == 0
}

func (ls *LinkedStack) Pop() {

	aux := ls.head

	for i := 0; aux != nil; i++ {
		if i == ls.tam-2 {
			aux.next = nil
			break
		}
		aux = aux.next
	}

	ls.tam--
}

func (ls *LinkedStack) Peek() int {

	aux := ls.head

	for i := 0; aux != nil; i++ {
		if i == ls.tam-1 {
			return aux.value
		}
		aux = aux.next
	}

	return -1
}

func (ls *LinkedStack) Size() int {
	return ls.tam
}
