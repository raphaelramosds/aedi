package helper

import "fmt"

// Basic structure

type ArrayQueue struct {
	front  int
	rear   int
	values []int
	tam    int
}

func (aq *ArrayQueue) Display() {

	fmt.Printf("[")

	if aq.rear < aq.front {
		for i := aq.front; i < len(aq.values); i++ {
			fmt.Printf("%d ", aq.values[i])
		}
		for i := 0; i <= aq.rear; i++ {
			fmt.Printf("%d ", aq.values[i])
		}
	} else {
		for i := aq.front; i <= aq.rear; i++ {
			fmt.Printf("%d ", aq.values[i])
		}
	}

	fmt.Printf("]\n")

}

func (aq *ArrayQueue) Init() {
	aq.values = make([]int, 10) // init with 10 elements
	aq.tam = 0
	aq.rear = -1
	aq.front = -1
}

func (aq *ArrayQueue) Push(value int) {
	if aq.tam == 0 {
		aq.rear = 0
		aq.front = 0
	} else if aq.rear == len(aq.values)-1 {
		aq.rear = 0
	} else {
		aq.rear++
	}

	aq.values[aq.rear] = value
	aq.tam++
}

func (aq *ArrayQueue) Pop() {

	if aq.front == aq.rear {
		aq.rear = -1
		aq.front = -1
	} else if aq.front == len(aq.values)-1 {
		aq.front = 0
	} else {
		aq.front++
	}

	aq.tam--

}

func (aq *ArrayQueue) Peek() int {
	return aq.values[aq.front]
}

func (aq *ArrayQueue) Size() int {
	return aq.tam
}

func (aq *ArrayQueue) IsEmpty() bool {
	return aq.front == -1 && aq.rear == -1
}

func (aq *ArrayQueue) IsFullFilled() bool {
	return (aq.rear+1)%10 == aq.front
}
