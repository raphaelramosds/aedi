package helper

import "fmt"

// Basic structures

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

	// O(1) case

	if ll.tam == 0 {

		ll.head = newNode

	} else {

		// O(n) case

		aux := ll.head

		for aux != nil {
			if aux.next == nil {
				aux.next = newNode
				break
			}
			aux = aux.next
		}
	}

	ll.tam++

	fmt.Printf("%d added at end (Size %d): ", value, ll.Size())
	ll.getAll()

}

func (ll *LinkedList) AddOnIndex(value int, index int) {

	if index >= ll.tam || index < 0 {
		return
	}

	newNode := initNode(value)
	aux := ll.head

	// O(1) case

	if index == 0 {
		ll.head = newNode
		newNode.next = aux
	} else {

		// O(n) case

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

				break
			}
			aux = aux.next
		}
	}

	ll.tam++

	fmt.Printf("%d was addded at %d (Size %d): ", value, index, ll.Size())
	ll.getAll()
}

func (ll *LinkedList) Remove() {

	if ll.tam == 0 {
		return
	}

	// go to the second last element
	// and make it pointing to null

	aux := ll.head

	for i := 0; aux != nil; i++ {

		if i == ll.tam-2 {
			aux.next = nil
		}

		aux = aux.next
	}

	ll.tam--

	fmt.Printf("Removed at end (Size %d): ", ll.Size())
	ll.getAll()

}

func (ll *LinkedList) RemoveOnIndex(index int) {

	if index >= ll.tam || index < 0 {
		return
	}

	aux := ll.head

	// O(1) case

	if index == 0 {
		ll.head = aux.next
		aux.next = nil
	} else {

		// O(n) case

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
	}

	ll.tam--

	fmt.Printf("Removed at %d (Size %d): ", index, ll.Size())
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

	if index >= ll.tam || index < 0 {
		return -1
	}

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
