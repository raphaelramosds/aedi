package helper

// Base structures

type Node struct {
	next *Node
	val  int
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	tam  int
}

// Internal functions

func initNode(val int) *Node {
	newNode := new(Node)
	newNode.val = val
	return newNode
}

// Public functions

func (dll *DoublyLinkedList) Init() {
	dll.head = nil
	dll.tam = 0
	dll.tail = nil
}

func (dll *DoublyLinkedList) Add(value int) {

	newNode := initNode(value)

	if dll.tam == 0 {
		dll.head = newNode
		dll.tail = newNode
		newNode.prev = dll.head
	} else {
		aux := dll.tail
		aux.next = newNode
		dll.tail = newNode
	}

	dll.tam++
}

func (dll *DoublyLinkedList) Get(value int) int {

	if value < 0 || value >= dll.tam {
		return -1
	}

	aux := dll.head

	for i := 0; aux != nil; i++ {
		if i == value {
			return aux.val
		}
		aux = aux.next
	}

	return 0
}

func (dll *DoublyLinkedList) Size() int {
	return dll.tam
}
