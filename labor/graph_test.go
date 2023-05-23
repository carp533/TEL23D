package main

import "fmt"

/*Beispiel: Abbildung eines Graphen mit einer Adjazenz-Matrix*/
/*Aufgabe: Zeichne den Graphen (ordne die Knoten in einem 3x3
Gitter an*/
var AdjacencyMatrix = [][]int{
	//////// 1, 2, 3, 4, 5, 6, 7, 8, 9
	/* 1 */ {0, 1, 0, 1, 0, 0, 0, 0, 0},
	/* 2 */ {1, 0, 1, 0, 1, 0, 0, 0, 0},
	/* 3 */ {0, 1, 0, 0, 0, 1, 0, 0, 0},
	/* 4 */ {1, 0, 0, 0, 1, 0, 1, 0, 0},
	/* 5 */ {0, 1, 0, 1, 0, 1, 0, 1, 0},
	/* 6 */ {0, 0, 1, 0, 1, 0, 0, 0, 1},
	/* 7 */ {0, 0, 0, 1, 0, 0, 0, 1, 0},
	/* 8 */ {0, 0, 0, 0, 1, 0, 1, 0, 1},
	/* 9 */ {0, 0, 0, 0, 0, 1, 0, 1, 0},
}

/*Beispiel: Abbildung eines Graphen mit einer Adjazenz-Liste*/
/*Aufgabe: Zeichne den Graphen (ordne die Knoten in einem 3x3
Gitter an*/
var AdjacencyList = map[int][]int{
	1: {2, 4},
	2: {3, 1},
	3: {2},
	// 2: {3, 5, 1},
	// 3: {6, 2},
	4: {1, 5, 7},
	5: {2, 4, 6, 8},
	6: {3, 5, 9},
	7: {4, 8},
	8: {5, 9, 7},
	9: {6, 8},
}

//Breitensuche
func bfs(start int, alist map[int][]int) (numLvls int, ordered []int, lvls [][]int) {
	visited := map[int]bool{start: true} // bereits besuchte Knoten

	queue := []int{start}
	for len(queue) > 0 {
		level := []int{} // besuchte Knoten zu jedem Level

		// besuche Knoten zu Level:
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			current := queue[0]
			queue = queue[1:]

			ordered = append(ordered, current)
			level = append(level, current)

			// speichere nicht besuchte Nachbarn für jeden Level
			for _, neighbor := range alist[current] {
				if visited[neighbor] {
					continue
				}
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
		numLvls++
		lvls = append(lvls, level)
	}

	return numLvls, ordered, lvls
}
func ExampleGraphBFS() {
	numLevels, ordered, levels := bfs(3, AdjacencyList)
	fmt.Println(numLevels)
	fmt.Println(ordered)
	fmt.Println(levels)
	//Output:
	//?
}

func ExampleGraphDFS() {
	actualPrevisit, actualPostvisit := dfs(1, AdjacencyList)
	fmt.Println(actualPrevisit)
	fmt.Println(actualPostvisit)
	//Output:
	//?
}

//Tiefensuche
func dfs(start int, list map[int][]int) (previsit, postvisit []int) {
	visited := map[int]bool{start: true} // bereits besucht Knoten

	var dfs func(key int)
	dfs = func(key int) {
		visited[key] = true

		previsit = append(previsit, key)

		for _, neighbor := range list[key] {
			fmt.Printf("%v, -> %v (besucht: %v)\n", key, neighbor, visited[neighbor])
			if visited[neighbor] {
				continue
			}
			dfs(neighbor)
		}

		postvisit = append(postvisit, key)
	}

	dfs(start)
	return previsit, postvisit
}
