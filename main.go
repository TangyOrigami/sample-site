package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type PageVariables struct {
	Title string
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/service1", ServicePage)
	http.HandleFunc("/service2", ServicePage)
	http.HandleFunc("/service3", ServicePage)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Running on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		Title: "Sample Site",
	}

	renderTemplate(w, "index.html", pageVariables)
}

func ServicePage(w http.ResponseWriter, r *http.Request) {
	serviceName := r.URL.Path[1:]
	pageVariables := PageVariables{
		Title: "Service " + serviceName,
	}

	renderTemplate(w, "service.html", pageVariables)
}

func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	tmpl, err := template.New(templateName).ParseFiles(
		filepath.Join("templates", templateName),
		filepath.Join("templates", "header.html"),
		filepath.Join("templates", "footer.html"),
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

