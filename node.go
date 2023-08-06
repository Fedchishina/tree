package tree

import "golang.org/x/exp/constraints"

type element[V constraints.Ordered] struct {
	key   V
	value any
}

// node is the structure of tree's node.
// node's key is any ordered type for type of
// node's value has type any
type node[V constraints.Ordered] struct {
	element element[V]
	parent  *node[V]
	left    *node[V]
	right   *node[V]
}

func (n *node[V]) hasNoChildren() bool {
	return n.left == nil && n.right == nil
}

func (n *node[V]) insertNode(newNode *node[V]) {
	if newNode.element.key < n.element.key {
		if n.left == nil {
			addLeaf[V](newNode, n, &n.left)
			return
		}
		n.left.insertNode(newNode)
		return
	}

	if n.right == nil {
		addLeaf[V](newNode, n, &n.right)
		return
	}

	n.right.insertNode(newNode)
}

func addLeaf[V constraints.Ordered](newNode, parentNode *node[V], nodePlace **node[V]) {
	*nodePlace = newNode
	newNode.parent = parentNode
}

func inOrderTreeWalk[V constraints.Ordered](n *node[V], d direction) []V {
	if n == nil {
		return []V{}
	}

	left := inOrderTreeWalk(n.left, d)
	right := inOrderTreeWalk(n.right, d)

	output := make([]V, 0)
	if d == Asc {
		return append(output, append(append(left, n.element.key), right...)...)
	}

	return append(output, append(append(right, n.element.key), left...)...)
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
