package helper

import "fmt"

// Basic Structure

type ArrayStack struct {
	values []int
	tam    int
}

// Internal functions

func (as *ArrayStack) double() {
	doubledArray := make([]int, 2*len(as.values))
	copy(as.values, doubledArray)
	as.values = doubledArray
}

// Public functions

func (as *ArrayStack) Display() {
	fmt.Printf("[")
	for i := 0; i < as.tam; i++ {
		fmt.Printf("%d", as.values[i])
		if i != as.tam-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("]")
}

func (as *ArrayStack) Init() {
	as.tam = 0
	as.values = make([]int, 10)
}

func (as *ArrayStack) Push(value int) {
	if as.tam >= len(as.values) {
		as.double()
	}
	as.values[as.tam] = value
	as.tam++
}

func (as *ArrayStack) IsEmpty() bool {
	return as.tam == 0
}

func (as *ArrayStack) Pop() {
	as.tam--
}

func (as *ArrayStack) Peek() int {
	return as.values[as.tam-1]
}

func (as *ArrayStack) Size() int {
	return as.tam
}
