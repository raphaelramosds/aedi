package helper

import "fmt"

type BSTNode struct {
	left  *BSTNode
	right *BSTNode
	value int
}

func Insert(root **BSTNode, value int) {
	if *root == nil {
		*root = &BSTNode{nil, nil, value}
	} else {
		if value > (*root).value {
			Insert(&(*root).right, value)
		} else {
			Insert(&(*root).left, value)
		}
	}
}

func Search(root **BSTNode, value int) bool {
	if *root == nil {
		return false
	} else {
		if (*root).value == value {
			return true
		} else if (*root).value > value {
			return Search(&(*root).left, value)
		} else {
			return Search(&(*root).right, value)
		}
	}
}

func PrintPre(root **BSTNode) {
	if *root != nil {
		fmt.Printf("%d ", (*root).value)
		PrintPre(&(*root).left)
		PrintPre(&(*root).right)
	}
}

func PrintIn(root **BSTNode) {
	if *root != nil {
		PrintIn(&(*root).left)
		fmt.Printf("%d ", (*root).value)
		PrintIn(&(*root).right)
	}
}

func PrintPost(root **BSTNode) {
	if *root != nil {
		PrintPost(&(*root).left)
		PrintPost(&(*root).right)
		fmt.Printf("%d ", (*root).value)
	}
}
