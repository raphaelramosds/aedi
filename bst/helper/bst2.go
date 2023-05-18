package helper

type BinarySearchTree struct {
	root *NodeBST
}

func (bst *BinarySearchTree) Insert(value int) {
	var parent *NodeBST

	curr := bst.root
	new_node := NodeBST{value, nil, nil}

	for curr != nil {
		parent = curr
		if value < curr.data {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	if parent == nil {
		bst.root = &new_node
	} else if value < parent.data {
		parent.left = &new_node
	} else {
		parent.right = &new_node
	}
}

func (bst *BinarySearchTree) Search(value int) bool {
	curr := bst.root
	for curr != nil {
		if value == curr.data {
			return true
		} else if value < curr.data {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return false
}
