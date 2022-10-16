package go_heaps

import (
	"github.com/void-developer/go-heaps/src/heaps"
	"github.com/void-developer/go-heaps/src/heaps/types"
)

const (
	MinQ int = 1
	MaxQ int = 0
)

type Heap[C types.Comparable] struct {
	Tree []C
	Size int
	Type int
}

func (h *Heap[C]) comparePriority(val1 C, val2 C) int {
	if h.Type == MaxQ {
		return val1.Compare(val2)
	} else if h.Type == MinQ {
		return val2.Compare(val1)
	} else {
		return 0
	}
}

func (h *Heap[C]) Peek() C {
	return h.Tree[0]
}

func (h *Heap[C]) Push(val C) {
	h.Tree[h.Size] = val
	h.Size += 1
	if h.Size > 1 {
		h.siftUp()
	}
}

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

func (h *Heap[C]) Replace(val C) C {
	if !h.IsFull() {
		h.Push(val)
	}
	popped := h.Tree[0]
	h.Tree[0] = val
	h.siftDown(0)
	return popped
}

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

func (h *Heap[C]) IsFull() bool {
	return h.Size == len(h.Tree)
}

func (h *Heap[C]) Init(size int, _type int) {
	h.Tree = make([]C, size)
	//for i := 0; i < size; i++ {
	//	h.Tree[i] = h.Nil
	//}
	h.Size = 0
	h.Type = _type
}

func (h *Heap[C]) InitFromArray(arr []C) {
	h.Tree = arr
	h.Size = len(arr)
}

func (h *Heap[C]) Heapify(arr []C) *Heap[C] {
	h.Tree = arr
	h.Size = len(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		h.siftDown(i)
	}
	return h
}
