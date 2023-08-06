// Package tree is a package for work with Binary trees.
package tree

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

const (
	// Desc specifies the sort direction to be descending.
	Desc direction = "desc"
	// Asc specifies the sort direction to be ascending.
	Asc direction = "asc"
)

// direction is a type which uses to set the direction (Asc or Desc).
type direction string

type Tree[V constraints.Ordered] struct {
	root *node[V]
}

// New is a function for creation empty tree
// - param should be `ordered type` (`int`, `string`, `float` etc)
func New[V constraints.Ordered]() *Tree[V] {
	return &Tree[V]{}
}

// NewWithElement is a function for creation tree with one element
// - param should be `ordered type` (`int`, `string`, `float` etc)
func NewWithElement[V constraints.Ordered](key V, value any) *Tree[V] {
	return &Tree[V]{
		root: &node[V]{
			element: element[V]{
				key:   key,
				value: value,
			},
		},
	}
}

// Insert is a function for inserting element into node
// - param key should be `ordered type` (`int`, `string`, `float` etc.)
// - param value can be any type
func (t *Tree[V]) Insert(key V, value any) {
	n := &node[V]{
		element: element[V]{
			key:   key,
			value: value,
		},
	}

	if t.root == nil {
		t.root = n
		return
	}

	t.root.insertNode(n)
}

// Min is a function for searching min element in tree (by key).
func (t *Tree[V]) Min() V {
	var result V
	n := t.root
	if n == nil {
		return result
	}

	for n.left != nil {
		n = n.left
	}

	return n.element.key
}

// Max is a function for searching max element in tree (by key).
func (t *Tree[V]) Max() V {
	var result V
	n := t.root
	if n == nil {
		return result
	}

	for n.right != nil {
		n = n.right
	}

	return n.element.key
}

// Exists is a function for searching element in node. If element exists in tree- return true, else - false
// - param key should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) Exists(key V) bool {
	searchNode := search(t.root, key)
	if searchNode == nil {
		return false
	}

	return true
}

// GetValue is a function for searching element in node and returning value of this element
// - param key should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) GetValue(key V) (any, error) {
	var result any
	searchNode := search(t.root, key)
	if searchNode == nil {
		return result, errors.New(fmt.Sprintf("element with key %v not found", key))
	}

	return searchNode.element.value, nil
}

// InOrderTreeWalk is a function for getting ordered array of tree's elements.
// - param key should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) InOrderTreeWalk(d direction) []V {
	if t.root == nil {
		return nil
	}

	return inOrderTreeWalk(t.root, d)
}

// PreOrderSuccessor is a function for searching preOrder key for income element (if we found it by input key param)
// - param key should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) PreOrderSuccessor(key V) (V, error) {
	var result V
	searchNode := search(t.root, key)
	if searchNode == nil || searchNode.parent == nil {
		return result, errors.New(fmt.Sprintf("PreOrderSuccessor for key %v not found", key))
	}

	return searchNode.parent.element.key, nil
}

// PostOrderSuccessor is a function for searching postOrder key for income element (if we found it by input key param)
// - param key should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) PostOrderSuccessor(key V) (V, error) {
	var result V
	searchNode := search(t.root, key)
	if searchNode == nil || searchNode.right == nil {
		return result, errors.New(fmt.Sprintf("postOrderSuccessor for key %v not found", key))
	}

	return searchNode.right.element.key, nil
}

// Delete is a function for deleting node in node
// - param key should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) Delete(key V) {
	delNode := search(t.root, key)
	if delNode == nil {
		return
	}

	if delNode == t.root && t.root.hasNoChildren() {
		t.root = nil
		return
	}

	// first case (node without children)
	if delNode.element.key == key && delNode.hasNoChildren() {
		if delNode.parent.right.element.key == key {
			delNode.parent.right = nil
			return
		}

		delNode.parent.left = nil
		return
	}

	// second case
	if delNode.left == nil && delNode.right != nil {
		delNode.right.parent = delNode.parent

		if delNode == t.root {
			t.root = t.root.right
			return
		}

		if delNode.parent.right == delNode {
			delNode.parent.right = delNode.right
			return
		}

		delNode.parent.left = delNode.left
		return
	}

	if delNode.left != nil && delNode.right == nil {
		delNode.left.parent = delNode.parent

		if delNode == t.root {
			t.root = t.root.left
			return
		}

		if delNode.parent.right == delNode {
			delNode.parent.right = delNode.left
			return
		}

		delNode.parent.left = delNode.left
		return
	}

	//third case
	m := min(delNode.right)
	minElement := m.element

	if m.parent == delNode {
		m.parent.right = nil
	} else {
		m.parent.left = nil
	}

	m.parent = nil

	delNode.element = minElement

	return
}
