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
