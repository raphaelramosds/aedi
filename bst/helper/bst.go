package helper

import "fmt"

type NodeBST struct {
	data  int
	left  *NodeBST
	right *NodeBST
}

func Insert(root **NodeBST, value int) {
	if *root == nil {
		*root = &NodeBST{value, nil, nil}
	} else {
		if value > (*root).data {
			Insert(&(*root).right, value)
		} else {
			Insert(&(*root).left, value)
		}
	}
}

func Search(root **NodeBST, value int) bool {
	if *root == nil {
		return false
	} else {
		if (*root).data == value {
			return true
		} else if (*root).data > value {
			return Search(&(*root).left, value)
		} else {
			return Search(&(*root).right, value)
		}
	}
}

func PreOrderFashion(root **NodeBST) {
	if *root != nil {
		fmt.Printf("%d ", (*root).data)
		InOrderFashion(&(*root).left)
		InOrderFashion(&(*root).right)
	}
}

func InOrderFashion(root **NodeBST) {
	if *root != nil {
		InOrderFashion(&(*root).left)
		fmt.Printf("%d ", (*root).data)
		InOrderFashion(&(*root).right)
	}
}

func PostOrderFashion(root **NodeBST) {
	if *root != nil {
		InOrderFashion(&(*root).left)
		InOrderFashion(&(*root).right)
		fmt.Printf("%d ", (*root).data)
	}
}
