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
					element: element[int]{
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
					element: element[int]{
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

	treeWithOneElement := New[int]()
	treeWithOneElement.Insert(15, 15)
	treeWithOneElement.Insert(25, 25)

	treeWithTwoElements := New[int]()
	treeWithTwoElements.Insert(15, 15)
	treeWithTwoElements.Insert(25, 25)
	treeWithTwoElements.Insert(35, 35)

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    Tree[int]{},
			args: args[int]{key: 15, value: 15},
			want: Tree[int]{
				root: &node[int]{
					element: element[int]{
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
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args[int]{key: 25, value: 25},
			want: *treeWithOneElement,
		},
		{
			name: "tree with root and two elements",
			t:    *treeWithOneElement,
			args: args[int]{key: 35, value: 35},
			want: *treeWithTwoElements,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tt.t.Insert(tt.args.key, tt.args.value)
			if !reflect.DeepEqual(tt.t, tt.want) {
				t1.Errorf("Insert() = %#+v, want %#+v", tt.t, tt.want)
			}
		})
	}
}

func TestTree_Min(t1 *testing.T) {
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		want V
	}

	treeWithOneElement := New[int]()
	treeWithOneElement.Insert(15, 15)
	treeWithOneElement.Insert(25, 25)

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    Tree[int]{root: nil},
			want: 0,
		},
		{
			name: "tree with one element",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			want: 15,
		},
		{
			name: "tree with root and one element",
			t:    *treeWithOneElement,
			want: 15,
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
		want V
	}

	treeWithOneElement := New[int]()
	treeWithOneElement.Insert(15, 15)
	treeWithOneElement.Insert(25, 25)

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    Tree[int]{root: nil},
			want: 0,
		},
		{
			name: "tree with one element",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			want: 15,
		},
		{
			name: "tree with root and one element",
			t:    *treeWithOneElement,
			want: 25,
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

func TestTree_Exist(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key V
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		args args[V]
		want bool
	}
	treeWithOneElement := New[int]()
	treeWithOneElement.Insert(15, 15)
	treeWithOneElement.Insert(25, 25)

	tests := []testCase[int]{
		{
			name: "empty tree",
			t:    Tree[int]{root: nil},
			args: args[int]{key: 1},
			want: false,
		},
		{
			name: "tree with one element - not found",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args[int]{key: 1},
			want: false,
		},
		{
			name: "tree with one element - found",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args[int]{key: 15},
			want: true,
		},
		{
			name: "tree with root and one element - found",
			t:    *treeWithOneElement,
			args: args[int]{key: 25},
			want: true,
		},
		{
			name: "tree with root and one element - not found",
			t:    *treeWithOneElement,
			args: args[int]{key: 35},
			want: false,
		},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.Exists(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Exists() = %v, want %v", got, tt.want)
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
		want []V
	}

	treeWithOneElement := New[int]()
	treeWithOneElement.Insert(15, 15)
	treeWithOneElement.Insert(25, 25)

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
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args{d: Asc},
			want: []int{15},
		},
		{
			name: "tree with one element - asc",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args: args{d: Desc},
			want: []int{15},
		},
		{
			name: "tree with root and one element - asc",
			t:    *treeWithOneElement,
			args: args{d: Asc},
			want: []int{15, 25},
		},
		{
			name: "tree with root and one element - desc",
			t:    *treeWithOneElement,
			args: args{d: Desc},
			want: []int{25, 15},
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

	treeWithOneElement := New[int]()
	treeWithOneElement.Insert(15, 15)
	treeWithOneElement.Insert(25, 25)

	treeWithTwoElements := New[int]()
	treeWithTwoElements.Insert(15, 15)
	treeWithTwoElements.Insert(25, 25)
	treeWithTwoElements.Insert(35, 35)

	treeResult := New[int]()
	treeResult.Insert(15, 15)
	treeResult.Insert(35, 35)

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
					element: element[int]{
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
					element: element[int]{
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
					element: element[int]{
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
			t:    *treeWithOneElement,
			args: args[int]{key: 85},
			want: *treeWithOneElement,
		},
		{
			name: "tree with elements - delete node without children",
			t:    *treeWithOneElement,
			args: args[int]{key: 25},
			want: Tree[int]{
				root: &node[int]{
					element: element[int]{
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
			t:    *treeWithTwoElements,
			args: args[int]{key: 25},
			want: *treeResult,
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
			element: element[int]{
				key:   15,
				value: 15,
			},
			parent: nil,
			left: &node[int]{
				element: element[int]{
					key:   10,
					value: 10,
				},
				parent: nil,
			},
			right: &node[int]{
				element: element[int]{
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
			element: element[int]{
				key:   25,
				value: 25,
			},
			left: &node[int]{
				element: element[int]{
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

	treeWithOneElement := New[int]()
	treeWithOneElement.Insert(15, 15)
	treeWithOneElement.Insert(25, 25)

	tests := []testCase[int]{
		{
			name: "tree with elements - delete root node with right children",
			t:    *treeWithOneElement,
			args: args[int]{key: 15},
			want: Tree[int]{
				root: &node[int]{
					element: element[int]{
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

func TestTree_PreOrderSuccessor(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key V
	}
	type testCase[V constraints.Ordered] struct {
		name    string
		t       Tree[V]
		args    args[V]
		want    V
		wantErr bool
	}
	treeWithOneElement := New[int]()
	treeWithOneElement.Insert(15, 15)
	treeWithOneElement.Insert(25, 25)

	tests := []testCase[int]{
		{
			name:    "empty tree",
			t:       Tree[int]{root: nil},
			args:    args[int]{key: 1},
			want:    0,
			wantErr: true,
		},
		{
			name: "tree with one element - not found",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args:    args[int]{key: 1},
			want:    0,
			wantErr: true,
		},
		{
			name: "tree with one element",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args:    args[int]{key: 15},
			want:    0,
			wantErr: true,
		},
		{
			name:    "tree with root and one element - found",
			t:       *treeWithOneElement,
			args:    args[int]{key: 25},
			want:    15,
			wantErr: false,
		},
		{
			name:    "tree with root and one element - not found",
			t:       *treeWithOneElement,
			args:    args[int]{key: 15},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got, err := tt.t.PreOrderSuccessor(tt.args.key)
			if (err != nil) != tt.wantErr {
				t1.Errorf("PreOrderSuccessor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("PreOrderSuccessor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_PostOrderSuccessor(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key V
	}
	type testCase[V constraints.Ordered] struct {
		name    string
		t       Tree[V]
		args    args[V]
		want    V
		wantErr bool
	}
	treeWithOneElement := New[int]()
	treeWithOneElement.Insert(15, 15)
	treeWithOneElement.Insert(25, 25)

	tests := []testCase[int]{
		{
			name:    "empty tree",
			t:       Tree[int]{root: nil},
			args:    args[int]{key: 1},
			want:    0,
			wantErr: true,
		},
		{
			name: "tree with one element - not found",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args:    args[int]{key: 1},
			want:    0,
			wantErr: true,
		},
		{
			name: "tree with one element - found element, but don't found postOrder",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args:    args[int]{key: 15},
			want:    0,
			wantErr: true,
		},
		{
			name:    "tree with root and one element - found",
			t:       *treeWithOneElement,
			args:    args[int]{key: 15},
			want:    25,
			wantErr: false,
		},
		{
			name:    "tree with root and one element - not found",
			t:       *treeWithOneElement,
			args:    args[int]{key: 25},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got, err := tt.t.PostOrderSuccessor(tt.args.key)
			if (err != nil) != tt.wantErr {
				t1.Errorf("PostOrderSuccessor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("PostOrderSuccessor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_GetValue(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		key V
	}
	type testCase[V constraints.Ordered] struct {
		name    string
		t       Tree[V]
		args    args[V]
		want    any
		wantErr bool
	}
	treeWithOneElement := New[int]()
	treeWithOneElement.Insert(15, 15)
	treeWithOneElement.Insert(25, 25)

	tests := []testCase[int]{
		{
			name:    "empty tree",
			t:       Tree[int]{root: nil},
			args:    args[int]{key: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "tree with one element - not found",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args:    args[int]{key: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "tree with one element - found",
			t: Tree[int]{
				root: &node[int]{
					element: element[int]{
						key:   15,
						value: 15,
					},
				},
			},
			args:    args[int]{key: 15},
			want:    15,
			wantErr: false,
		},
		{
			name:    "tree with root and one element - found",
			t:       *treeWithOneElement,
			args:    args[int]{key: 25},
			want:    25,
			wantErr: false,
		},
		{
			name:    "tree with root and one element - not found",
			t:       *treeWithOneElement,
			args:    args[int]{key: 35},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got, err := tt.t.GetValue(tt.args.key)
			if (err != nil) != tt.wantErr {
				t1.Errorf("GetValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("GetValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}
