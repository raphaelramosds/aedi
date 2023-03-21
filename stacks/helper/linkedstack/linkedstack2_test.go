package helper

import (
	"testing"
)

func TestInspectTop(t *testing.T) {

	var stack LinkedStack

	stack.Init()

	length := 10

	for i := 0; i < length; i++ {
		stack.Push(i)
	}

	for i := 0; i < 2; i++ {
		stack.Pop()
		length--
	}

	result := stack.Peek()
	expected := length - 1

	if result != expected {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestInspectIfItsEmpty(t *testing.T) {
	var stack LinkedStack

	stack.Init()

	for i := 0; i < 10; i++ {
		stack.Push(i)
	}

	for i := 0; i < 10; i++ {
		stack.Pop()
	}

	result := stack.IsEmpty()
	expected := true

	if result != expected {
		t.Errorf("Expected %t but got %t", expected, result)
	}
}
