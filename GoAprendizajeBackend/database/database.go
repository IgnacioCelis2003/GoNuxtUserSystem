package database

import (
	"fmt"
	"log"

	"go-aprendizaje/config"
	"go-aprendizaje/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// gorm.DB es la instancia de la base de datos que se utilizará en toda la aplicación
var DB *gorm.DB

// Funcion para conectar a la base de datos de Postgres
func ConnectDatabase() {
	var err error

	host := config.GetEnv("DB_HOST", "localhost")
	user := config.GetEnv("DB_USER", "postgres")
	password := config.GetEnv("DB_PASSWORD", "mysecretpassword")
	dbname := config.GetEnv("DB_NAME", "mi_api_db")
	port := config.GetEnv("DB_PORT", "5432")

	// Formatear la cadena de conexión con la base de datos de Postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, dbname, port)

	// gorm.Open abre la conexión a la base de datos en base a la cadena dsn
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("¡Conexión a la base de datos (Postgres) exitosa!")

	// AutoMigrate crea las tablas en la base de datos basándose en los modelos definidos
	DB.AutoMigrate(&models.User{})
}
