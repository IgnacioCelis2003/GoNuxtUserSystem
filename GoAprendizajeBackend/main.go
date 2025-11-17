package main

import (
	"go-aprendizaje/config"
	"go-aprendizaje/core"
	"go-aprendizaje/database"
	"go-aprendizaje/logging"
	"go-aprendizaje/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// Cargar la configuración
	//config.LoadConfig()
	port := config.GetEnv("SERVER_PORT", "8080")

	// Inicializar el sistema de logging
	logging.InitLogging()

	// Conectar a la base de datos
	database.ConnectDatabase()
	database.ConnectToMongoDB()

	// Inicializar los repositorios globales de MongoDB
	core.InitMongoRepositories()

	// Configurar el router
	router := gin.Default()
	routes.SetupRoutes(router)

	logging.Log.Info("Servidor iniciando en el puerto " + port)
	log.Println("Servidor iniciando en el puerto " + port)
	router.Run(":" + port)
}
