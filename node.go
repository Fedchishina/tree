package tree

import "golang.org/x/exp/constraints"

type Element[V constraints.Ordered] struct {
	key   V
	value any
}

// node is the structure of tree's node.
// node's key is any ordered type for type of
// node's value has type any
type node[V constraints.Ordered] struct {
	element Element[V]
	parent  *node[V]
	left    *node[V]
	right   *node[V]
}

func (n *node[V]) insertNode(newNode *node[V]) {
	if newNode.element.key < n.element.key {
		if n.left == nil {
			n.left = newNode
			newNode.parent = n
			return
		}
		n.left.insertNode(newNode)
		return
	}

	if n.right == nil {
		n.right = newNode
		newNode.parent = n
		return
	}

	n.right.insertNode(newNode)
}

func inOrderTreeWalk[V constraints.Ordered](n *node[V], d direction) []Element[V] {
	if n == nil {
		return []Element[V]{}
	}

	left := inOrderTreeWalk(n.left, d)
	right := inOrderTreeWalk(n.right, d)

	output := make([]Element[V], 0)
	if d == Asc {
		return append(output, append(append(left, n.element), right...)...)
	}

	return append(output, append(append(right, n.element), left...)...)
}

func search[V constraints.Ordered](n *node[V], key V) *node[V] {
	for n != nil && key != n.element.key {
		if key < n.element.key {
			n = n.left
			continue
		}
		n = n.right
	}

	return n
}

func min[V constraints.Ordered](n *node[V]) *node[V] {
	if n == nil {
		return nil
	}

	for n.left != nil {
		n = n.left
	}

	return n
}
