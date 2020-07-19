package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tmplt *template.Template

func init() {
	tmplt = template.Must(template.ParseGlob("templates/*.html"))
}

func homeF(res http.ResponseWriter, req *http.Request) {
	err := tmplt.ExecuteTemplate(res, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	fmt.Println("Start server localhost:2016")

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", homeF)
	http.ListenAndServe(":2016", nil)
}
