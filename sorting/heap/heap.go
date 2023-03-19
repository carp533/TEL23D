package heap

import "sort"

func Sort(a []int) []int {

	heapSort(sort.IntSlice(a))
	return a
}
func heapSort(list sort.Interface) {
	for start := (list.Len() - 2) / 2; start >= 0; start-- {
		siftDown(list, start, list.Len()-1)
	}
	for end := list.Len() - 1; end > 0; end-- {
		list.Swap(0, end)
		siftDown(list, 0, end-1)
	}
}

func siftDown(list sort.Interface, start, end int) {
	for root := start; root*2+1 <= end; {
		child := root*2 + 1
		if child+1 <= end && list.Less(child, child+1) {
			child++
		}
		if !list.Less(root, child) {
			return
		}
		list.Swap(root, child)
		root = child
	}
}
