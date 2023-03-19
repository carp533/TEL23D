package selection

func Sort(a []int) []int {
	for i := 0; i < len(a)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[minIndex] {
				a[j], a[minIndex] = a[minIndex], a[j]
			}
		}
	}
	return a
}
