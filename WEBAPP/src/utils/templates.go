package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// Add all html templates on var templates
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// Render html pages
func RunTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
