package main

import "fmt"

func main() {
	// anonyme Funktion
	// achte auf die Klammern ()
	func() {
		fmt.Println("ich bin eine anonyme Funktion")
	}()

	// anonyme Funktion wird einer Variablen zugewiesen
	a := func() {
		fmt.Println("ich bin eine anonyme Funktion II")
	}
	a()
	fmt.Printf("Signatur der Funktion a: %T\n", a)

	// anonyme Funktion mit Eingabeparameter
	func(n string) {
		fmt.Println("anonyme Funktion mit Eingabeparameter. Hello", n)
	}("TEL22D")

	// Funktionen als Eingabeparameter
	f := func(a, b int) int {
		return a * b
	}
	// die Funktion funcInput hat eine Funktion (f) als Eingabe
	funcInput(f)

	f2 := funcOutput()
	fmt.Println("Funktion als Ausgabe. ", f2(6, 7))

	//closure: eine anonyme Funktion, die auf eine "äußere" Variable zugreift
	z := 5
	func() {
		fmt.Println("z =", z)
	}()

	//closure -> jede closure "hält" ihre Variablen
	c1 := appendStr()
	c2 := appendStr()
	fmt.Println(c1("World"))
	fmt.Println(c2("Everyone"))

	fmt.Println(c1("Gopher"))
	fmt.Println(c2("!"))

}

// die Funktion funcInput hat eine Funktion als Eingabe
func funcInput(a func(a, b int) int) {
	fmt.Println("Funktion als Eingabe. ", a(4, 5))
}

// die Funktion funcOutput hat eine Funktion als Rückgabeparameter
func funcOutput() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

// Beispiel für closure & Gültigkeit Variablen
func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
