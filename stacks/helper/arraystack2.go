package helper

import "errors"

// Basic Structure

type ArrayStack struct {
	values []int
	size   int
}

// Internal functions

func (as *ArrayStack) double() {
	doubledArray := make([]int, 2*len(as.values))
	for i := 0; i < as.size; i++ {
		doubledArray[i] = as.values[i]
	}
	as.values = doubledArray
}

func (as *ArrayStack) Init(size int) error {
	if size > 0 {
		as.size = 0
		as.values = make([]int, size)
		return nil
	} else {
		return errors.New("can't init an array with size <= 0")
	}
}

// Interface functions

func (as *ArrayStack) Push(value int) {
	if as.size == len(as.values) {
		as.double()
	}
	as.values[as.size] = value
	as.size++
}

func (as *ArrayStack) Pop() (int, error) {
	if as.size == 0 {
		return -1, errors.New("stack is empty")
	} else {
		removedElement := as.values[as.size-1]
		as.size--
		return removedElement, nil
	}
}

func (as *ArrayStack) Peek() (int, error) {
	if as.size == 0 {
		return -1, errors.New("stack is empty")
	} else {
		return as.values[as.size-1], nil
	}
}

func (as *ArrayStack) Empty() bool {
	return as.size == 0
}

func (as *ArrayStack) Size() int {
	return as.size
}
