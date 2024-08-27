package main

import (
	"log"
	"net/http"
	"os"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/routes"
	"socialNetwork/config"
)

func main() {
	//Initialisation de la base de données
	db, err := sqlite.NewDatabase()
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}
	defer db.Close()
	// On passe le db ouvert a global qui sera utiliser dans les service
	sqlite.GlobalDB = db


	// Obtenir les configurations du serveur
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ":8080"
	}
	// Initialiser les services
	config.InitServices()
	// Initialiser les routes
	router := routes.InitializeRoutes()

	// Log avant le démarrage du serveur
	log.Printf("Starting server on http://localhost%s", serverPort)

	// Démarrer le serveur HTTP
	err = http.ListenAndServe(serverPort, router)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}