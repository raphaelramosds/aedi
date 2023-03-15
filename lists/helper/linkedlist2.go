package helper

import "fmt"

type LinkedList struct {
	head *No
	tam  int
}

type No struct {
	value int
	next  *No
}

// Internal functions

func initNode(value int) *No {

	newNode := new(No)
	newNode.next = nil
	newNode.value = value

	return newNode
}

func (ll *LinkedList) getAll() {

	aux := ll.head

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

// Public functions

func (ll *LinkedList) Add(value int) {

	newNode := initNode(value)

	if ll.head == nil {
		ll.head = newNode
	} else {
		temp := ll.head
		ll.head = newNode
		newNode.next = temp
		ll.tam++
	}

	ll.getAll()
}

func (ll *LinkedList) AddOnIndex(value int, index int) {

	newNode := initNode(value)
	aux := ll.head

	// iterate until the index - 1 element

	for i := 0; aux != nil; i++ {
		if i == index-1 {

			rear := aux
			front := aux.next

			rear.next = newNode
			newNode.next = front

			ll.tam++
			ll.getAll()

			return
		}
		aux = aux.next
	}

}

func (ll *LinkedList) Get(value int) int {
	aux := ll.head

	for i := 0; aux != nil; i++ {
		if i == value {
			return aux.value
		}
		aux = aux.next
	}

	return -1
}

func (ll *LinkedList) Set(value int, index int) int {
	aux := ll.head
	for i := 0; aux != nil; i++ {
		if i == index {
			aux.value = value
		}
		aux = aux.next
	}

	ll.getAll()

	return -1
}

func (ll *LinkedList) Size() int {
	return ll.tam
}
