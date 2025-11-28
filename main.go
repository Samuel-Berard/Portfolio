package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
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
		Title: "Portfolio - BERARD Samuel",
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

	// Enregistrer le message dans un fichier texte
	filename, fileContent, err := saveContactToFile(name, email, message)
	if err != nil {
		log.Printf("Erreur lors de l'enregistrement: %v", err)
		http.Error(w, "Erreur lors de l'enregistrement du message", http.StatusInternalServerError)
		return
	}

	// Ici vous pouvez traiter le message (envoyer un email, sauvegarder en DB, etc.)
	log.Printf("Message reçu de %s (%s): %s", name, email, message)

	// Proposer le téléchargement du fichier
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(fileContent))
}

func saveContactToFile(name, email, message string) (string, string, error) {
	// Créer le dossier "contacts" s'il n'existe pas
	if err := os.MkdirAll("contacts", 0755); err != nil {
		return "", "", err
	}

	// Générer un nom de fichier unique avec timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("contact_%s_%s.txt", timestamp, name)
	filepath := fmt.Sprintf("contacts/%s", filename)

	// Créer le contenu du fichier
	content := fmt.Sprintf("=== MESSAGE DE CONTACT ===\n\n")
	content += fmt.Sprintf("Date: %s\n", time.Now().Format("02/01/2006 à 15:04:05"))
	content += fmt.Sprintf("Nom: %s\n", name)
	content += fmt.Sprintf("Email: %s\n", email)
	content += fmt.Sprintf("\nMessage:\n%s\n", message)
	content += fmt.Sprintf("\n=========================\n")

	// Créer le fichier sur le serveur
	file, err := os.Create(filepath)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return "", "", err
	}

	return filename, content, nil
}
