package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// fmt.Println("test")
	// Initialisation de la base de données
	// db, err := sqlite.NewDatabase()
	// if err != nil {
	// 	log.Fatalf("Error initializing the database: %v", err)
	// }
	// defer db.Close()

	// // Charger les variables d'environnement depuis le fichier .env
	// err = godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	// // Obtenir les configurations du serveur
	// serverPort := os.Getenv("SERVER_PORT")
	// if serverPort == "" {
	// 	serverPort = ":8080" // Port par défaut si non spécifié
	// }

	// // Initialiser les routes
	// router := routes.InitializeRoutes(db)

	// // Log avant le démarrage du serveur
	// log.Printf("Starting server on http://localhost%s", serverPort)

	// // Démarrer le serveur HTTP
	// err = http.ListenAndServe(serverPort, router)
	// if err != nil {
	// 	log.Fatalf("Error starting server: %v", err)
	// }

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("1"), bcrypt.DefaultCost)
	fmt.Println(string(hashedPassword))
}
