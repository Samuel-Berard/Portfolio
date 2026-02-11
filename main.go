package main

import (
	"log"
	"net/http"
	"os"

	"portfolio/src/controller"
	"portfolio/src/router"

	"github.com/joho/godotenv"
)

func main() {
	// Charger les variables d'environnement depuis .env
	if err := godotenv.Load(); err != nil {
		log.Println("Fichier .env non trouvé, utilisation des variables d'environnement système")
	}

	// Initialisation de la base de données
	if err := controller.InitDB(); err != nil {
		log.Fatalf("Erreur lors de l'initialisation de la base de données: %v", err)
	}

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
