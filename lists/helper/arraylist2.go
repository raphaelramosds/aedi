package helper

import "errors"

// Basic structures

type ArrayList struct {
	values []int
	size   int
}

// Internal functions

func (al *ArrayList) double() {
	doubledValues := make([]int, len(al.values)*2)
	for i := 0; i < len(al.values); i++ {
		doubledValues[i] = al.values[i]
	}
	al.values = doubledValues
}

// Public functions

func (al *ArrayList) Init(size int) error {
	if size > 0 {
		al.values = make([]int, size)
		return nil
	} else {
		return errors.New("can't init arraylist with size <= 0")
	}
}

func (al *ArrayList) Add(value int) {

	if al.size >= len(al.values) {
		al.double()
	}

	al.values[al.size] = value
	al.size++
}

func (al *ArrayList) AddOnIndex(value int, index int) error {

	if al.size >= len(al.values) {
		al.double()
	}

	// Shift all elements to the right starting with the last element

	for i := al.size; i > index; i-- {
		al.values[i] = al.values[i-1]
	}

	al.values[index] = value
	al.size++

	return nil
}

func (l *ArrayList) Remove() {
	if l.size > 0 {
		l.size--
	}
}

func (al *ArrayList) RemoveOnIndex(index int) error {

	// Put right side values on the place of the excluded element

	for i := index; i < al.size-1; i++ {
		al.values[i] = al.values[i+1]
	}

	al.size--

	return nil
}

func (al *ArrayList) Get(index int) (int, error) {
	return al.values[index], nil
}

func (al *ArrayList) Set(value, index int) error {
	al.values[index] = value
	return nil
}

func (al *ArrayList) Size() int {
	return al.size
}
