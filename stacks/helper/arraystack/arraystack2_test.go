package helper

import (
	"testing"
)

func TestInspectTop(t *testing.T) {

	var stack ArrayStack

	stack.Init()

	length := len(stack.values)

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

func TestAddIntoAFilledStack(t *testing.T) {
	var stack ArrayStack

	stack.Init()

	length := 2 * len(stack.values)

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
