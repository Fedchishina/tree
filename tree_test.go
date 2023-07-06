package main

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
			if got := tt.t.preOrderSuccessor(tt.args.el); !reflect.DeepEqual(got, tt.want) {
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
			if got := tt.t.postOrderSuccessor(tt.args.el); !reflect.DeepEqual(got, tt.want) {
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
