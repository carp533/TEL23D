package main

import (
	"fmt"
	"sort"
)

var pr = fmt.Println

/*Aufgabe 1: closures
* programmiere die Funktion adder. die Funktion soll ein closure (func(int) int) nutzen, um sich die
* Summe der Zahlen zu merken
 */
func adder() func(int) int {
	//TODO
	// sum := 0
	return nil
}
func demoAdder() {
	pr("Aufgabe 1 - closures")
	add1 := adder()
	add2 := adder()
	for i := 0; i < 5; i++ {
		fmt.Println("add1:", add1(i))
		fmt.Println("add2:", add2(i*2))
	}
}

/* Aufgabe 2: higher order functions
* die Funktion multiplyBy hat eine Funktion als Rückgabeparameter
* die zurückgegebene Funktion soll die Eingabe mit der Eingabe von
* multiplyBy multiplizieren
 */
func multiplyBy(x int) func(int) int {
	//TODO
	return nil
}
func demoMultiplyBy() {
	pr("Aufgabe 2 - higher order functions")
	timesTwo := multiplyBy(2)
	fmt.Println("2 * 3 =", timesTwo(21))
}

/* Aufgabe 3 - Bonusaufgabe
* ändere die Signatur von func2 von string zu int und schau dir die Fehlermeldungen an
* gib den Datentyp von func1-3 aus (nutze fmt.Pfintf("%T",var))
 */
func func1(f func(string) (string, string), g func(string, string) string) string {
	return g(f("a string"))
}
func func2(s string) (string, string) {
	return "func2 returns1 (" + s + ")", "func2 returns2 (hugo)"
}
func func3(s, t string) string { return "func3 called with (" + s + "," + t + ")" }
func Aufgabe3() {
	pr("Aufgabe 3 - functions gone crazy")
	fmt.Println(func1(func2, func3))
	//TODO
}

/* Aufgabe 4 - go builtin sort package
* programmiere eine Funktion, die eine Person nach
* 1. Nachname
* 2. Vorname
* 3. Alter
* sortiert
 */

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s: %d  ", p.FirstName, p.LastName, p.Age)
}

func DemoGoSort() {
	pr("Aufgabe 4 - Nutzen der eingebauten GO Sortierfunktionen")
	people := []Person{
		{"Bob", "Maier", 31},
		{"Bob", "Maier", 13},
		{"John", "Meier", 42},
		{"John", "Müller", 42},
	}
	sort.Slice(people, func(i, j int) bool {
		//TODO
		//Hinweis: dir Lösung beinhaltet mehrere if Abfragen
		return false
	})
	fmt.Println(people)
}

func main() {
	demoAdder()
	demoMultiplyBy()
	DemoGoSort()
	Aufgabe3()

}
