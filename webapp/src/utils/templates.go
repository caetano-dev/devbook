package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

//LoadTemplates loads all templates
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecuteTemplate executes a template
func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
