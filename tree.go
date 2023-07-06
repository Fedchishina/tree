package main

import (
	"golang.org/x/exp/constraints"
)

const (
	// Desc specifies the sort direction to be descending.
	Desc direction = "desc"
	// Asc specifies the sort direction to be ascending.
	Asc direction = "asc"
)

// direction is a type which uses to set the sort direction.
type direction string

type tree[V constraints.Ordered] struct {
	value V
	left  *tree[V]
	right *tree[V]
}

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

func (t *tree[V]) Min() *tree[V] {
	for t.left != nil {
		t = t.left
	}

	return t
}

func (t *tree[V]) Max() *tree[V] {
	for t.right != nil {
		t = t.right
	}

	return t
}

func (t *tree[V]) preOrderSuccessor(el *tree[V]) *tree[V] {
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

func (t *tree[V]) postOrderSuccessor(el *tree[V]) *tree[V] {
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

func (t *tree[V]) Insert(val V) {
	elem := tree[V]{
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
