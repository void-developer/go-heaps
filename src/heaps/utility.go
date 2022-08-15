package heaps

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
