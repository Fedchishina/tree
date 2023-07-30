package tree

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestNew(t *testing.T) {
	type testCase[V constraints.Ordered] struct {
		name string
		want *Tree[V]
	}
	testInt := testCase[int]{
		name: "int empty tree",
		want: &Tree[int]{root: nil},
	}
	t.Run(testInt.name, func(t *testing.T) {
		if got := New[int](); !reflect.DeepEqual(got, testInt.want) {
			t.Errorf("CreateNode() = %v, want %v", got, testInt.want)
		}
	})

	testString := testCase[string]{
		name: "int empty tree",
		want: &Tree[string]{root: nil},
	}
	t.Run(testString.name, func(t *testing.T) {
		if got := New[int](); !reflect.DeepEqual(got, testInt.want) {
			t.Errorf("CreateNode() = %v, want %v", got, testInt.want)
		}
	})
}

func TestNewWithElement(t *testing.T) {
	type args[V constraints.Ordered] struct {
		key   V
		value any
	}
	type testCase[V constraints.Ordered] struct {
		name string
		args args[V]
		want *Tree[V]
	}

	intTests := []testCase[int]{
		{
			name: "empty value",
			args: args[int]{key: 1, value: nil},
			want: &Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   1,
						value: nil,
					},
				},
			},
		},
		{
			name: "one element",
			args: args[int]{key: 15, value: 15},
			want: &Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
				},
			},
		},
	}
	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWithElement(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWithElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Insert(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key   V
		value any
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		args args[V]
		want Tree[V]
	}

	treeWithOneElement := Tree[int]{
		root: &node[int]{
			element: Element[int]{
				key:   15,
				value: 15,
			},
			parent: nil,
			left:   nil,
			right: &node[int]{
				element: Element[int]{
					key:   25,
					value: 25,
				},
				parent: nil,
			},
		},
	}
	treeWithOneElement.root.right.parent = treeWithOneElement.root

	right := &node[int]{
		element: Element[int]{
			key:   25,
			value: 25,
		},
		parent: nil,
		right: &node[int]{
			element: Element[int]{
				key:   35,
				value: 35,
			},
		},
	}
	treeWithTwoElements := Tree[int]{
		root: &node[int]{
			element: Element[int]{
				key:   15,
				value: 15,
			},
			parent: nil,
			left:   nil,
			right:  right,
		},
	}

	treeWithTwoElements.root.right.parent = treeWithTwoElements.root
	treeWithTwoElements.root.right.right.parent = right

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    Tree[int]{},
			args: args[int]{key: 15, value: 15},
			want: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
					parent: nil,
					left:   nil,
					right:  nil,
				},
			},
		},
		{
			name: "tree with root and one element",
			t: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args[int]{key: 25, value: 25},
			want: treeWithOneElement,
		},
		{
			name: "tree with root and two elements",
			t:    treeWithOneElement,
			args: args[int]{key: 35, value: 35},
			want: treeWithTwoElements,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tt.t.Insert(tt.args.key, tt.args.value)
			if !reflect.DeepEqual(tt.t, tt.want) {
				t1.Errorf("Insert() = %v, want %v", tt.t, tt.want)
			}
		})
	}
}

func TestTree_Min(t1 *testing.T) {
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		want *Element[V]
	}

	treeWithOneElement := Tree[int]{
		root: &node[int]{
			element: Element[int]{
				key:   15,
				value: 15,
			},
			parent: nil,
			left:   nil,
			right: &node[int]{
				element: Element[int]{
					key:   25,
					value: 25,
				},
				parent: nil,
			},
		},
	}
	treeWithOneElement.root.right.parent = treeWithOneElement.root

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    Tree[int]{root: nil},
			want: nil,
		},
		{
			name: "tree with one element",
			t: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			want: &Element[int]{
				key:   15,
				value: 15,
			},
		},
		{
			name: "tree with root and one element",
			t:    treeWithOneElement,
			want: &Element[int]{
				key:   15,
				value: 15,
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.Min(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Max(t1 *testing.T) {
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		want *Element[V]
	}

	treeWithOneElement := Tree[int]{
		root: &node[int]{
			element: Element[int]{
				key:   15,
				value: 15,
			},
			parent: nil,
			left:   nil,
			right: &node[int]{
				element: Element[int]{
					key:   25,
					value: 25,
				},
				parent: nil,
			},
		},
	}
	treeWithOneElement.root.right.parent = treeWithOneElement.root

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    Tree[int]{root: nil},
			want: nil,
		},
		{
			name: "tree with one element",
			t: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			want: &Element[int]{
				key:   15,
				value: 15,
			},
		},
		{
			name: "tree with root and one element",
			t:    treeWithOneElement,
			want: &Element[int]{
				key:   25,
				value: 25,
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.Max(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Search(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key V
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		args args[V]
		want *Element[V]
	}
	treeWithOneElement := Tree[int]{
		root: &node[int]{
			element: Element[int]{
				key:   15,
				value: 15,
			},
			parent: nil,
			left:   nil,
			right: &node[int]{
				element: Element[int]{
					key:   25,
					value: 25,
				},
				parent: nil,
			},
		},
	}
	treeWithOneElement.root.right.parent = treeWithOneElement.root

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    Tree[int]{root: nil},
			args: args[int]{key: 1},
			want: nil,
		},
		{
			name: "tree with one element - not found",
			t: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args[int]{key: 1},
			want: nil,
		},
		{
			name: "tree with one element - found",
			t: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args[int]{key: 15},
			want: &Element[int]{
				key:   15,
				value: 15,
			},
		},
		{
			name: "tree with root and one element - found",
			t:    treeWithOneElement,
			args: args[int]{key: 25},
			want: &Element[int]{
				key:   25,
				value: 25,
			},
		},
		{
			name: "tree with root and one element - not found",
			t:    treeWithOneElement,
			args: args[int]{key: 35},
			want: nil,
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.Search(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_InOrderTreeWalk(t1 *testing.T) {
	type args struct {
		d direction
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		args args
		want []Element[V]
	}

	treeWithOneElement := Tree[int]{
		root: &node[int]{
			element: Element[int]{
				key:   15,
				value: 15,
			},
			parent: nil,
			left:   nil,
			right: &node[int]{
				element: Element[int]{
					key:   25,
					value: 25,
				},
				parent: nil,
			},
		},
	}
	treeWithOneElement.root.right.parent = treeWithOneElement.root
	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    Tree[int]{root: nil},
			args: args{d: Asc},
			want: nil,
		},
		{
			name: "tree with one element - asc",
			t: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args{d: Asc},
			want: []Element[int]{{key: 15, value: 15}},
		},
		{
			name: "tree with one element - asc",
			t: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args{d: Desc},
			want: []Element[int]{{key: 15, value: 15}},
		},
		{
			name: "tree with root and one element - asc",
			t:    treeWithOneElement,
			args: args{d: Asc},
			want: []Element[int]{
				{key: 15, value: 15},
				{key: 25, value: 25},
			},
		},
		{
			name: "tree with root and one element - desc",
			t:    treeWithOneElement,
			args: args{d: Desc},
			want: []Element[int]{
				{key: 25, value: 25},
				{key: 15, value: 15},
			},
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.InOrderTreeWalk(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("InOrderTreeWalk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_PreOrderSuccessor(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key V
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		args args[V]
		want *Element[V]
	}
	treeWithOneElement := Tree[int]{
		root: &node[int]{
			element: Element[int]{
				key:   15,
				value: 15,
			},
			parent: nil,
			left:   nil,
			right: &node[int]{
				element: Element[int]{
					key:   25,
					value: 25,
				},
				parent: nil,
			},
		},
	}
	treeWithOneElement.root.right.parent = treeWithOneElement.root

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    Tree[int]{root: nil},
			args: args[int]{key: 1},
			want: nil,
		},
		{
			name: "tree with one element - not found",
			t: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args[int]{key: 1},
			want: nil,
		},
		{
			name: "tree with one element",
			t: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args[int]{key: 15},
			want: nil,
		},
		{
			name: "tree with root and one element - found",
			t:    treeWithOneElement,
			args: args[int]{key: 25},
			want: &Element[int]{
				key:   15,
				value: 15,
			},
		},
		{
			name: "tree with root and one element - not found",
			t:    treeWithOneElement,
			args: args[int]{key: 15},
			want: nil,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.PreOrderSuccessor(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("PreOrderSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_PostOrderSuccessor(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key V
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		args args[V]
		want *Element[V]
	}
	treeWithOneElement := Tree[int]{
		root: &node[int]{
			element: Element[int]{
				key:   15,
				value: 15,
			},
			parent: nil,
			left:   nil,
			right: &node[int]{
				element: Element[int]{
					key:   25,
					value: 25,
				},
				parent: nil,
			},
		},
	}
	treeWithOneElement.root.right.parent = treeWithOneElement.root

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    Tree[int]{root: nil},
			args: args[int]{key: 1},
			want: nil,
		},
		{
			name: "tree with one element - not found",
			t: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args[int]{key: 1},
			want: nil,
		},
		{
			name: "tree with one element - found element, but don't found postOrder",
			t: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args[int]{key: 15},
			want: nil,
		},
		{
			name: "tree with root and one element - found",
			t:    treeWithOneElement,
			args: args[int]{key: 15},
			want: &Element[int]{
				key:   25,
				value: 25,
			},
		},
		{
			name: "tree with root and one element - not found",
			t:    treeWithOneElement,
			args: args[int]{key: 25},
			want: nil,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.PostOrderSuccessor(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("PostOrderSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Delete(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key V
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		args args[V]
		want Tree[V]
	}

	treeWithOneElement := Tree[int]{
		root: &node[int]{
			element: Element[int]{
				key:   15,
				value: 15,
			},
			parent: nil,
			left:   nil,
			right: &node[int]{
				element: Element[int]{
					key:   25,
					value: 25,
				},
				parent: nil,
			},
		},
	}
	treeWithOneElement.root.right.parent = treeWithOneElement.root

	treeResult := Tree[int]{
		root: &node[int]{
			element: Element[int]{
				key:   15,
				value: 15,
			},
			parent: nil,
			left:   nil,
			right: &node[int]{
				element: Element[int]{
					key:   35,
					value: 35,
				},
				parent: nil,
			},
		},
	}
	treeResult.root.right.parent = treeResult.root

	right := &node[int]{
		element: Element[int]{
			key:   25,
			value: 25,
		},
		parent: nil,
		right: &node[int]{
			element: Element[int]{
				key:   35,
				value: 35,
			},
		},
	}
	treeWithTwoElements := Tree[int]{
		root: &node[int]{
			element: Element[int]{
				key:   15,
				value: 15,
			},
			parent: nil,
			left:   nil,
			right:  right,
		},
	}

	treeWithTwoElements.root.right.parent = treeWithTwoElements.root
	treeWithTwoElements.root.right.right.parent = right

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    Tree[int]{},
			args: args[int]{key: 1},
			want: Tree[int]{},
		},
		{
			name: "tree only with root - without changes",
			t: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
					parent: nil,
					left:   nil,
					right:  nil,
				},
			},
			args: args[int]{key: 1},
			want: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
					parent: nil,
					left:   nil,
					right:  nil,
				},
			},
		},
		{
			name: "tree only with root - delete root",
			t: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
					parent: nil,
					left:   nil,
					right:  nil,
				},
			},
			args: args[int]{key: 15},
			want: Tree[int]{},
		},
		{
			name: "tree with elements - without changes",
			t:    treeWithOneElement,
			args: args[int]{key: 85},
			want: treeWithOneElement,
		},
		{
			name: "tree with elements - delete node without children",
			t:    treeWithOneElement,
			args: args[int]{key: 25},
			want: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   15,
						value: 15,
					},
					parent: nil,
					left:   nil,
					right:  nil,
				},
			},
		},
		{
			name: "tree with elements - delete node with right children",
			t:    treeWithTwoElements,
			args: args[int]{key: 25},
			want: treeResult,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tt.t.Delete(tt.args.key)
			if !reflect.DeepEqual(tt.t, tt.want) {
				t1.Errorf("Delete() = %v, want %v", tt.t, tt.want)
			}
		})
	}
}

func TestTree_DeleteThirdCase(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key V
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		args args[V]
		want Tree[V]
	}

	tree := Tree[int]{
		root: &node[int]{
			element: Element[int]{
				key:   15,
				value: 15,
			},
			parent: nil,
			left: &node[int]{
				element: Element[int]{
					key:   10,
					value: 10,
				},
				parent: nil,
			},
			right: &node[int]{
				element: Element[int]{
					key:   25,
					value: 25,
				},
				parent: nil,
			},
		},
	}
	tree.root.right.parent = tree.root
	tree.root.left.parent = tree.root

	treeResult := Tree[int]{
		root: &node[int]{
			element: Element[int]{
				key:   25,
				value: 25,
			},
			left: &node[int]{
				element: Element[int]{
					key:   10,
					value: 10,
				},
				parent: nil,
			},
		},
	}
	treeResult.root.left.parent = treeResult.root

	tests := []testCase[int]{
		{
			name: "third case",
			t:    tree,
			args: args[int]{key: 15},
			want: treeResult,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tt.t.Delete(tt.args.key)
			if !reflect.DeepEqual(tt.t, tt.want) {
				t1.Errorf("Delete() = %v, want %v", tt.t, tt.want)
			}
		})
	}
}

func TestTree_DeleteSecondCase(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key V
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		args args[V]
		want Tree[V]
	}

	treeWithOneElement := Tree[int]{
		root: &node[int]{
			element: Element[int]{
				key:   15,
				value: 15,
			},
			parent: nil,
			left:   nil,
			right: &node[int]{
				element: Element[int]{
					key:   25,
					value: 25,
				},
				parent: nil,
			},
		},
	}
	treeWithOneElement.root.right.parent = treeWithOneElement.root

	tests := []testCase[int]{
		{
			name: "tree with elements - delete root node with right children",
			t:    treeWithOneElement,
			args: args[int]{key: 15},
			want: Tree[int]{
				root: &node[int]{
					element: Element[int]{
						key:   25,
						value: 25,
					},
					parent: nil,
					left:   nil,
					right:  nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tt.t.Delete(tt.args.key)
			if !reflect.DeepEqual(tt.t, tt.want) {
				t1.Errorf("Delete() = %v, want %v", tt.t, tt.want)
			}
		})
	}
}
