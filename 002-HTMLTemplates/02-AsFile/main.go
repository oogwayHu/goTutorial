package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "John Doe"
	age := "15"

	template := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>"HelloWorld"</title>
	</head>
	<body>
	<h1>My name is ` + name + ` and my age is ` + age + `</h1>
	</body>
	</html>
	`

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("ERRRRRRR")
	}
	defer file.Close()

	io.Copy(file, strings.NewReader(template))
}
