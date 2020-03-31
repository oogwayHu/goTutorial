package main

import (
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "tutanotaa")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	customServeMux()
	// defaultServeMux()
}

// Custom ServeMux
func customServeMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	// mux.Handle("/", http.HandlerFunc(index))

	http.ListenAndServe(":8080", mux)
}

func defaultServeMux() {
	http.HandleFunc("/", index)
	// http.Handle("/", http.HandlerFunc(index))

	http.ListenAndServe(":8080", nil)
}
