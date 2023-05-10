package main

import "fmt"

func ExampleTopk() {
	array := []int{1, 1, 1, 1, 2, 2, 2, 3, 5, 3, 4}
	fmt.Println(topKFrequent(array, 3))
	//Output:
	//[1 2 3]
}
