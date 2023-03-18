package helper

import "fmt"

// Basic structures

type ArrayList struct {
	values []int
	tam    int
}

// Internal functions

func (al *ArrayList) double() {
	doubledValues := make([]int, len(al.values)*2)
	copy(al.values, doubledValues)
	al.values = doubledValues
}

func (al *ArrayList) getAll() {
	fmt.Printf("[")
	for _, element := range al.values {
		fmt.Printf("%d", element)
	}
	fmt.Printf("]\n")
}

// Public functions

func (al *ArrayList) Init() {
	al.tam = 0
	al.values = make([]int, 10)
}

func (al *ArrayList) Add(value int) {

	if al.tam >= len(al.values) {
		al.double()
	}

	al.values[al.tam] = value
	al.tam++
}

func (al *ArrayList) AddOnIndex(value int, index int) {

	if al.tam >= len(al.values) {
		al.double()
	}

	// Shift all elements to the right
	// starting with the last element

	for i := al.tam; i > index; i-- {
		al.values[i] = al.values[i-1]
	}

	al.values[index] = value
	al.tam++
}

func (l *ArrayList) Remove() {
	if l.tam > 0 {
		l.tam--
	}
}

func (al *ArrayList) RemoveOnIndex(index int) {

	// Put right side values on the place
	// of the excluded element

	for i := index; i < al.tam-1; i++ {
		al.values[i] = al.values[i+1]
	}

	al.tam--
}

func (al *ArrayList) Get(index int) int {
	return al.values[index]
}

func (al *ArrayList) Set(value, index int) {
	al.values[index] = value
}

func (al *ArrayList) Size() int {
	return al.tam
}
