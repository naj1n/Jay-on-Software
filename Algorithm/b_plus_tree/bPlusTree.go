package main

import (
	"fmt"
)

// TreeNode represents a node in the B+ tree
type TreeNode struct {
	keys       []int
	children   []*TreeNode
	isLeaf     bool
	next       *TreeNode // Pointer to next leaf node
}

// BPlusTree represents the B+ tree structure
type BPlusTree struct {
	root *TreeNode
	t    int // Maximum number of children a node can have
}

// insert inserts a key into the B+ tree
func (tree *BPlusTree) insert(key int) {
	if tree.root == nil {
		tree.root = &TreeNode{keys: []int{key}, isLeaf: true}
		return
	}

	// Start insertion from the root
	tree.insertNode(tree.root, key)
}

// insertNode handles the logic of inserting a key into a node
func (tree *BPlusTree) insertNode(node *TreeNode, key int) {
	// Find position to insert key
	i := 0
	for i < len(node.keys) && key > node.keys[i] {
		i++
	}

	// Leaf node
	if node.isLeaf {
		node.keys = append(node.keys[:i], append([]int{key}, node.keys[i:]...)...)
		if len(node.keys) > tree.t-1 {
			tree.splitLeafNode(node)
		}
	} else { // Internal node
		tree.insertNode(node.children[i], key)
	}
}

// splitLeafNode splits a leaf node into two
func (tree *BPlusTree) splitLeafNode(node *TreeNode) {
	mid := len(node.keys) / 2
	newLeaf := &TreeNode{
		keys:   node.keys[mid:],
		isLeaf: true,
		next:   node.next,
	}
	node.keys = node.keys[:mid]
	node.next = newLeaf

	// TODO: Insert new leaf's first key into parent
}

func main() {
	tree := &BPlusTree{t: 3}
	keys := []int{10, 20, 5, 6, 12, 30, 7}

	for _, key := range keys {
		tree.insert(key)
	}

	fmt.Println("B+ Tree constructed with elements:", keys)
	// Output will depend on how the tree structure is printed
}
