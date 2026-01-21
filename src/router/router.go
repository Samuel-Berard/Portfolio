package router

import (
	"net/http"

	"portfolio/src/controller"
)

// SetupRoutes configure toutes les routes de l'application
func SetupRoutes() {
	// Servir les fichiers statiques (CSS, JS, Images)
	fs := http.FileServer(http.Dir("src/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes de l'application
	http.HandleFunc("/", controller.HomeHandler)
	http.HandleFunc("/contact", controller.ContactHandler)

	// Routes de la Todo App
	http.HandleFunc("/todo", controller.TodoHandler)
	http.HandleFunc("/todo/add", controller.AddTodoHandler)
	http.HandleFunc("/todo/toggle", controller.ToggleTodoHandler)
	http.HandleFunc("/todo/delete", controller.DeleteTodoHandler)
}
