package controller

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"portfolio/src/models"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var templates *template.Template
var db *sql.DB

// InitDB initialise la connexion à la base de données MySQL
func InitDB() error {
	var err error
	// Récupérer les variables de connexion depuis l'environnement
	dbUser := os.Getenv("SCALINGO_MYSQL_USER")
	dbPassword := os.Getenv("SCALINGO_MYSQL_PASSWORD")
	dbHost := os.Getenv("SCALINGO_MYSQL_HOST")
	dbPort := os.Getenv("SCALINGO_MYSQL_PORT")
	dbName := os.Getenv("SCALINGO_MYSQL_DB")

	// Vérifier que toutes les variables sont définies
	if dbUser == "" || dbHost == "" || dbPort == "" || dbName == "" {
		return fmt.Errorf("variables de base de données manquantes (DB_USER, DB_HOST, DB_PORT, DB_NAME)")
	}

	// Construire le DSN: user:password@tcp(host:port)/database?parseTime=true
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture de la connexion: %v", err)
	}

	// Vérifier la connexion
	if err = db.Ping(); err != nil {
		return fmt.Errorf("erreur lors du ping de la base de données: %v", err)
	}

	log.Println("Connexion à la base de données MySQL réussie")
	return nil
}

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

// loadTodos charge les todos depuis la base de données MySQL
func loadTodos() ([]models.Todo, error) {
	rows, err := db.Query("SELECT id, title, completed, created_at FROM todo ORDER BY id DESC")
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la requête SELECT: %v", err)
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt); err != nil {
			return nil, fmt.Errorf("erreur lors du scan: %v", err)
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erreur lors du parcours des lignes: %v", err)
	}

	return todos, nil
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

	// Insérer le nouveau todo dans la base de données
	_, err := db.Exec("INSERT INTO todo (title, completed) VALUES (?, false)", title)
	if err != nil {
		log.Printf("Erreur lors de l'insertion: %v", err)
		http.Error(w, "Erreur lors de l'insertion", http.StatusInternalServerError)
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

	// Basculer l'état du todo
	_, err = db.Exec("UPDATE todo SET completed = NOT completed WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de la mise à jour: %v", err)
		http.Error(w, "Erreur lors de la mise à jour", http.StatusInternalServerError)
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

	// Supprimer le todo
	_, err = db.Exec("DELETE FROM todo WHERE id = ?", id)
	if err != nil {
		log.Printf("Erreur lors de la suppression: %v", err)
		http.Error(w, "Erreur lors de la suppression", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/todo", http.StatusSeeOther)
}
