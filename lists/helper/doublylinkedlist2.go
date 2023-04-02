package helper

import "errors"

// Base structures

type Node2P struct {
	value int
	next  *Node2P
	prev  *Node2P
}

type DoublyLinkedList struct {
	head *Node2P
	tail *Node2P
	size int
}

// Interface functions

func (dll *DoublyLinkedList) Add(value int) {
	newNode := Node2P{value, nil, nil}
	if dll.size == 0 {
		dll.head = &newNode
		dll.tail = &newNode
	} else {
		curr := dll.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = &newNode
		newNode.prev = curr
	}
	dll.size++
}

func (dll *DoublyLinkedList) AddOnIndex(value int, index int) error {
	if index >= 0 && index <= dll.size {
		newNode := Node2P{value, nil, nil}
		if index == 0 {
			if dll.head != nil {
				newNode.next = dll.head
			}
			dll.head = &newNode
			dll.tail = &newNode
		} else {
			curr := dll.head
			for i := 0; i < index-1; i++ {
				curr = curr.next
			}
			newNode.next = curr.next // front
			curr.next = &newNode     // prev
			newNode.prev = curr

			// avoids prev of nil in case of add at end

			if curr.next != nil {
				curr.next.prev = &newNode
			}
		}

		dll.size++
		return nil

	} else {
		if index < 0 {
			return errors.New("can't add into a negative index")
		} else {
			return errors.New("can't add into an index > list size")
		}
	}
}

func (dll *DoublyLinkedList) Remove() {}

func (dll *DoublyLinkedList) RemoveOnIndex(index int) error {
	if dll.size == 0 {
		return errors.New("empty list")
	} else if index >= 0 && index < dll.size {
		curr := dll.head
		if index == 0 {
			dll.head = curr.next
			dll.tail = curr.next

			// in case of list has more than one element

			if curr.next != nil {
				curr.next.prev = nil
			}
		} else {
			for i := 0; i < index-1; i++ {
				curr = curr.next
			}
			curr.next = curr.next.next

			// avoids prev of nil in case of removing last element

			if curr.next != nil {
				curr.next.next.prev = curr
			}
		}
		dll.size--
		return nil
	} else {
		if index < 0 {
			return errors.New("can't remove from a negative index")
		} else {
			return errors.New("can't remove from an index > list size")
		}
	}
}

func (dll *DoublyLinkedList) Get(index int) (int, error) {
	if index >= 0 && index < dll.size {
		curr := dll.head
		for i := 0; i < index; i++ {
			curr = curr.next
		}

		return curr.value, nil

	} else {
		if index < 0 {
			return -1, errors.New("can't get a negative index")
		} else {
			return -1, errors.New("can't get an index > list size")
		}
	}
}

func (dll *DoublyLinkedList) Set(value int, index int) error {
	if index >= 0 && index < dll.size {
		curr := dll.head
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

func (dll *DoublyLinkedList) Size() int {
	return dll.size
}
