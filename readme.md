# Data Structures

All implementations are built in Golang, a modern programming language

**Tips:** Below some tips for those implementations

- The parameter `size` on Abstract Data Types means how many elements you've inserted into the data structure! It differs from `len()` that represents how many spaces of memory you've allocated for storing your elements

- Index verification is presented as follow. Note that is totally safe if user entries with `index = 0`. However, get, remove and set from `index = list.size` is a problem because this index doesn't exist yet

```go
if (index >= 0 && index < list.size) {
	// get, remove, set
}
```

But, add into `index = list.size` is a safe operation! It means that you want to put an element into a new index. Therefore, for adding operations

```go
if (index >= 0 && index <= list.size) {
	// add
}
```

- It's easier to iterate over linked lists using for loops. For exemple, if you want to reach the n-th element, you only want to iterate until de n - 1 index. That's because your iterate variable `curr` always assumes the next address. So, once the (n-1)-th element is reached, `curr` automatically take the address of the n-th element

```go 
curr := list.head

for i := 0; i < n; i++ {
	curr = curr.next
}
```

## Lists

Lists implementations must follow this interface

```go
type IList interface {
	Add(value int)
	AddOnIndex(value int, index int) error
	RemoveOnIndex(index int) error
	Get(index int) (int, error)
	Set(value int, index int) error
	Size() int
}
```

**Note:** One must follow the set of rules below

- Methods `Add(value int)` and `Remove()` addes/removes element into the end of the list, regardless if it is an ArrayList, LinkedList or DoublyLinkedList

- Methods `AddOnIndex(value int, index int)` should shift elements to right in order to fit the new element


### Pros and caveats

| **Operation**                                      | **ArrayList** | **LinkedList** | **DoublyLinkedList** |
|:---------------------------------------------------|---------------|----------------|----------------------|
| Add at end                                         |      O(1)     |      O(n)      |         O(1)         |
| Add at end (full filled array)                     |      O(n)     |      O(n)      |         O(1)         |
| Add at begin                                       |      O(n)     |      O(1)      |         O(1)         |
| Add at index                                       |      O(n)     |      O(n)      |         O(n)         |
| Get                                                |      O(1)     |      O(n)      |         O(n)         |
| Set                                                |      O(1)     |      O(n)      |         O(n)         |
| Size                                               |      O(1)     |      O(1)      |         O(1)         |
| Remove from the end                                |      O(1)     |      O(n)      |         O(1)         |
| Remove from the beginning                          |      O(n)     |      O(1)      |         O(1)         |
| Remove from index                                  |      O(n)     |      O(n)      |         O(n)         |


## Stacks

Stacks implementations must follow this interface

```go
type IStack interface {
	Push(value int)
	Pop() (int, error)
	Peek() (int, error)
	Empty() bool
	Size() int
}
```