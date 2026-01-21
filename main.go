package main

import (
	"log"
	"net/http"
	"os"

	"portfolio/src/controller"
	"portfolio/src/router"
)

func main() {
	// Initialisation des templates
	controller.InitTemplates()

	// Configuration des routes
	router.SetupRoutes()

	// Récupération du port depuis l'environnement
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Port par défaut pour le développement local
	}

	// Démarrage du serveur
	log.Printf("Serveur démarré sur le port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
