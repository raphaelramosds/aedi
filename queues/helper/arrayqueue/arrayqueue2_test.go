package helper

import (
	"fmt"
	"testing"
)

func TestInspectPeek(t *testing.T) {
	var queue ArrayQueue

	queue.Init()

	for i := 0; i < 10; i++ {
		queue.Push(i)
	}

	// remove 9 elements

	for i := 0; i < 9; i++ {
		queue.Pop()
	}

	queue.Push(12)
	queue.Pop()

	result := queue.Peek()
	expected := 12

	if result != expected {
		fmt.Printf("Expected %d but got %d", expected, result)
	}
}
