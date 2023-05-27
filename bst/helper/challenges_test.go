package helper

import "testing"

func TestIsBstWithANonBST(t *testing.T) {

	root := &BstNode{value: 5}

	root.left = &BstNode{value: 2}
	root.left.left = &BstNode{value: 1}
	root.left.right = &BstNode{value: 4}

	root.right = &BstNode{value: 7}
	root.right.left = &BstNode{value: 1}
	root.right.right = &BstNode{value: 8}
	root.right.right.left = &BstNode{value: 10}
	root.right.right.right = &BstNode{value: 3}

	result := root.isBst()

	if result {
		t.Errorf("This binary tree isn't BST")
	}
}

func TestIsBstWithABST(t *testing.T) {
	root := &BstNode{value: 5}

	root.left = &BstNode{value: 2}
	root.left.left = &BstNode{value: 1}
	root.left.right = &BstNode{value: 4}

	root.right = &BstNode{value: 7}
	root.right.left = &BstNode{value: 1}
	root.right.right = &BstNode{value: 8}
	root.right.right.left = &BstNode{value: 3}
	root.right.right.right = &BstNode{value: 10}

	result := root.isBst()

	if !result {
		t.Errorf("This binary tree is a BST")
	}
}
