package models

import "time"

// PageData représente les données passées aux templates
type PageData struct {
	Title       string
	VisitCount  int
	FormMessage string
}

// Todo représente une tâche à faire
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// TodoPageData représente les données de la page Todo
type TodoPageData struct {
	Title string
	Todos []Todo
}
