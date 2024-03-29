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
		newPeek := as.values[as.size-1]
		as.size--
		return newPeek, nil
	}
}

func (as *ArrayStack) Peek() (int, error) {
	// TODO peek is the at the beginning
}

func (as *ArrayStack) Empty() bool {
	return as.size == 0
}

func (as *ArrayStack) Size() int {
	return as.size
}
