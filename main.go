package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bmizerany/pat"
)

type PageVariables struct {
	Title string
}

func main() {
	fmt.Print("asd;flkjas;dlfkja")
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(HomePage))
	mux.Get("/service1", http.HandlerFunc(ServicePage))
	mux.Get("/service2", http.HandlerFunc(ServicePage))
	mux.Get("/service3", http.HandlerFunc(ServicePage))

	mux.Get("/form", http.HandlerFunc(Form))
	mux.Post("/form", http.HandlerFunc(HandleForm))
	mux.Get("/confirmation", http.HandlerFunc(Confirmation))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	port := 8080
	serverAddr := fmt.Sprintf("http://localhost:%d", port)
	log.Printf("Running on: %s\n", serverAddr)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.Fatal("Server startup error:", err)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", PageVariables{Title: "Sample Site"})
}

func ServicePage(w http.ResponseWriter, r *http.Request) {
	serviceName := r.URL.Path[1:]
	renderTemplate(w, "service.html", PageVariables{Title: "Service " + serviceName})
}

func Form(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "confirmation.html", PageVariables{Title: "Sample Site"})
}

func HandleForm(w http.ResponseWriter, r *http.Request) {
	// Handle the form submission here
	// Retrieve form data using r.FormValue("fieldname")
	// Perform necessary actions (e.g., send email)

	// Respond with the confirmation template
	pageVariables := PageVariables{Title: "Sample Site"}
	renderTemplate(w, "confirmation.html", pageVariables)
}

func Confirmation(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "confirmation.html", PageVariables{Title: "Sample Site"})
}

func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	// Load templates once and cache them for better performance
	templates := template.Must(template.ParseFiles(
		filepath.Join("templates", templateName),
		filepath.Join("templates", "header.html"),
		filepath.Join("templates", "footer.html"),
	))

	err := templates.ExecuteTemplate(w, templateName, data)
	if err != nil {
		log.Printf("Error executing template %s: %v", templateName, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
