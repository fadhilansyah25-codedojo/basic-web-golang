package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name:    "Fadil",
			Gender:  "male",
			Hobbies: []string{"reading", "travelling", "buying things"},
			Info:    Info{"Webops Intern", "Sahid Sudirmans"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func (t Info) GetAffiliationDetailInfo() string {
	return "Have 31 Divisions"
}
