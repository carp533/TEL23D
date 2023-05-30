package main

import (
	"fmt"
	"math/rand"
)

func ascending(n int) []int {
	s := make([]int, n)
	v := 1
	for i := 0; i < n; {
		if rand.Intn(3) == 0 {
			s[i] = v
			i++
		}
		v++
	}
	return s
}
func ExampleAscending() {
	fmt.Println(ascending(10))
	//Output:
	//1
}
func shuffled(n int) []int {
	return rand.Perm(n)
}
func ExampleSort() {
	var a = []int{170, 45, 75, -90, -802, 24, 2, 66}
	//a = shuffled(10)
	for inc := len(a) / 2; inc > 0; inc = (inc + 1) * 5 / 11 {
		for i := inc; i < len(a); i++ {
			j, temp := i, a[i]
			for ; j >= inc && a[j-inc] > temp; j -= inc {
				a[j] = a[j-inc]
			}
			a[j] = temp
		}
	}
	//Output:
	//1
}

func ForEach[T any](list []T, f func(T)) {
	for _, v := range list {
		f(v)
	}
}

func Sum(list []int) int {
	sum := 0
	f := func(v int) {
		sum += v
	}

	ForEach(list, f)

	return sum
}

func ExampleSum() {
	l := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(l))
	//Output:
	//15
}
