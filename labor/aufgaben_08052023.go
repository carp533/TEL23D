package main

import (
	"fmt"
	"strings"
)

/*
Aufgabe: programmiere die funktion unique. die Funktion soll eine
Ae A als Eingabe bekommen und eine Liste ohne die doppelten Elemente
zurück geben.
Hinweis: benutze eine map, um die eindeutigen Werte zu ermitteln.
*/
func unique(A []int) []int {
	unique_set := make(map[int]bool, 0)
	for _, value := range A {
		unique_set[value] = true
	}
	result := make([]int, 0)
	for index := range unique_set {
		result = append(result, index)
	}
	return result
}

/*
Aufgabe: schreibe die Funktion unique_ord. Die Funktion soll die Reihenfolge
der Elemente in A beibehalten.
*/
func unique_ord(A []int) []int {
	unique_set := make(map[int]int, 0)
	i := 0
	for _, value := range A {
		if _, ok := unique_set[value]; !ok {
			unique_set[value] = i
			i++
		}

	}
	fmt.Println(unique_set)
	result := make([]int, len(unique_set))
	for key, val := range unique_set {
		result[val] = key
	}
	return result
}

/*
Aufgabe: programmiere die Funktion removeVowels. Die Funktion soll die Vokale
aus einem String entfernen.
Hinweis: verwende string.Builder, string.Builder.WriteRune, string.Builder.String
und string.ContainsAny
*/
func removeVowels(s string) string {
	vokale := "aeiouAEIOU"
	var sb strings.Builder
	for _, ch := range s {
		// fmt.Println(ch)
		if !strings.ContainsAny(string(ch), vokale) {
			sb.WriteRune(ch)
		}
	}
	return sb.String()
}

/*
Aufgabe: programmiere die Funktion symmdiff; diese soll die symmetrische
Differenz (a ohne b) plus (b ohne a) berechnen.
*/
func symdiff(a, b map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for key, _ := range a {
		if _, ok := b[key]; !ok {
			result[key] = true
		}
	}
	for key, _ := range b {
		if _, ok := a[key]; !ok {
			result[key] = true
		}
	}
	return result
}

/* */

func main() {
	// l := []int{1, 5, 3, 5, 3, 7}
	// fmt.Println(l)
	// fmt.Println(unique_ord(l))
	var a = map[string]bool{"John": true, "Bob": true, "Mary": true, "Serena": true}
	var b = map[string]bool{"Jim": true, "Mary": true, "John": true, "Bob": true}
	fmt.Println(symdiff(a, b))
	// fmt.Println(removeVowels("Hallo AAABBBCCC!"))

}
