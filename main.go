package main

import (
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

    http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
    pageVariables := PageVariables{
        Title: "Sample Site",
    }

    tmpl, err := template.New("index.html").ParseFiles(filepath.Join("templates", "index.html"), filepath.Join("templates", "header.html"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, pageVariables)
}

func ServicePage(w http.ResponseWriter, r *http.Request) {
    serviceName := r.URL.Path[1:]

    pageVariables := PageVariables{
        Title: "Service " + serviceName,
    }

    tmpl, err := template.New("service.html").ParseFiles(filepath.Join("templates", "service.html"), filepath.Join("templates", "header.html"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, pageVariables)
}

