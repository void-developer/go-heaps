package go_heaps

import (
	"github.com/void-developer/go-heaps/src/heaps/types"
	"testing"
)

type IntHeap struct {
	Val int
}

func (h IntHeap) Compare(other types.Comparable) int {
	return h.Val - other.(IntHeap).Val
}

func (h IntHeap) IsNull() bool {
	return h.Val == -1
}

func (h IntHeap) Equals(other types.Comparable) bool {
	return h.Val == other.(IntHeap).Val
}

func TestHeapifyArray(t *testing.T) {
	intArray := []IntHeap{{2}, {4}, {1}, {10}, {7}}
	heap := Heap[IntHeap]{}
	heap.Heapify(intArray)
	if heap.Pop().Val != 10 || heap.Pop().Val != 7 || heap.Pop().Val != 4 || heap.Pop().Val != 2 || heap.Pop().Val != 1 {
		t.Fatalf("The heapification returned a non valid heap")
	}
}
