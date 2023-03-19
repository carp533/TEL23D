package merge

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func ExampleSort() {
	array := []int{11, 14, 3, 8, 18, 17, 43}
	fmt.Println(Sort(array))
	//Output:
	//[3 8 11 14 17 18 43]
}
func BenchmarkSort(b *testing.B) {
	size := 100000
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	b.ResetTimer()
	Sort(slice)
}
