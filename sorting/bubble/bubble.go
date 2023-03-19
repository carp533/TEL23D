package bubble

func Sort(a []int) []int {
	var (
		n      = len(a)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				a[i+1], a[i] = a[i], a[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
	return a
}
