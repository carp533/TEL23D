package main

/* Aufgabe: ermittle die K häufigsten Elemente einer Liste
 */

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	// TODO: erstelle die map count mit den Elementen aus nums und deren Häufigkeit
	//count := make(map[int]int)

	nums_ := make(Nums, 0)
	// TODO: wandle die map in die Liste nums_ um

	heap.Init(&nums_)
	var res []int
	// TODO: rufe k mal Pop auf des Heaps und schreibe das Ergebnis in die Liste res
	return res
}

// Num stores its value and frequency as Count.
type Num struct {
	Val   int
	Count int
}

// Nums struct for impl Interface
type Nums []Num

// Len sort Interface
func (n Nums) Len() int { return len(n) }

// Swap sort Interface
func (n Nums) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

// Less sort Interface
func (n Nums) Less(i, j int) bool { return n[i].Count >= n[j].Count }

// Push heap Interface
func (n *Nums) Push(num interface{}) {
	m := num.(Num)
	*n = append(*n, m)
}

// Pop heap Interface
func (n *Nums) Pop() interface{} {
	res := (*n)[len(*n)-1]
	*n = (*n)[:len(*n)-1]
	return res
}
