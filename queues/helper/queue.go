package helper

type IQueue interface {
	Enqueue(value int)
	Dequeue() (int, error)
	Peek() (int, error)
	Empty() bool
	Size() int
}
