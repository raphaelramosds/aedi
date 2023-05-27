package helper

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

	// Análise do filhos
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
	return -1
}

func createBst(v []int) *BstNode {
	bstNode := &BstNode{}
	return bstNode
}
