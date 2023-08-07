package tree

import "testing"

func BenchmarkTreeInsert(b *testing.B) {
	tree := &Tree[int]{}

	for i := 0; i < b.N; i++ {
		tree.Insert(i, i)
	}
}

func BenchmarkTree_InsertWithoutRecursion(b *testing.B) {
	tree := &Tree[int]{}

	for i := 0; i < b.N; i++ {
		tree.InsertWithoutRecursion(i, i)
	}
}
