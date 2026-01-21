package template

import (
	"html/template"
	"log"
)

var templates *template.Template

// InitTemplates charge tous les templates au démarrage de l'application
func InitTemplates() {
	var err error
	templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Erreur lors du chargement des templates: %v", err)
	}
	log.Println("Templates chargés avec succès")
}

// RenderTemplate exécute un template avec les données fournies
func RenderTemplate(name string, data interface{}) (*template.Template, error) {
	return templates, nil
}
