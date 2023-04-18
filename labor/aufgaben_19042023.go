package main

import "fmt"

/*Aufgabe  1: variadic functions
* schreibe die Funktion average, die variabel viele Eingabeparamerter entgegen nimmt
* und den Durchschnitt der Zahlen berechnet
* func average(nums ...int)
 */
func average(nums ...int) float32 {
	if len(nums) == 0 {
		return 0
	}
	total := 0
	for _, num := range nums {
		total += num
	}
	return float32(total) / float32(len(nums))
}

/*Aufgabe  2:
// Erwartet eine sclice von Zahlen.
// Liefert true, falls die sclice an irgendeiner
// Stelle eine aufsteigende Kette von mindestens
// drei Zahlen enthält.
// die Zahlen müssen nicht direkt aufeinanderfolgen.*/
func ContainsChain2(list []int) bool {
	// TODO
	for i := 0; i < len(list)-2; i++ {

	}
	return false
}

/*Aufgabe  3: binary search tree implementation
* löse die TODO's
 */

// Struktur für einen Knoten
type Node struct {
	//TODO: vervollständige die Struktur für einen Binärbaum
	data rune
	// left
	// right
}

// Wurzel des Baums
type BinarySearchTree struct {
	root *Node
}

func (t *BinarySearchTree) insert(value rune) {
	if t.root != nil {
		t.root.insert(value)
	} else {
		//TODO Fall Baum leer (&Node ...)
	}
}

func (node *Node) insert(value rune) {
	if value < node.data {
		//TODO insert falls linker Knoten leer, sonst rekursiver insert
	}
	if value > node.data {
		//TODO insert falls rechter Knoten leer, sonst rekursiver insert
	}
}

func inorderTraversal(root *Node) []rune {
	// Traversiere den linken Teilbaum Inorder
	// Besuche die Wurzel
	// Traversiere den rechtenTeilbaum Inorder
	if root == nil {
		return nil
	}
	output := make([]rune, 0)
	//TODO inorder
	return output
}

func preorderTraversal(root *Node) []rune {
	// Besuche die Wurzel
	// Traversiere den linken Teilbaum Preorder
	// Traversiere den rechtenTeilbaum Preorder
	if root == nil {
		return nil
	}
	output := make([]rune, 0)
	//TODO preorder
	return output
}

func postorderTraversal(root *Node) []rune {
	// Traversiere den linken Teilbaum Postorder
	// Traversiere den rechtenTeilbaum Postorder
	// Besuche die Wurzel
	if root == nil {
		return nil
	}
	output := make([]rune, 0)
	//TODO postorder
	return output
}

func levelOrder(root *Node) [][]rune {
	if root == nil {
		return [][]rune{}
	}

	queue := []*Node{}
	queue = append(queue, root)
	result := [][]rune{}

	for len(queue) > 0 {
		sz := len(queue)
		level := []rune{}
		for i := 0; i < sz; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, rune(node.data))
			// TODO einkommentieren
			// if node.left != nil {
			// 	queue = append(queue, node.left)
			// }
			// if node.right != nil {
			// 	queue = append(queue, node.right)
			// }
		}
		result = append(result, level)
	}
	return result
}

func main() {

	nums := []int{}
	fmt.Println(average(nums...))
	nums = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(average(nums...))
	fmt.Println(average(1, 2, 3))

	fmt.Println(ContainsChain2([]int{1, 2, 1, 1, 1, 3}))
	fmt.Println(ContainsChain2([]int{1, 3, 2, 2, 1, 1}))
	fmt.Println(ContainsChain2([]int{3, 2, 3}))

	tree := &BinarySearchTree{}
	chars := []rune{'r', 'y', 't', 'b', 'i', 'n', 'a', 'r', 'e', 'e'}

	//chars := []rune{'b', 'i', 'n', 'a', 'r', 'y', 't', 'r', 'e', 'e'}

	for _, v := range chars {
		tree.insert(v)
	}

	fmt.Printf("Preorder:\n")
	for _, c := range preorderTraversal(tree.root) {
		fmt.Printf("%c ", c)
	}
	fmt.Printf("\nInorder:\n")
	for _, c := range inorderTraversal(tree.root) {
		fmt.Printf("%c ", c)
	}
	fmt.Printf("\nPostorder:\n")
	for _, c := range postorderTraversal(tree.root) {
		fmt.Printf("%c ", c)
	}
	fmt.Printf("\nLevelorder:\n")
	for _, arr := range levelOrder(tree.root) {
		for _, c := range arr {
			fmt.Printf("%c ", c)
		}
		fmt.Printf("\n")
	}

}
