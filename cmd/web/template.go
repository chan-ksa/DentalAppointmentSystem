package handlers

import (
	"html/template"
)

var tpl *template.Template

func TemplateInit() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
