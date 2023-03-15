# Data Structures and Algorithms (DCA0208)

All implementations are built with Golang, a modern programming language

## Lists

Lists implementations must follow this interface

```go
package list

type IList interface {
	Add(value int)
	AddOnIndex(value int, index int)
	Remove()
	RemoveOnIndex(index int)
	Get(value int)
	Set(value int, index int)
	Size()
}
```

**Note:** `AddOnIndex(value int, index int)` should shift elements to right in order to fit the new element