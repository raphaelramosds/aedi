package helper

import (
	"errors"
	"fmt"
)

type CyclicBuffer struct {
	values []int
	rear   int
	front  int
	size   int
}

func (aq *CyclicBuffer) Init(size int) {
	aq.values = make([]int, size)
	aq.front = -1
	aq.rear = -1
	aq.size = 0
}

func (aq *CyclicBuffer) Display() {
	fmt.Printf("%d %d : ", aq.front, aq.rear)
	index := aq.front
	fmt.Printf("[")
	for j := 0; j < aq.size; j++ {
		fmt.Printf("%d ", aq.values[index%len(aq.values)])
		index++
	}
	fmt.Printf("]\n")
}

func (aq *CyclicBuffer) Double() {
	doubledQueue := make([]int, 2*aq.size)
	curr := aq.front
	for j := 0; j < aq.size; j++ {
		doubledQueue[j] = aq.values[curr%len(aq.values)]
		curr++
	}
	aq.front = 0
	aq.rear = aq.size - 1
	aq.values = doubledQueue
}

func (aq *CyclicBuffer) Enqueue(value int) {
	if (aq.rear+1)%len(aq.values) == aq.front {
		aq.Double()
	}
	if aq.front == -1 && aq.rear == -1 {
		aq.rear++
		aq.front++
	} else {
		aq.rear = (aq.rear + 1) % len(aq.values)
	}
	aq.values[aq.rear] = value
	aq.size++
}

func (aq *CyclicBuffer) Dequeue() error {
	if aq.front != -1 && aq.rear != -1 {
		aq.front = (aq.front + 1) % len(aq.values)
		aq.size--
		return nil
	} else {
		return errors.New("Queue is empty")
	}
}

func (aq *CyclicBuffer) Peek() int {
	if aq.size != 0 {
		return aq.values[aq.front]
	}
	return -1
}
