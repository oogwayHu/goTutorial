package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var s []string
	var input string

	// Open text file with read and write permissions
	file, err := os.OpenFile("sampleText.txt", os.O_RDWR, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read through each line in file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		s = append(s, text)
	}

	// If request is post, add input into file
	if r.Method == http.MethodPost {
		input = r.FormValue("input")

		s = append(s, input)
		input = "\n" + input

		_, err = io.WriteString(file, input)
		// file.WriteString(input)
		if err != nil {
			log.Fatal(err)
		}
	}

	// run template
	tpl.Execute(w, s)
}
