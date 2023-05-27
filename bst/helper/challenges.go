package helper

import (
	"sort"
)

type BstNode struct {
	left  *BstNode
	value int
	right *BstNode
}

func (bstNode *BstNode) isBst() bool {
	// Caso base: a folha é uma bst, por padrão
	if bstNode.left == nil && bstNode.right == nil {
		return true
	}

	// Inspeção do filhos esquerdo e direito
	if bstNode.left != nil && bstNode.value < bstNode.left.value {
		return false
	}

	if bstNode.right != nil && bstNode.value > bstNode.right.value {
		return false
	}

	// Uma subárvore é BST se o seu nó esquerdo for menor que o pai E o nó direito é maior que o pai
	return bstNode.right.isBst() && bstNode.left.isBst()
}

func (bstNode *BstNode) Size() int {
	// Se visitei um nó, a contagem deve começar com 1
	count := 1

	// Inspeção dos filhos esquerdo e direito

	if bstNode.left != nil {
		// Acumule a quantidade nós da subárvore anterior
		count += bstNode.left.Size()
	}

	if bstNode.right != nil {
		count += bstNode.right.Size()
	}

	// Retorna quantidade de nós da subárvore
	return count
}

func (bstNode *BstNode) addBstNode(v []int, start int, end int) {
	if start < end {
		mid := (start + end) / 2
		bstNode.value = v[mid]
		bstNode.left.addBstNode(v, start, mid-1)
		bstNode.right.addBstNode(v, mid+1, end)
	}
}

func createBst(v []int) *BstNode {
	sort.Ints(v)
	root := &BstNode{}
	root.addBstNode(v, 0, len(v)-1)
	return root
}
