package main

import (
	"fmt"
	"net/http"
)

func handleForm(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "Es ist ein Fehler bei ParseForm() aufgetreteten", err)
	}
	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(w, "Sie haben folgene Daten eingegeben:\n")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Adresse = %s\n", address)
}
func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/form", handleForm)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.ListenAndServe(":8090", nil)
}
