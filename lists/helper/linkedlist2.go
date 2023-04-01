package helper

import "errors"

// Basic structures

type LinkedList struct {
	head *Node1P
	size int
}

type Node1P struct {
	value int
	next  *Node1P
}

// Interface functions

func (ll *LinkedList) Add(value int) {

	newNode := Node1P{value, nil}

	if ll.size == 0 {
		ll.head = &newNode
	} else {
		curr := ll.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = &newNode
	}

	ll.size++
}

func (ll *LinkedList) AddOnIndex(value int, index int) error {

	if index >= 0 && index <= ll.size {

		newNode := Node1P{value, nil}

		if index == 0 {

			if ll.head != nil {
				newNode.next = ll.head
			}

			ll.head = &newNode

		} else {
			curr := ll.head
			for i := 0; i < index-1; i++ {
				curr = curr.next
			}
			newNode.next = curr.next
			curr.next = &newNode
		}

		ll.size++
		return nil

	} else {
		if index < 0 {
			return errors.New("can't add into a negative index")
		} else {
			return errors.New("can't add into an index > list size")
		}
	}
}

func (ll *LinkedList) Remove() {}

func (ll *LinkedList) RemoveOnIndex(index int) error {
	if ll.size == 0 {
		return errors.New("list is empty")
	} else if index >= 0 && index <= ll.size {
		curr := ll.head
		if index == 0 {
			ll.head = curr.next
			curr.next = nil
		} else {
			for i := 0; i < index-1; i++ {
				curr = curr.next
			}
			curr.next = curr.next.next
		}

		ll.size--
		return nil
	} else {
		if index < 0 {
			return errors.New("can't remove from a negative index")
		} else {
			return errors.New("can't remove from an index > list size")
		}
	}
}

func (ll *LinkedList) Get(index int) (int, error) {
	if index >= 0 && index < ll.size {
		curr := ll.head
		for i := 0; i < index; i++ {
			curr = curr.next
		}
		return curr.value, nil

	} else {
		if index < 0 {
			return -1, errors.New("can't add into a negative index")
		} else {
			return -1, errors.New("can't add into an index > list size")
		}
	}
}

func (ll *LinkedList) Set(value int, index int) error {
	if index >= 0 && index <= ll.size {
		curr := ll.head
		for i := 0; i < index; i++ {
			curr = curr.next
		}
		curr.value = value
		return nil
	} else {
		if index < 0 {
			return errors.New("can't set a negative index")
		} else {
			return errors.New("can't set an index > list size")
		}
	}
}

func (ll *LinkedList) Size() int {
	return ll.size
}
