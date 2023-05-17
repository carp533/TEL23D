package main

import "fmt"

func maxl(x []int) []int {
	max := 0
	res := make([]int, len(x))
	for i := 0; i < len(x); i++ {
		if max < x[i] {
			max = x[i]
		}
		res[i] = max
	}
	return res
}
func maxr(x []int) []int {
	max := 0
	res := make([]int, len(x))
	for i := len(x) - 1; i >= 0; i-- {
		if max < x[i] {
			max = x[i]
		}
		res[i] = max
	}
	return res
}

func min(x, y []int) []int {
	res := make([]int, len(x))
	for i := 0; i < len(x); i++ {
		if x[i] < y[i] {
			res[i] = x[i]
		} else {
			res[i] = y[i]
		}
	}
	return res
}

// berechnet x[i] - y[i]
func diff(x, y []int) []int {
	res := make([]int, len(x))
	for i := 0; i < len(x); i++ {
		res[i] = x[i] - y[i]
	}
	return res
}

func sum(x []int) int {
	res := 0
	for _, v := range x {
		res += v
	}
	return res
}

func calculate(ein []int) int {
	ml := maxl(ein)
	mr := maxr(ein)
	min := min(ml, mr)
	diff := diff(min, ein)
	return sum(diff)
}
func ExampleCalc() {
	x := []int{3, 0, 0, 4, 6}
	fmt.Println(calculate(x))
	x = []int{3, 4, 5, 4, 1}
	fmt.Println(calculate(x))
	//Output:
	//6
	//0

}
func ExampleMin() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{3, 3, 3, 3, 3}
	fmt.Println(min(x, y))
	//Output:
	//[1 2 3 3 3]
}

func ExampleMaxL() {
	fmt.Println(maxl([]int{3, 2, 2, 6, 5, 5}))
	// Output:
	// [3 3 3 6 6 6]
}

func ExampleMaxR() {
	fmt.Println(maxr([]int{3, 2, 2, 6, 5, 5}))
	// Output:
	// [6 6 6 6 5 5]
}
