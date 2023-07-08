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
		t    Tree[V]
		args args
		want []V
	}

	intTests := []testCase[int]{
		{
			name: "empty",
			t:    Tree[int]{},
			args: args{d: Asc},
			want: []int{0},
		},
		{
			name: "one element",
			t:    Tree[int]{value: 5},
			args: args{d: Asc},
			want: []int{5},
		},
		{
			name: "ascending",
			t: Tree[int]{
				value: 5,
				left:  &Tree[int]{value: 1},
				right: &Tree[int]{value: 10},
			},
			args: args{d: Asc},
			want: []int{1, 5, 10},
		},
		{
			name: "descending",
			t: Tree[int]{
				value: 5,
				left:  &Tree[int]{value: 1},
				right: &Tree[int]{value: 10},
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
		t    Tree[V]
		args args[V]
		want *Tree[V]
	}

	intTests := []testCase[int]{
		{
			name: "empty",
			t:    Tree[int]{},
			args: args[int]{val: 0},
			want: &Tree[int]{value: 0},
		},
		{
			name: "not found",
			t:    Tree[int]{},
			args: args[int]{val: 10},
			want: nil,
		},
		{
			name: "found left",
			t: Tree[int]{
				value: 5,
				left:  &Tree[int]{value: 1},
				right: &Tree[int]{value: 10},
			},
			args: args[int]{val: 1},
			want: &Tree[int]{value: 1},
		},
		{
			name: "found right",
			t: Tree[int]{
				value: 5,
				left:  &Tree[int]{value: 1},
				right: &Tree[int]{value: 10},
			},
			args: args[int]{val: 10},
			want: &Tree[int]{value: 10},
		},
		{
			name: "found root",
			t: Tree[int]{
				value: 5,
				left:  &Tree[int]{value: 1},
				right: &Tree[int]{value: 10},
			},
			args: args[int]{val: 5},
			want: &Tree[int]{
				value: 5,
				left:  &Tree[int]{value: 1},
				right: &Tree[int]{value: 10},
			},
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
		t    Tree[V]
		want *Tree[V]
	}

	intTests := []testCase[int]{
		{
			name: "empty",
			t:    Tree[int]{},
			want: &Tree[int]{value: 0},
		},
		{
			name: "one element",
			t:    Tree[int]{value: 15},
			want: &Tree[int]{value: 15},
		},
		{
			name: "found min",
			t: Tree[int]{
				value: 5,
				left:  &Tree[int]{value: 1},
				right: &Tree[int]{value: 10},
			},
			want: &Tree[int]{value: 1},
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
		t    Tree[V]
		want *Tree[V]
	}
	intTests := []testCase[int]{
		{
			name: "empty",
			t:    Tree[int]{},
			want: &Tree[int]{value: 0},
		},
		{
			name: "one element",
			t:    Tree[int]{value: 15},
			want: &Tree[int]{value: 15},
		},
		{
			name: "found max",
			t: Tree[int]{
				value: 5,
				left:  &Tree[int]{value: 1},
				right: &Tree[int]{value: 10},
			},
			want: &Tree[int]{value: 10},
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
		el *Tree[V]
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		args args[V]
		want *Tree[V]
	}

	intTests := []testCase[int]{
		{
			name: "empty",
			t:    Tree[int]{},
			args: args[int]{el: &Tree[int]{value: 0}},
			want: &Tree[int]{value: 0},
		},
		{
			name: "one element",
			t:    Tree[int]{value: 15},
			args: args[int]{el: &Tree[int]{value: 15}},
			want: &Tree[int]{value: 0},
		},
		{
			name: "found preOrder",
			t: Tree[int]{
				value: 5,
				left:  &Tree[int]{value: 1},
				right: &Tree[int]{value: 10},
			},
			args: args[int]{el: &Tree[int]{
				value: 5,
				left:  &Tree[int]{value: 1},
				right: &Tree[int]{value: 10},
			}},
			want: &Tree[int]{value: 1},
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
		el *Tree[V]
	}
	type testCase[V constraints.Ordered] struct {
		name string
		t    Tree[V]
		args args[V]
		want *Tree[V]
	}
	intTests := []testCase[int]{
		{
			name: "empty",
			t:    Tree[int]{},
			args: args[int]{el: &Tree[int]{value: 0}},
			want: &Tree[int]{value: 0},
		},
		{
			name: "one element",
			t:    Tree[int]{value: 15},
			args: args[int]{el: &Tree[int]{value: 15}},
			want: &Tree[int]{value: 0},
		},
		{
			name: "found postOrder",
			t: Tree[int]{
				value: 5,
				left:  &Tree[int]{value: 1},
				right: &Tree[int]{value: 10},
			},
			args: args[int]{el: &Tree[int]{
				value: 5,
				left:  &Tree[int]{value: 1},
				right: &Tree[int]{value: 10},
			}},
			want: &Tree[int]{value: 10},
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
		t    Tree[V]
		args args[V]
	}
	intTests := []testCase[int]{
		{
			name: "empty",
			t:    Tree[int]{},
			args: args[int]{val: 0},
		},
		{
			name: "one element",
			t:    Tree[int]{value: 10},
			args: args[int]{val: 23},
		},
		{
			name: "several elements",
			t: Tree[int]{
				value: 5,
				left:  &Tree[int]{value: 1},
				right: &Tree[int]{value: 10},
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
		want *Tree[V]
	}

	testInt := testCase[int]{
		name: "int node",
		args: args[int]{val: 10},
		want: &Tree[int]{value: 10},
	}
	t.Run(testInt.name, func(t *testing.T) {
		if got := CreateNode(testInt.args.val); !reflect.DeepEqual(got, testInt.want) {
			t.Errorf("CreateNode() = %v, want %v", got, testInt.want)
		}
	})

	testString := testCase[string]{
		name: "string node",
		args: args[string]{val: "test"},
		want: &Tree[string]{value: "test"},
	}
	t.Run(testString.name, func(t *testing.T) {
		if got := CreateNode(testString.args.val); !reflect.DeepEqual(got, testString.want) {
			t.Errorf("CreateNode() = %v, want %v", got, testString.want)
		}
	})
}
