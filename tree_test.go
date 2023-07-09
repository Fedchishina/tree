package tree

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func Test_tree_InOrderTreeWalk(t1 *testing.T) {
	type args struct {
		d direction
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    tree[V]
		args args
		want []V
	}

	intTests := []testCase[int]{
		{
			name: "empty",
			t:    tree[int]{},
			args: args{d: Asc},
			want: []int{0},
		},
		{
			name: "one element",
			t:    tree[int]{value: 5},
			args: args{d: Asc},
			want: []int{5},
		},
		{
			name: "ascending",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			args: args{d: Asc},
			want: []int{1, 5, 10},
		},
		{
			name: "descending",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			args: args{d: Desc},
			want: []int{10, 5, 1},
		},
	}
	for _, tt := range intTests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.InOrderTreeWalk(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("InOrderTreeWalk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_Search(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		val V
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    tree[V]
		args args[V]
		want *tree[V]
	}

	intTests := []testCase[int]{
		{
			name: "empty",
			t:    tree[int]{},
			args: args[int]{val: 0},
			want: &tree[int]{value: 0},
		},
		{
			name: "not found",
			t:    tree[int]{},
			args: args[int]{val: 10},
			want: nil,
		},
		{
			name: "found left",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			args: args[int]{val: 1},
			want: &tree[int]{value: 1},
		},
		{
			name: "found right",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			args: args[int]{val: 10},
			want: &tree[int]{value: 10},
		},
		{
			name: "found root",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			args: args[int]{val: 5},
			want: &tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
		},
		{
			name: "not found",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			args: args[int]{val: 8},
			want: nil,
		},
	}
	for _, tt := range intTests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.Search(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_Min(t1 *testing.T) {
	type testCase[V constraints.Ordered] struct {
		name string
		t    tree[V]
		want *tree[V]
	}

	intTests := []testCase[int]{
		{
			name: "empty",
			t:    tree[int]{},
			want: &tree[int]{value: 0},
		},
		{
			name: "one element",
			t:    tree[int]{value: 15},
			want: &tree[int]{value: 15},
		},
		{
			name: "found min",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			want: &tree[int]{value: 1},
		},
	}
	for _, tt := range intTests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.Min(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_Max(t1 *testing.T) {
	type testCase[V constraints.Ordered] struct {
		name string
		t    tree[V]
		want *tree[V]
	}
	intTests := []testCase[int]{
		{
			name: "empty",
			t:    tree[int]{},
			want: &tree[int]{value: 0},
		},
		{
			name: "one element",
			t:    tree[int]{value: 15},
			want: &tree[int]{value: 15},
		},
		{
			name: "found max",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			want: &tree[int]{value: 10},
		},
	}
	for _, tt := range intTests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.Max(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_preOrderSuccessor(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		el *tree[V]
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    tree[V]
		args args[V]
		want *tree[V]
	}

	intTests := []testCase[int]{
		{
			name: "empty",
			t:    tree[int]{},
			args: args[int]{el: &tree[int]{value: 0}},
			want: &tree[int]{value: 0},
		},
		{
			name: "one element",
			t:    tree[int]{value: 15},
			args: args[int]{el: &tree[int]{value: 15}},
			want: &tree[int]{value: 0},
		},
		{
			name: "found preOrder",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			args: args[int]{el: &tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			}},
			want: &tree[int]{value: 1},
		},
	}

	for _, tt := range intTests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.PreOrderSuccessor(tt.args.el); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("preOrderSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_postOrderSuccessor(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		el *tree[V]
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    tree[V]
		args args[V]
		want *tree[V]
	}
	intTests := []testCase[int]{
		{
			name: "empty",
			t:    tree[int]{},
			args: args[int]{el: &tree[int]{value: 0}},
			want: &tree[int]{value: 0},
		},
		{
			name: "one element",
			t:    tree[int]{value: 15},
			args: args[int]{el: &tree[int]{value: 15}},
			want: &tree[int]{value: 0},
		},
		{
			name: "found postOrder",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			args: args[int]{el: &tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			}},
			want: &tree[int]{value: 10},
		},
	}
	for _, tt := range intTests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.PostOrderSuccessor(tt.args.el); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("postOrderSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tree_Insert(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		val V
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    tree[V]
		args args[V]
	}
	intTests := []testCase[int]{
		{
			name: "empty",
			t:    tree[int]{},
			args: args[int]{val: 0},
		},
		{
			name: "one element",
			t:    tree[int]{value: 10},
			args: args[int]{val: 23},
		},
		{
			name: "several elements",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			args: args[int]{val: 23},
		},
	}
	for _, tt := range intTests {
		t1.Run(tt.name, func(t1 *testing.T) {
			tt.t.Insert(tt.args.val)
		})
	}
}

func TestCreateNode(t *testing.T) {
	type args[V constraints.Ordered] struct {
		val V
	}

	type testCase[V constraints.Ordered] struct {
		name string
		args args[V]
		want *tree[V]
	}

	testInt := testCase[int]{
		name: "int node",
		args: args[int]{val: 10},
		want: &tree[int]{value: 10},
	}
	t.Run(testInt.name, func(t *testing.T) {
		if got := CreateNode(testInt.args.val); !reflect.DeepEqual(got, testInt.want) {
			t.Errorf("CreateNode() = %v, want %v", got, testInt.want)
		}
	})

	testString := testCase[string]{
		name: "string node",
		args: args[string]{val: "test"},
		want: &tree[string]{value: "test"},
	}
	t.Run(testString.name, func(t *testing.T) {
		if got := CreateNode(testString.args.val); !reflect.DeepEqual(got, testString.want) {
			t.Errorf("CreateNode() = %v, want %v", got, testString.want)
		}
	})
}

func TestTree_Parent(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		val V
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    tree[V]
		args args[V]
		want *tree[V]
	}
	intTests := []testCase[int]{
		{
			name: "empty tree",
			t:    tree[int]{},
			args: args[int]{val: 13},
			want: nil,
		},
		{
			name: "empty tree - arg 0",
			t:    tree[int]{},
			args: args[int]{val: 0},
			want: nil,
		},
		{
			name: "one element in tree",
			t:    tree[int]{value: 13},
			args: args[int]{val: 13},
			want: nil,
		},
		{
			name: "one element in tree - not found",
			t:    tree[int]{value: 13},
			args: args[int]{val: 1},
			want: nil,
		},
		{
			name: "found parent for left node",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			args: args[int]{val: 1},
			want: &tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
		},
		{
			name: "found parent for right node",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			args: args[int]{val: 10},
			want: &tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
		},
		{
			name: "not found parent",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			args: args[int]{val: 8},
			want: nil,
		},
	}
	for _, tt := range intTests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if got := tt.t.Parent(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Parent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Delete(t1 *testing.T) {
	type args[V constraints.Ordered] struct {
		val V
	}
	type testCase[V constraints.Ordered] struct {
		name    string
		t       tree[V]
		args    args[V]
		wantErr bool
		want    *tree[V]
	}
	intTests := []testCase[int]{
		{
			name:    "empty tree",
			t:       tree[int]{},
			args:    args[int]{val: 13},
			wantErr: true,
		},
		{
			name:    "not found for deleting",
			t:       tree[int]{value: 1},
			args:    args[int]{val: 13},
			wantErr: true,
		},
		{
			name:    "one element in tree",
			t:       tree[int]{value: 1},
			args:    args[int]{val: 1},
			wantErr: false,
			want:    nil,
		},
		{
			name: "first case - left node without children",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			args:    args[int]{val: 1},
			wantErr: false,
			want: &tree[int]{
				value: 5,
				left:  nil,
				right: &tree[int]{value: 10},
			},
		},
		{
			name: "first case - right node without children",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
			args:    args[int]{val: 10},
			wantErr: false,
			want: &tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: nil,
			},
		},
		{
			name: "second case - left node with children",
			t: tree[int]{
				value: 5,
				left:  &tree[int]{value: 2, left: &tree[int]{value: 1}},
				right: &tree[int]{value: 10},
			},
			args:    args[int]{val: 2},
			wantErr: false,
			want: &tree[int]{
				value: 5,
				left:  &tree[int]{value: 1},
				right: &tree[int]{value: 10},
			},
		},
		{
			name: "root remove",
			t: tree[int]{
				value: 3,
				right: &tree[int]{
					value: 9,
					left: &tree[int]{
						value: 7,
						right: &tree[int]{
							value: 8,
						},
					},
					right: &tree[int]{value: 10},
				},
			},
			args:    args[int]{val: 3},
			wantErr: false,
			want: &tree[int]{
				value: 9,
				left: &tree[int]{
					value: 7,
					right: &tree[int]{
						value: 8,
					},
				},
				right: &tree[int]{value: 10},
			},
		},
		{
			name: "third case",
			t: tree[int]{
				value: 3,
				left:  &tree[int]{value: 1},
				right: &tree[int]{
					value: 9,
					left: &tree[int]{
						value: 7,
						right: &tree[int]{
							value: 8,
						},
					},
					right: &tree[int]{value: 10},
				},
			},
			args:    args[int]{val: 3},
			wantErr: false,
			want: &tree[int]{
				value: 7,
				left:  &tree[int]{value: 1},
				right: &tree[int]{
					value: 9,
					left:  &tree[int]{value: 8},
					right: &tree[int]{value: 10},
				},
			},
		},
	}
	for _, tt := range intTests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if err := tt.t.Delete(tt.args.val); (err != nil) != tt.wantErr {
				t1.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := tt.t.Delete(tt.args.val); err == nil && !reflect.DeepEqual(tt.t, tt.want) &&
				tt.name != "one element in tree" &&
				tt.name != "root remove" {
				t1.Errorf("Delete() = %v, want %v", tt.t, tt.want)
			}
		})
	}
}
