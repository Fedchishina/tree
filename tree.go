// Package tree is a package for work with Binary trees.
package tree

import (
	"errors"

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

// tree is the structure of tree. You can use any ordered type for type of tree's values
type tree[V constraints.Ordered] struct {
	value V
	left  *tree[V]
	right *tree[V]
}

// CreateNode is a function for creation tree with one node
// - param should be `ordered type` (`int`, `string`, `float` etc)
func CreateNode[V constraints.Ordered](val V) *tree[V] {
	return &tree[V]{
		value: val,
		left:  nil,
		right: nil,
	}
}

// InOrderTreeWalk is a function when you can make tree traversal.
//   - param should be `type direction` (`direction.Asc` or `direction.Desc`)
func (t *tree[V]) InOrderTreeWalk(d direction) []V {
	if t == nil {
		return nil
	}

	left := t.left.InOrderTreeWalk(d)
	right := t.right.InOrderTreeWalk(d)

	output := make([]V, 0)

	if d == Asc {
		output = append(output, left...)
	} else {
		output = append(output, right...)
	}

	output = append(output, t.value)

	if d == Asc {
		output = append(output, right...)
	} else {
		output = append(output, left...)
	}

	return output
}

// Search is a function for searching element in tree.
//   - param should be `ordered type` (`int`, `string`, `float` etc)
func (t *tree[V]) Search(val V) *tree[V] {
	for t != nil && val != t.value {
		if val < t.value {
			t = t.left
		} else {
			t = t.right
		}
	}

	return t
}

// Min is a function for searching min element in tree.
func (t *tree[V]) Min() *tree[V] {
	for t.left != nil {
		t = t.left
	}

	return t
}

// Max is a function for searching max element in tree.
func (t *tree[V]) Max() *tree[V] {
	for t.right != nil {
		t = t.right
	}

	return t
}

// PreOrderSuccessor is a function for searching preOrder element for income element of tree
//   - param should be `type tree`
func (t *tree[V]) PreOrderSuccessor(el *tree[V]) *tree[V] {
	if el.left != nil {
		return el.left.Max()
	}

	var r tree[V]
	for t != nil {
		if t.value > el.value {
			t = t.left
		} else if t.value < el.value {
			r = *t
			t = t.right
		} else {
			break
		}
	}
	return &r
}

// PostOrderSuccessor is a function for searching postOrder element for income element of tree
//   - param should be `type tree`
func (t *tree[V]) PostOrderSuccessor(el *tree[V]) *tree[V] {
	if el.right != nil {
		return el.right.Min()
	}

	var r tree[V]
	for t != nil {
		if t.value < el.value {
			t = t.right
		} else if t.value > el.value {
			r = *t
			t = t.left
		} else {
			break
		}
	}
	return &r
}

// Insert is a function for inserting element into tree
//   - param should be `ordered type` (`int`, `string`, `float` etc.)
func (t *tree[V]) Insert(val V) {
	node := tree[V]{
		value: val,
	}

	if node.value < t.value {
		if t.left == nil {
			t.left = &node
		} else {
			t.left.Insert(val)
		}
	} else {
		if t.right == nil {
			t.right = &node
		} else {
			t.right.Insert(val)
		}
	}
}

// Parent is a function for founding parent for value
//   - param should be `ordered type` (`int`, `string`, `float` etc.)
//
// if tree does not contain node with input val function returns nil
func (t *tree[V]) Parent(val V) *tree[V] {
	if t.left == nil && t.right == nil {
		return nil
	}

	var parent *tree[V]
	for t != nil && val != t.value {
		parent = t
		if val < t.value {
			t = t.left
		} else {
			t = t.right
		}
	}

	if t == nil {
		return nil
	}

	return parent
}

// Delete is a function for deleting node in tree
//   - param should be `ordered type` (`int`, `string`, `float` etc.)
func (t *tree[V]) Delete(val V) error {
	//tree with one node
	if t.left == nil && t.right == nil && val == t.value {
		t = nil

		return nil
	}

	delNode := t.Search(val)
	if delNode == nil {
		return errors.New("node for deleting not found")
	}

	parent := t.Parent(val)

	// first case
	if delNode.left == nil && delNode.right == nil {
		if parent.left.value == val {
			parent.left = nil
		} else {
			parent.right = nil
		}

		return nil
	}

	// second case
	if delNode.left == nil && delNode.right != nil {
		if parent == nil {
			t = delNode.right

			return nil
		}

		if parent.left.value == val {
			parent.left = delNode.right
		} else {
			parent.right = delNode.right
		}

		return nil
	}

	if delNode.left != nil && delNode.right == nil {
		if parent == nil {
			t = delNode.left

			return nil
		}
		if parent.left.value == val {
			parent.left = delNode.left
		} else {
			parent.right = delNode.left
		}

		return nil
	}

	//third case
	min := delNode.right.Min()
	minVal := min.value
	err := t.Delete(min.value)

	if err != nil {
		return err
	}

	delNode.value = minVal

	return nil
}
