package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)

type BlogPost struct {
	Title string
	Body  []byte
}

var templates = template.Must(template.ParseFiles("view.html"))

func (p *BlogPost) renderTemplate(w http.ResponseWriter) {
	err := templates.ExecuteTemplate(w, "view.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loadPost(title string) (*BlogPost, error) {
	filename := "posts/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &BlogPost{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	title = strings.TrimSuffix(title, "/")
	p, err := loadPost(title)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	p.renderTemplate(w)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
