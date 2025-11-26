package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title       string
	VisitCount  int
	FormMessage string
}

func main() {
	// Servir les fichiers statiques (CSS, JS, Images)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Route principale
	http.HandleFunc("/", homeHandler)

	// Route pour le formulaire de contact
	http.HandleFunc("/contact", contactHandler)

	log.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title: "Portfolio - [Entrez votre nom ici]",
	}

	tmpl.Execute(w, data)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Récupération des données du formulaire
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	// Validation côté serveur
	if name == "" || email == "" || message == "" {
		http.Error(w, "Tous les champs sont requis", http.StatusBadRequest)
		return
	}

	// Ici vous pouvez traiter le message (envoyer un email, sauvegarder en DB, etc.)
	log.Printf("Message reçu de %s (%s): %s", name, email, message)

	// Redirection avec message de succès
	http.Redirect(w, r, "/?success=true", http.StatusSeeOther)
}
