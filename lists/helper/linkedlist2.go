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
	}

	ll.tam++

	fmt.Printf("%d was addded. List: ", value)
	ll.getAll()
}

func (ll *LinkedList) AddOnIndex(value int, index int) {

	newNode := initNode(value)
	aux := ll.head

	for i := 0; aux != nil; i++ {

		// we have to put one element between two elements
		// so, we iterate til the first, indexed at index - 1

		if i == index-1 {

			// take his address

			left := aux

			// and the third one's

			right := aux.next

			// update pointers

			left.next = newNode
			newNode.next = right

			ll.tam++

			break
		}
		aux = aux.next
	}

	fmt.Printf("%d was addded at %d. List: ", value, index)
	ll.getAll()
}

func (ll *LinkedList) Remove() {

	aux := ll.head

	for i := 0; aux != nil; i++ {

		// just make the second to last element
		// pointing to null

		if i == ll.tam-2 {
			aux.next = nil
			ll.tam--
			break
		}

		aux = aux.next
	}

	fmt.Printf("Removed at end. List: ")
	ll.getAll()
}

func (ll *LinkedList) RemoveOnIndex(index int) {

	aux := ll.head

	for i := 0; aux != nil; i++ {
		if i == index-1 {
			left := aux
			middle := left.next
			right := middle.next

			left.next = right
			middle.next = nil

			break
		}
		aux = aux.next
	}

	fmt.Printf("Removed at %d. List: ", index)
	ll.getAll()
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
