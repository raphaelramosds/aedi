package helper

type BSTNode struct {
	left   *BSTNode
	value  int
	height int
	bf     int
	right  *BSTNode
}

func (node *BSTNode) RotRight() *BSTNode {
	left := node.left
	node.left = left.right
	left.right = node
	node.UpdateProperties()
	left.UpdateProperties()
	return left
}

func (node *BSTNode) RotLeft() *BSTNode {
	right := node.right
	node.right = right.left
	right.left = node
	node.UpdateProperties()
	right.UpdateProperties()
	return right
}

func (node *BSTNode) UpdateProperties() {
	heightRight := 0
	heightLeft := 0
	if node.right == nil && node.left == nil {
		node.height = 0
	} else {
		if node.right != nil {
			heightRight = node.right.height
		}
		if node.left != nil {
			heightLeft = node.left.height
		}
		if heightRight > heightLeft {
			node.height = 1 + heightRight
		} else {
			node.height = 1 + heightLeft
		}
	}

	node.bf = heightRight - heightLeft
}
