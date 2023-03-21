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

## Pros and caveats

<style>

.custom-table tr td ~ td{ text-align:center; }

.custom-table tr:nth-child(1) td:nth-child(2),
.custom-table tr:nth-child(1) td:nth-child(4),
.custom-table tr:nth-child(2) td:nth-child(4),
.custom-table tr:nth-child(3) td:nth-child(4),
.custom-table tr:nth-child(3) td:nth-child(3),
.custom-table tr:nth-child(5) td:nth-child(2),
.custom-table tr:nth-child(6) td:nth-child(2),
.custom-table tr:nth-child(7) td:nth-child(2),
.custom-table tr:nth-child(7) td:nth-child(4),
.custom-table tr:nth-child(7) td:nth-child(3),
.custom-table tr:nth-child(8) td:nth-child(2),
.custom-table tr:nth-child(8) td:nth-child(4),
.custom-table tr:nth-child(9) td:nth-child(4),
.custom-table tr:nth-child(9) td:nth-child(3)  {
	background-color: green;
}
	
</style>

<div class="custom-table">

| **Operação**                                       | **ArrayList** | **LinkedList** | **DoublyLinkedList** |
|----------------------------------------------------|---------------|----------------|----------------------|
| Inserir elemento no fim (ainda há espaço no array) |      O(1)     |      O(n)      |         O(1)         |
| Inserir elemento no fim (não há espaço no array)   |      O(n)     |      O(n)      |         O(1)         |
| Inserir elemento no início                         |      O(n)     |      O(1)      |         O(1)         |
| Inserir elemento em posição                        |      O(n)     |      O(n)      |         O(n)         |
| Obter elemento em posição                          |      O(1)     |      O(n)      |         O(n)         |
| Atualizar elemento em posição                      |      O(1)     |      O(n)      |         O(n)         |
| Obter tamanho da lista                             |      O(1)     |      O(1)      |         O(1)         |
| Remover elemento no fim                            |      O(1)     |      O(n)      |         O(1)         |
| Remover elemento no início                         |      O(n)     |      O(1)      |         O(1)         |
| Remover elemento em posição                        |      O(n)     |      O(n)      |         O(n)         |

</div>