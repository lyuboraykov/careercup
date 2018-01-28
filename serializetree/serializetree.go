package serializetree

import (
	"errors"
	"fmt"
)

// BinarySearchTree is a structure for searching values quickly
type BinarySearchTree struct {
	root *node
}

// NewTree constructs a tree
func NewTree() (tree *BinarySearchTree) {
	tree = &BinarySearchTree{
		nil,
	}
	return
}

// Insert adds a new value in the tree
func (t *BinarySearchTree) Insert(v int) error {
	if t.root == nil {
		t.root = newNode(v)
		return nil
	}
	return t.root.insert(v)
}

func (n *node) insert(v int) error {
	if n.value == v {
		return errors.New("Value is already present in the tree")
	}
	if v < n.value {
		if n.left == nil {
			n.left = newNode(v)
			return nil
		}
		return n.left.insert(v)
	}
	if n.right == nil {
		n.right = newNode(v)
		return nil
	}
	return n.right.insert(v)
}

// Marshal returns the binary search tree as json
func (t *BinarySearchTree) Marshal() string {
	result := `{"head": null}`
	if t.root != nil {
		result = fmt.Sprintf(`{"head": %s}`, t.root.Marshal())
	}
	return result
}

type node struct {
	value int
	left  *node
	right *node
}

func newNode(v int) *node {
	return &node{
		v,
		nil,
		nil,
	}
}

func (n *node) Marshal() string {
	left, right := "null", "null"
	if n.left != nil {
		left = n.left.Marshal()
	}
	if n.right != nil {
		right = n.right.Marshal()
	}
	return fmt.Sprintf(`{"val": %d, "left": %s, "right": %s}`, n.value, left, right)
}
