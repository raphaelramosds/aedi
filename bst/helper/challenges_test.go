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

func TestSizeWithRootNode(t *testing.T) {
	root := &BstNode{value: 5}
	result := root.Size()
	expected := 1
	if result != expected {
		t.Errorf("This three has only one node")
	}
}

func TestSize(t *testing.T) {
	root := &BstNode{value: 5}

	root.left = &BstNode{value: 2}
	root.left.left = &BstNode{value: 1}
	root.left.right = &BstNode{value: 4}

	root.right = &BstNode{value: 7}
	root.right.left = &BstNode{value: 1}
	root.right.right = &BstNode{value: 8}
	root.right.right.left = &BstNode{value: 3}
	root.right.right.right = &BstNode{value: 10}

	result := root.Size()
	expected := 9

	if result != expected {
		t.Errorf("Expected %d nodes, but got %d", expected, result)
	}
}

// TODO segmentation fault
func TestCreateBst(t *testing.T) {
	v := [6]int{5, 1, 2, 3, 4, 8}
	root := createBst(v[:])
	result := root.isBst()
	if !result {
		t.Errorf("This binary tree is a BST")
	}
}
