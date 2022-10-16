package go_heaps

import (
	"github.com/void-developer/go-heaps/src/heaps"
	"github.com/void-developer/go-heaps/src/heaps/types"
)

const (
	MinQ int = 1
	MaxQ int = 0
)

// Heap is the base data structure for the binary heap. The heap has the basic tree and size properties.
// Additional properties are:
//   - Type: the type of the heap. Either MinQ or MaxQ.
//   - Nil: the value that represents a null value in the heap.
type Heap[C types.Comparable] struct {
	Tree []C
	Size int
	Type int
}

// comparePriority uses the Compare method of the Heap type to compare the priorities of the two values.
// It compares val1 to val2 if the heap is of type MaxQ, and it compares val2 to val1 if it's of type MinQ
func (h *Heap[C]) comparePriority(val1 C, val2 C) int {
	if h.Type == MaxQ {
		return val1.Compare(val2)
	} else if h.Type == MinQ {
		return val2.Compare(val1)
	} else {
		return 0
	}
}

// Peek returns the value of the top of the heap without removing it
func (h *Heap[C]) Peek() C {
	return h.Tree[0]
}

// Push pushes an element to the rightmost child of the last element of the binary tree and then rebalances the tree
// using the siftup operation
func (h *Heap[C]) Push(val C) {
	h.Tree[h.Size] = val
	h.Size += 1
	if h.Size > 1 {
		h.siftUp()
	}
}

// siftUp moves the element at index ix up the tree until it's it's parent has a higher priority than it
func (h *Heap[C]) siftUp() {
	currentNode := h.Size - 1
	for h.comparePriority(h.Tree[currentNode], h.Tree[heaps.GetParent(currentNode)]) > 0 {
		parentNode := heaps.GetParent(currentNode)
		temp := h.Tree[parentNode]
		h.Tree[parentNode] = h.Tree[currentNode]
		h.Tree[currentNode] = temp
		currentNode = parentNode
	}
}

// siftDown moves the element at index ix down the tree until its children have a lower priority than it
func (h *Heap[C]) siftDown(ix int) {
	currentNode := ix
	children := heaps.GetChildren(currentNode)
	for children[0] < h.Size && !h.Tree[children[0]].IsNull() {
		swap := -1
		if h.comparePriority(h.Tree[children[0]], h.Tree[currentNode]) > 0 {
			swap = children[0]
		}

		if children[1] < h.Size && (swap == -1 || h.comparePriority(h.Tree[children[1]], h.Tree[swap]) > 0) && h.comparePriority(h.Tree[children[1]], h.Tree[currentNode]) > 0 {
			swap = children[1]
		}

		if swap == -1 {
			break
		}

		temp := h.Tree[swap]
		h.Tree[swap] = h.Tree[currentNode]
		h.Tree[currentNode] = temp
		currentNode = swap
		children = heaps.GetChildren(currentNode)
	}
}

// Replace replaces the value at the root of the heap, returns the old value and rebalances the tree performing
// a siftdown operation
func (h *Heap[C]) Replace(val C) C {
	if !h.IsFull() {
		h.Push(val)
		return *new(C)
	}
	popped := h.Tree[0]
	h.Tree[0] = val
	h.siftDown(0)
	return popped
}

// Pop removes the top of the heap and rebalances the tree using the siftdown operation
func (h *Heap[C]) Pop() C {
	if h.Size == 0 {
		return *new(C)
	}

	popped := h.Tree[0]
	h.Size -= 1
	if h.Size > 0 {
		h.Tree[0] = h.Tree[h.Size]
		//h.Tree[h.Size] = h.Nil
		h.siftDown(0)
	}
	return popped
}

// IsFull returns true if the heap is full
func (h *Heap[C]) IsFull() bool {
	return h.Size == len(h.Tree)
}

// Init initializes the heap with the given type and nil value
func (h *Heap[C]) Init(size int, _type int) {
	h.Tree = make([]C, size)
	//for i := 0; i < size; i++ {
	//	h.Tree[i] = h.Nil
	//}
	h.Size = 0
	h.Type = _type
}

// InitFromArray initializes the heap with the given type and nil value and then populates the heap with the given array
func (h *Heap[C]) InitFromArray(arr []C) {
	h.Tree = arr
	h.Size = len(arr)
}

// Heapify creates a heap out of the given array. The heapification is an in-place operation
func (h *Heap[C]) Heapify(arr []C) *Heap[C] {
	h.Tree = arr
	h.Size = len(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}
