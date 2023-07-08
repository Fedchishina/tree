// Package tree is a package for work with trees.
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

// Tree is the structure of tree. You can use any ordered type for type of tree's values
type Tree[V constraints.Ordered] struct {
	value V
	left  *Tree[V]
	right *Tree[V]
}

// InOrderTreeWalk is a function when you can make Tree traversal.
//   - param should be `type direction` (`direction.Asc` or `direction.Desc`)
func (t *Tree[V]) InOrderTreeWalk(d direction) []V {
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
func (t *Tree[V]) Search(val V) *Tree[V] {
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
func (t *Tree[V]) Min() *Tree[V] {
	for t.left != nil {
		t = t.left
	}

	return t
}

// Max is a function for searching max element in tree.
func (t *Tree[V]) Max() *Tree[V] {
	for t.right != nil {
		t = t.right
	}

	return t
}

// PreOrderSuccessor is a function for searching preOrder element for income element of tree
//   - param should be `type Tree`
func (t *Tree[V]) PreOrderSuccessor(el *Tree[V]) *Tree[V] {
	if el.left != nil {
		return el.left.Max()
	}

	var r Tree[V]
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
//   - param should be `type Tree`
func (t *Tree[V]) PostOrderSuccessor(el *Tree[V]) *Tree[V] {
	if el.right != nil {
		return el.right.Min()
	}

	var r Tree[V]
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
//   - param should be `ordered type` (`int`, `string`, `float` etc)
func (t *Tree[V]) Insert(val V) {
	elem := Tree[V]{
		value: val,
	}
	if elem.value < t.value {
		if t.left == nil {
			t.left = &elem
		} else {
			t.left.Insert(val)
		}
	} else {
		if t.right == nil {
			t.right = &elem
		} else {
			t.right.Insert(val)
		}
	}
}
