package insertion

func Sort(a []int) []int {
	for j := 1; j < len(a)-1; j++ {
		elem := a[j]
		i := j - 1
		for i > -1 && a[i] > elem {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = elem
	}
	return a
}
