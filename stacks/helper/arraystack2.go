package helper

// Basic Structure

type ArrayStack struct {
	values []int
	tam    int
}

// Internal functions

func (as *ArrayStack) double() {

}

func (as *ArrayStack) Init(size int) {

}

// Interface functions

func (as *ArrayStack) Push(value int) {}

func (as *ArrayStack) Pop() (int, error) {
	return -1, nil
}

func (as *ArrayStack) Peek() (int, error) {
	return -1, nil
}

func (as *ArrayStack) Empty() bool {
	return as.tam == 0
}

func (as *ArrayStack) Size() int {
	return as.tam
}
