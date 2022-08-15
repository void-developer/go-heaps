package heaps

//func HeapifyByType(arr []int, _type int) *HeapO {
//	h := &HeapO{
//		Tree: arr,
//		Size: len(arr),
//		Type: _type,
//	}
//	for i := len(arr) - 1; i >= 0; i-- {
//		h.siftDown(i)
//	}
//	return h
//}
//
//func Heapify(arr []int) *HeapO {
//	return HeapifyByType(arr, MaxQ)
//}

func GetParent(ix int) int {
	if ix == 0 {
		return ix
	}
	return (ix - 1) / 2
}

func GetChildren(ix int) []int {
	return []int{ix*2 + 1, ix*2 + 2}
}
