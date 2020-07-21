package bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	MinVal = -5
	MaxVal = 15
)

var (
	input = []int{9, 3, MinVal, 7, 2, MaxVal, 5, 12, 4}
)

func TestIntBinarySearchTree_InsertAndSearch(t *testing.T) {
	bst := NewIntBinarySearchTree(input)
	val := 13
	assert.False(t, bst.Search(val))
	bst.Insert(val)
	assert.True(t, bst.Search(val))
}

func TestIntBinarySearchTree_Remove(t *testing.T) {
	bst := NewIntBinarySearchTree(input)
	val := 2
	assert.True(t, bst.Search(val))
	bst.Remove(val)
	assert.False(t, bst.Search(val))
}

func TestIntBinarySearchTree_Min(t *testing.T) {
	bst := NewIntBinarySearchTree(input)
	minVal, _ := bst.Min()
	assert.Equal(t, MinVal, minVal)
	newMinVal := MinVal - 10
	bst.Insert(newMinVal)
	minVal, _ = bst.Min()
	assert.Equal(t, newMinVal, minVal)
}

func TestIntBinarySearchTree_Max(t *testing.T) {
	bst := NewIntBinarySearchTree(input)
	maxVal, _ := bst.Max()
	assert.Equal(t, MaxVal, maxVal)
	newMaxVal := MaxVal + 10
	bst.Insert(newMaxVal)
	maxVal, _ = bst.Max()
	assert.Equal(t, newMaxVal, maxVal)
}

func TestIntBinarySearchTree_EmptyTree(t *testing.T) {
	bst := NewIntBinarySearchTree([]int{})
	val := 42
	assert.False(t, bst.Search(val))
	bst.Remove(val)
	_, err := bst.Max()
	assert.Equal(t, err, ErrNoRootInTree)
	_, err = bst.Min()
	assert.Equal(t, err, ErrNoRootInTree)
}
