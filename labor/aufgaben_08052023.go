package main

import (
	"fmt"
)

/*
Aufgabe: programmiere die funktion unique. die Funktion soll eine
Ae A als Eingabe bekommen und eine Ae ohne die doppelten Elemente
zurück geben.
Hinweis: benutze eine map, um die eindeutigen Werte zu ermitteln.
*/
func unique(A []int) []int {
	//unique_set := make(map[int]bool, len(A))
	return nil
}

/*
Aufgabe: schreibe die Funktion unique_ord. Die Funktion soll die Reihenfolge
der Elemente in A beibehalten.
*/
func unique_ord(A []int) []int {
	return nil
}

/*
Aufgabe: programmiere die Funktion removeVowels. Die Funktion soll die Vokale
aus einem String entfernen.
Hinweis: verwende string.Builder, string.Builder.WriteRune, string.Builder.String
und string.ContainsAny
*/
func removeVowels(s string) string {
	//vokale := "aeiouAEIOU"
	return ""
}

/*
Aufgabe: programmiere die Funktion symmdiff; diese soll die symmetrische
Differenz (a ohne b) plus (b ohne a) berechnen.
*/
func symdiff(a, b map[string]bool) map[string]bool {
	return nil
}

/* */

func main() {
	fmt.Println(unique([]int{1, 2, 3, 2, 3, 4}))
	var a = map[string]bool{"John": true, "Bob": true, "Mary": true, "Serena": true}
	var b = map[string]bool{"Jim": true, "Mary": true, "John": true, "Bob": true}
	fmt.Println(symdiff(a, b))

}
