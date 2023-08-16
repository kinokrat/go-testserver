// main
package main

import (
	"fmt"
	"net/http"
	"log"
)

func theformAhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "|X| r.ParseForm() got an error | %v\n", err)
		return
	}
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	address := r.FormValue("address")

	fmt.Fprintf(w, "YOUR NAME:\t%s\n", name)
	fmt.Fprintf(w, "YOUR SURNAME:\t%s\n", surname)
	fmt.Fprintf(w, "YOUR ADDRESS:\t%s\n", address)
}

func theformBhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "|X| r.ParseForm() got an error | %v\n", err)
		return
	}
	dr_license := r.FormValue("driving_license")
	insurance := r.FormValue("insurance")
	blood_type := r.FormValue("blood_type")

	fmt.Fprintf(w, "DRIVING LICENSE:\t%s\n", dr_license)
	fmt.Fprintf(w, "INSURANCE:\t\t%s\n", insurance)
	fmt.Fprintf(w, "BLOOD TYPE:\t\t%s\n", blood_type)
}

func theeasteregghandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/easteregg.html" {
		http.Error(w, "Well, you've found a place from outer space.", http.StatusNotFound)
		return
	}
	fmt.Println("You've accessed the Easter Egg!")
}

func main() {
	theserver := http.FileServer(http.Dir("./static"))
	fmt.Println("G'day!")
	http.Handle("/", theserver)
	http.HandleFunc("/formA", theformAhandler)
	http.HandleFunc("/formB", theformBhandler)
	http.HandleFunc("/easteregg", theeasteregghandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
