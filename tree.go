// Package tree is a package for work with Binary trees.
package tree

import (
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
			element: Element[V]{
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
		element: Element[V]{
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
func (t *Tree[V]) Min() *Element[V] {
	n := t.root
	if n == nil {
		return nil
	}

	for n.left != nil {
		n = n.left
	}

	return &n.element
}

// Max is a function for searching max element in tree (by key).
func (t *Tree[V]) Max() *Element[V] {
	n := t.root
	if n == nil {
		return nil
	}

	for n.right != nil {
		n = n.right
	}

	return &n.element
}

// Search is a function for searching element in node.
// - param key should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) Search(key V) *Element[V] {
	searchNode := search(t.root, key)
	if searchNode == nil {
		return nil
	}

	return &searchNode.element
}

// InOrderTreeWalk is a function for getting ordered array of tree's elements.
// - param key should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) InOrderTreeWalk(d direction) []Element[V] {
	if t.root == nil {
		return nil
	}

	return inOrderTreeWalk(t.root, d)
}

// PreOrderSuccessor is a function for searching preOrder element for income element (if we found it by input key param)
// - param key should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) PreOrderSuccessor(key V) *Element[V] {
	searchNode := search(t.root, key)
	if searchNode == nil || searchNode.parent == nil {
		return nil
	}

	return &searchNode.parent.element
}

// PostOrderSuccessor is a function for searching postOrder element for income element (if we found it by input key param)
// - param key should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) PostOrderSuccessor(key V) *Element[V] {
	searchNode := search(t.root, key)
	if searchNode == nil || searchNode.right == nil {
		return nil
	}

	return &searchNode.right.element
}

// Delete is a function for deleting node in node
// - param key should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) Delete(key V) {
	if t.root != nil && t.root.element.key == key && t.root.left == nil && t.root.right == nil {
		t.root = nil
		return
	}

	delNode := search(t.root, key)
	if delNode == nil {
		return
	}

	// first case (node without children)
	if delNode.element.key == key && delNode.left == nil && delNode.right == nil {
		if delNode.parent.right.element.key == key {
			delNode.parent.right = nil
			return
		}

		delNode.parent.left = nil
		return
	}

	// second case
	if delNode.left == nil && delNode.right != nil && delNode.parent == nil {
		delNode.right.parent = nil
		t.root = delNode.right
		return
	}
	if delNode.left != nil && delNode.right == nil && delNode.parent == nil {
		delNode.left.parent = nil
		t.root = delNode.left
		return
	}

	if delNode.left == nil && delNode.right != nil && delNode.parent != nil {
		delNode.right.parent = delNode.parent
		if delNode.parent.right.element.key == key {
			delNode.parent.right = delNode.right
			return
		}

		delNode.parent.left = delNode.left
		return
	}

	if delNode.left != nil && delNode.right == nil && delNode.parent != nil {
		delNode.left.parent = delNode.parent
		if delNode.parent.right.element.key == key {
			delNode.parent.right = delNode.left
			return
		}

		delNode.parent.left = delNode.left
		return
	}

	//third case
	m := min(delNode.right)
	minElement := m.element
	t.Delete(m.element.key)

	delNode.element = minElement

	return
}
