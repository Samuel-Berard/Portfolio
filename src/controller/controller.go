package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"portfolio/src/models"
	"strconv"
	"time"
)

var templates *template.Template

// InitTemplates charge tous les templates au démarrage de l'application
func InitTemplates() {
	var err error

	// Créer un FuncMap avec les fonctions helper
	funcMap := template.FuncMap{
		"countCompleted": func(todos []models.Todo) int {
			count := 0
			for _, todo := range todos {
				if todo.Completed {
					count++
				}
			}
			return count
		},
		"countRemaining": func(todos []models.Todo) int {
			count := 0
			for _, todo := range todos {
				if !todo.Completed {
					count++
				}
			}
			return count
		},
		"formatDate": func(t time.Time) string {
			return t.Format("02/01/2006 15:04")
		},
	}

	templates, err = template.New("").Funcs(funcMap).ParseGlob("src/templates/*.html")
	if err != nil {
		log.Fatalf("Erreur lors du chargement des templates: %v", err)
	}
	log.Println("Templates chargés avec succès")
}

// HomeHandler gère la page d'accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := models.PageData{
		Title: "Portfolio - BERARD Samuel",
	}

	if templates == nil {
		http.Error(w, "Templates non initialisés", http.StatusInternalServerError)
		return
	}

	err := templates.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ContactHandler gère la soumission du formulaire de contact
func ContactHandler(w http.ResponseWriter, r *http.Request) {
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

// saveContactToFile enregistre les informations de contact dans un fichier
func saveContactToFile(name, email, message string) (string, string, error) {
	// Créer le dossier "src/temp" s'il n'existe pas
	if err := os.MkdirAll("src/temp", 0755); err != nil {
		return "", "", err
	}

	// Générer un nom de fichier unique avec timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("contact_%s_%s.txt", timestamp, name)
	filepath := fmt.Sprintf("src/temp/%s", filename)

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

// ========== TODO APP HANDLERS ==========

const todosFile = "src/temp/todos.json"

// loadTodos charge les todos depuis le fichier JSON
func loadTodos() ([]models.Todo, error) {
	// Créer le dossier temp s'il n'existe pas
	if err := os.MkdirAll("src/temp", 0755); err != nil {
		return nil, err
	}

	// Si le fichier n'existe pas, retourner une liste vide
	if _, err := os.Stat(todosFile); os.IsNotExist(err) {
		return []models.Todo{}, nil
	}

	// Lire le fichier
	data, err := ioutil.ReadFile(todosFile)
	if err != nil {
		return nil, err
	}

	// Parser le JSON
	var todos []models.Todo
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

// saveTodos sauvegarde les todos dans le fichier JSON
func saveTodos(todos []models.Todo) error {
	// Créer le dossier temp s'il n'existe pas
	if err := os.MkdirAll("src/temp", 0755); err != nil {
		return err
	}

	// Convertir en JSON
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	// Écrire dans le fichier
	return ioutil.WriteFile(todosFile, data, 0644)
}

// TodoHandler affiche la page des todos
func TodoHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := loadTodos()
	if err != nil {
		log.Printf("Erreur lors du chargement des todos: %v", err)
		http.Error(w, "Erreur lors du chargement des todos", http.StatusInternalServerError)
		return
	}

	data := models.TodoPageData{
		Title: "Ma Todo App",
		Todos: todos,
	}

	if templates == nil {
		http.Error(w, "Templates non initialisés", http.StatusInternalServerError)
		return
	}

	err = templates.ExecuteTemplate(w, "todo.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// AddTodoHandler ajoute un nouveau todo
func AddTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/todo", http.StatusSeeOther)
		return
	}

	title := r.FormValue("title")
	if title == "" {
		http.Error(w, "Le titre est requis", http.StatusBadRequest)
		return
	}

	todos, err := loadTodos()
	if err != nil {
		log.Printf("Erreur lors du chargement des todos: %v", err)
		http.Error(w, "Erreur lors du chargement des todos", http.StatusInternalServerError)
		return
	}

	// Trouver le prochain ID
	nextID := 1
	for _, todo := range todos {
		if todo.ID >= nextID {
			nextID = todo.ID + 1
		}
	}

	// Créer le nouveau todo
	newTodo := models.Todo{
		ID:        nextID,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	todos = append(todos, newTodo)

	// Sauvegarder
	if err := saveTodos(todos); err != nil {
		log.Printf("Erreur lors de la sauvegarde: %v", err)
		http.Error(w, "Erreur lors de la sauvegarde", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/todo", http.StatusSeeOther)
}

// ToggleTodoHandler bascule l'état complété d'un todo
func ToggleTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/todo", http.StatusSeeOther)
		return
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	todos, err := loadTodos()
	if err != nil {
		log.Printf("Erreur lors du chargement des todos: %v", err)
		http.Error(w, "Erreur lors du chargement des todos", http.StatusInternalServerError)
		return
	}

	// Trouver et basculer le todo
	found := false
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = !todos[i].Completed
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Todo non trouvé", http.StatusNotFound)
		return
	}

	// Sauvegarder
	if err := saveTodos(todos); err != nil {
		log.Printf("Erreur lors de la sauvegarde: %v", err)
		http.Error(w, "Erreur lors de la sauvegarde", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/todo", http.StatusSeeOther)
}

// DeleteTodoHandler supprime un todo
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/todo", http.StatusSeeOther)
		return
	}

	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	todos, err := loadTodos()
	if err != nil {
		log.Printf("Erreur lors du chargement des todos: %v", err)
		http.Error(w, "Erreur lors du chargement des todos", http.StatusInternalServerError)
		return
	}

	// Trouver et supprimer le todo
	newTodos := []models.Todo{}
	for _, todo := range todos {
		if todo.ID != id {
			newTodos = append(newTodos, todo)
		}
	}

	// Sauvegarder
	if err := saveTodos(newTodos); err != nil {
		log.Printf("Erreur lors de la sauvegarde: %v", err)
		http.Error(w, "Erreur lors de la sauvegarde", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/todo", http.StatusSeeOther)
}
