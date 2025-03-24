package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"
)

type TemplateParam struct {
	Title   string
	Message string
}

//go:embed templates
var f embed.FS

func main() {
	tmpl, err := template.ParseFS(f, "templates/index.html")
	if err != nil {
		fmt.Printf("Error: %v", err)

	}
	param := TemplateParam{Title: "Hello", Message: "World"}
	err = tmpl.Execute(os.Stdout, param)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
