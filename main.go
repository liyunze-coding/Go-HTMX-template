package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var PORT = "8008"

// PageData represents data to be passed to the HTML template
type PageData struct {
	Title string
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	pageData := PageData{
		Title: "Go + HTMX",
	}

	renderTemplate(w, "index.html", pageData)
}

func handleGreet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostFormValue("name")

	var greeting = fmt.Sprintf("Hello, %s! You've been greeting from Go!", name)

	fmt.Fprint(w, greeting)
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/greet", handleGreet)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/lib/", http.StripPrefix("/lib/", http.FileServer(http.Dir("lib"))))

	var display = fmt.Sprintf("Server is running on http://localhost:%s", PORT)

	fmt.Println(display)
	http.ListenAndServe(":8008", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := fmt.Sprintf("templates/%s", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
