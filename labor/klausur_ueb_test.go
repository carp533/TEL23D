package main

import (
	"errors"
	"fmt"
	"sort"
)

/*
*********************************************************************
Aufgabe 1:
Programmiere einen Min-Stack, d.h. ein Datentyp, der die
Methoden push, pop, top und Min unterstützt.
Min soll das kleinste Element des Stacks in konstanter Zeit liefern
**********************************************************************
*/

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   []int{},
	}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	min, err := s.Min()
	if err == nil {
		if x < min {
			s.min = append(s.min, x)
		} else {
			s.min = append(s.min, min)
		}
	} else {
		s.min = append(s.min, x)
	}
}

func (s *MinStack) Pop() error {
	l1 := len(s.stack)
	if l1 == 0 {
		return errors.New("Empty Stack!")
	}
	s.stack = s.stack[:l1-1]
	s.min = s.min[:l1-1]
	return nil
}

func (s *MinStack) Top() (int, error) {
	l1 := len(s.stack)
	if l1 == 0 {
		return 0, errors.New("Empty Stack")
	}
	return s.stack[l1-1], nil
}

func (s *MinStack) Min() (int, error) {
	l1 := len(s.min)
	if l1 == 0 {
		return 0, errors.New("Empty Stack")
	}
	return s.min[l1-1], nil
}

func ExampleMinStack() {
}

/**********************************************************************
* Aufgabe 2 :
* AUFGABENSTELLUNG:
* Gegeben ist ein Binärbaum. Berechnen Sie die Summe der "inneren" Knoten
* des Baumes
* Hinweis: Ein Blatt ist ein Knoten, dessen linker und rechter Teilbaum
* nil ist.
**********************************************************************/

type BinNode struct {
	Val   int
	Left  *BinNode
	Right *BinNode
}

func InnerSum(root *BinNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	return root.Val + InnerSum(root.Left) + InnerSum(root.Right)
}

func (root *BinNode) InnerSum2() int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 0
	}
	return root.Val + root.Left.InnerSum2() + root.Right.InnerSum2()
}

/**********************************************************************
*Aufgabe 3:
*Programmiere die Sortierung nach Namen und Gewicht
/**********************************************************************/

type Grams int

func (g Grams) String() string { return fmt.Sprintf("%dg", int(g)) }

type Organ struct {
	Name   string
	Weight Grams
}

type Organs []*Organ

func (s Organs) Len() int      { return len(s) }
func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByName struct{ Organs }

func (s ByName) Less(i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}

type ByWeight struct{ Organs }

func (s ByWeight) Less(i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func ExampleSortO() {
	s := []*Organ{
		{"brain", 1340},
		{"heart", 290},
		{"liver", 1494},
		{"pancreas", 131},
		{"prostate", 62},
		{"spleen", 162},
	}

	sort.Sort(ByWeight{s})
	fmt.Println("Organs by weight:")
	printOrgans(s)

	sort.Sort(ByName{s})
	fmt.Println("Organs by name:")
	printOrgans(s)
	//Output:
	//1
}

func printOrgans(s []*Organ) {
	for _, o := range s {
		fmt.Printf("%-8s (%v)\n", o.Name, o.Weight)
	}
}

/**********************************************************************
* Aufgabe 4
* AUFGABENSTELLUNG:
* Gegeben sind zwei Listen. Berechne die Vereinigung beider Listen.
* Hinweis: Zahlen können mehrfach in den Listen vorkommen. Im Ergebnis
* sollen die Zahlen nur einmal vorkommen.
**********************************************************************/

func Union(l1, l2 []int) []int {
	result := []int{}
	m := make(map[int]bool)
	for i := 0; i < len(l1); i++ {
		m[l1[i]] = true
	}
	for i := 0; i < len(l2); i++ {
		m[l2[i]] = true
	}

	for key, _ := range m {
		result = append(result, key)
	}
	// alternative:
	// l1 = append(l1, l2...)
	// for i := 0; i < len(l1); i++ {
	// 	found := false
	// 	for j := 0; j < len(result); j++ {
	// 		if result[j] == l1[i] {
	// 			found = true
	// 		}
	// 	}
	// 	if !found {
	// 		result = append(result, l1[i])
	// 	}
	// }

	return result
}
