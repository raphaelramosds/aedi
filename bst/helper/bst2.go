package helper

import "fmt"

func (node *BSTNode) Insert(value int) {
	if node.left == nil && node.right == nil && node.value == 0 {
		node.value = value
	} else {
		if value <= node.value {
			if node.left != nil {
				node.left.Insert(value)
			} else {
				node.left = &BSTNode{value: value}
			}
		} else {
			if node.right != nil {
				node.right.Insert(value)
			} else {
				node.right = &BSTNode{value: value}
			}
		}
	}
}

func (node *BSTNode) Search(value int) bool {
	if value == node.value {
		return true
	} else if value < node.value && node.left != nil {
		return node.left.Search(value)
	} else if value > node.value && node.right != nil {
		return node.right.Search(value)
	}
	return false
}

func (node *BSTNode) Min() int {
	if node.left == nil {
		return node.value
	} else {
		return node.left.Min()
	}
}

func (node *BSTNode) Max() int {
	if node.right == nil {
		return node.value
	} else {
		return node.right.Max()
	}
}

func (node *BSTNode) PrintPre() {
	fmt.Println(node.value)
	if node.left != nil {
		node.left.PrintPre()
	}
	if node.right != nil {
		node.right.PrintPre()
	}
}

func (node *BSTNode) PrintIn() {
	if node.left != nil {
		node.left.PrintIn()
	}
	fmt.Println(node.value)
	if node.right != nil {
		node.right.PrintIn()
	}
}

func (node *BSTNode) PrintPos() {
	if node.left != nil {
		node.left.PrintPos()
	}
	if node.right != nil {
		node.right.PrintPos()
	}
	fmt.Println(node.value)
}

func (node *BSTNode) Height() int {

	// Alturas da subárvore esquerda e direita
	left_height := 0
	right_height := 0

	// Calcula altura da subárvore direita
	if node.right != nil {
		return 1 + node.right.Height()
	}

	// Calcula altura da subárvore esquerda
	if node.left != nil {
		return 1 + node.left.Height()
	}

	// Retorna maior altura
	if left_height > right_height {
		return left_height
	} else {
		return right_height
	}
}

func (node *BSTNode) Remove(value int) *BSTNode {

	// Procurar nó a ser removido

	if value > node.value {
		node.right.Remove(value)
	} else if value < node.value {
		node.left.Remove(value)
	} else {

		// Nó encontrado

		if node.left == nil && node.right == nil { // Caso 1: nó folha (1 filho)
			// O pai dele apontará para nil
			return nil

		} else if node.right == nil && node.left != nil { // Caso 2: nó com um filho esquerdo
			// O pai dele apontará para o filho esquerdo do nó removido
			return node.left

		} else if node.right != nil && node.left == nil { // Caso 2: nó com um filho direito
			// O pai dele apontará para o filho direito do nó removido
			return node.right

		} else { // Caso 3: nó com dois filhos (Abordagem II)

			// Recupera menor valor da subárvore direita
			right_min := node.right.Min()

			// Copia esse valor no nó a ser removido
			node.value = right_min

			// Remove esse exato valor da subárvore direita
			node.right = node.right.Remove(right_min)

			return node
		}
	}

	return node
}
