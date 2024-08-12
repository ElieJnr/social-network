package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"socialNetwork/pkg/db/sqlite"

	"github.com/joho/godotenv"
)

func main() {
	db, err := sqlite.NewDatabase()
	if err != nil {
		fmt.Println("ok")
		log.Println(err)
	}
	// Charger les variables d'environnement depuis le fichier .env
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dbURL := os.Getenv("DATABASE_URL")
	serverPort := os.Getenv("SERVER_PORT")

	// Log avant le démarrage du serveur
	log.Printf("Starting server on port %s with database %s", serverPort, dbURL)

	// Démarrer le serveur HTTP
	err = http.ListenAndServe(serverPort, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	db.Close()
}
