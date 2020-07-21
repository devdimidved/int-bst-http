// package bst provides functionality of binary search tree
package bst

import (
	"errors"
	"sync"
)

var (
	ErrNoRootInTree = errors.New("the tree has no root")
)

// Node is an element of a binary search tree
type Node struct {
	val   int
	left  *Node
	right *Node
}

// IntBinarySearchTree is an implementation of a binary search tree on integers
type IntBinarySearchTree struct {
	mtx  sync.RWMutex
	root *Node
}

// NewIntBinarySearchTree creates a new instance of IntBinarySearchTree
// and populates it with provided integers
func NewIntBinarySearchTree(ints []int) *IntBinarySearchTree {
	t := &IntBinarySearchTree{}
	for i := range ints {
		t.Insert(ints[i])
	}
	return t
}

// Insert inserts an element into the tree
func (t *IntBinarySearchTree) Insert(v int) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	n := &Node{v, nil, nil}
	if t.root == nil {
		t.root = n
	} else {
		insert(t.root, n)
	}
}

// internal recursive function to insert a new element in the tree
func insert(existing, new *Node) {
	if new.val < existing.val {
		if existing.left == nil {
			existing.left = new
		} else {
			insert(existing.left, new)
		}
	} else {
		if existing.right == nil {
			existing.right = new
		} else {
			insert(existing.right, new)
		}
	}
}

// Search returns true if v exists in the tree
func (t *IntBinarySearchTree) Search(v int) bool {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	return search(t.root, v)
}

// internal recursive function to search value in the tree
func search(node *Node, v int) bool {
	if node == nil {
		return false
	}
	if v < node.val {
		return search(node.left, v)
	}
	if v > node.val {
		return search(node.right, v)
	}
	return true
}

// Min returns min value in the tree or error in case the tree is empty
func (t *IntBinarySearchTree) Min() (int, error) {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	if t.root == nil {
		return 0, ErrNoRootInTree
	}
	return min(t.root), nil
}

// internal recursive function to find min value in the tree
func min(node *Node) int {
	if node.left == nil {
		return node.val
	}
	return min(node.left)
}

// Max returns max value in the non-empty tree or error in case the tree is empty
func (t *IntBinarySearchTree) Max() (int, error) {
	t.mtx.RLock()
	defer t.mtx.RUnlock()
	if t.root == nil {
		return 0, ErrNoRootInTree
	}
	return max(t.root), nil
}

// internal recursive function to find max value in the tree
func max(node *Node) int {
	if node.right == nil {
		return node.val
	}
	return max(node.right)
}

// Remove removes a value from the tree
func (t *IntBinarySearchTree) Remove(v int) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	remove(t.root, v)
}

// internal recursive function to remove value from the tree
func remove(node *Node, v int) *Node {
	if node == nil {
		return nil
	}
	if v < node.val {
		node.left = remove(node.left, v)
		return node
	}
	if v > node.val {
		node.right = remove(node.right, v)
		return node
	}
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}
	minRightSide := node.right
	for {
		//find smallest value on the right side
		if minRightSide != nil && minRightSide.left != nil {
			minRightSide = minRightSide.left
		} else {
			break
		}
	}
	node.val = minRightSide.val
	node.right = remove(node.right, node.val)
	return node
}
