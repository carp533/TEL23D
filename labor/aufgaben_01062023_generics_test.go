package main

/*Aufgaben Labor 6.1.2023*/

// FindCond sucht alle Indizes, für die f(v) wahr ist. v sind die Elemente
// der Liste list
// TODO: Ändere den Datentyp, sodass die Funktion eine Liste mit einem generischen
// Typ verarbeiten kann (type constraint any)
func FindCond(list []int, f func(int) bool) []int {
	res := make([]int, 0)
	for i, v := range list {
		if f(v) {
			res = append(res, i)
		}
	}
	return res
}

func Greater50(v int) bool {
	return v > 50
}

// FindIfGreater50 erwartet eine Liste und liefert die Position aller Elemente,
// die größer als 50 sind.
// TODO: programmiere die Funktion und nutze Greater50
func FindIfGreater50(list []int) []int {
	return nil
	//TODO
}

// FindIfOdd erwartet eine Liste und liefert die Position aller Elemente,
// die ungerade sind.
// TODO: programmiere die Funktion und nutze eine anonyme Funktion
func FindIfOdd(list []int) []int {
	return nil
	//TODO
}

// FindIfBetween erwartet eine Liste und zwei Zahlen a,b.
// Sie liefert alle Position, die zwischen a und b liegen.
//TODO: programmiere die Funktion, mit Hilfe einer anonymen Funktion
func FindIfBetween(list []int, a, b int) []int {
	return nil
	//TODO
}

//TODO: Programmiere eine Example Funktion, die alle obigen Funktionen testet

//TODO: definiere einen Datentyp Numbers, der die Datentypen int, int64, float und float64
//beinhaltet und programmiere die funktion FindIfBetween für diesen Datentyp.
