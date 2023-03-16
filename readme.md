# Data Structures and Algorithms (DCA0208)

All implementations are built in Golang, a modern programming language

## Lists

Lists implementations must follow this interface

```go
package list

type IList interface {
	Init(length int)
	Add(value int)
	AddOnIndex(value int, index int)
	Remove()
	RemoveOnIndex(index int)
	Get(value int)
	Set(value int, index int)
	Size()
}
```

**Note:** One must follow the set of rules below

- On structs, `tam` represents how many elements you've already added into the list. It's totally diferent from `len(list)` that represents how many elements you've asked memory to allocate

- Methods `Add(value int)` and `Remove()` addes/removes element into the end of the list, regardless if it is an ArrayList, LinkedList or DoublyLinkedList

- Methods `AddOnIndex(value int, index int)` should shift elements to right in order to fit the new element
