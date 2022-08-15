# Go Heaps

This package offers the Heap structure implemented with any type that extends `Comparable` (_/heaps/types_).

## Heap Container Structure

In order to use a structure inside the heap, this structure must be comparable with other similar items. So implementing the `Comparable` interface, the heap structure, must have a `Equals`, `Compare`, and a `isNull` method.
An example would be a simple Int wrapper:

```go
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
```
This particular wrapper considers NULL the numbers with -1 (but that is the user's choice)

## Creating a Heap

To create a Heap it's necessary to specify the type of element that the Heap contains/will contain at initialization.

```go
heap := go_heaps.Heap[IntHeap]{}
```

### Using the Heap

The heap structure exposes the main heap operations which are:
* **Push** -> adds an element to the heap and performs a sift up operation to rebalance the heap
* **Pop** -> removes the element with the highest priority and performs a sift down operation to rebalance the heap
* **Peek** -> returns the element with the highest priority without removing it from the heap
* **Size** -> returns the number of elements in the heap
* **Replace** -> replaces the element with the highest priority with the given element and performs a sift down operation to rebalance the heap (if the heap is not full it simply pushes the element)

