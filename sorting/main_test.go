package main

import (
	"bubble"
	"heap"
	"insertion"
	"merge"
	"quick"
	"selection"
	"testing"
)

var size = 1000

func BenchmarkBubble(b *testing.B) {
	slice := generateSlice(size)
	b.ResetTimer()
	bubble.Sort(slice)
}
func BenchmarkHeap(b *testing.B) {
	slice := generateSlice(size)
	b.ResetTimer()
	heap.Sort(slice)
}
func BenchmarkInsertion(b *testing.B) {
	slice := generateSlice(size)
	b.ResetTimer()
	insertion.Sort(slice)
}
func BenchmarkMerge(b *testing.B) {
	slice := generateSlice(size)
	b.ResetTimer()
	merge.Sort(slice)
}
func BenchmarkQuick(b *testing.B) {
	slice := generateSlice(size)
	b.ResetTimer()
	quick.Sort(slice)
}
func BenchmarkSelection(b *testing.B) {
	slice := generateSlice(size)
	b.ResetTimer()
	selection.Sort(slice)
}
