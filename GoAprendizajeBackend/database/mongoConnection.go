package database

import (
	"context"
	"go-aprendizaje/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo es la instancia global del cliente de MongoDB (la base de datos u conexión a esta que utilizaremos en toda la aplicación)
var Mongo *mongo.Database

// ConnectToMongoDB establece la conexión a la base de datos MongoDB
func ConnectToMongoDB() {
	// Obtener la URI de conexión desde la configuración
	uri := config.GetEnv("MONGO_URI", "mongodb://localhost:27017")
	dbName := config.GetEnv("MONGO_DB_NAME", "goAprendizaje")

	// Crear un "Context"
	// Mongo usa 'context' para manejar timeouts y cancelaciones.
	// context.WithTimeout le dice: "si no te conectas en 10 seg, falla".
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 'defer cancel()' es crucial: asegura que los recursos del context
	// se liberen cuando la función termine.
	defer cancel()

	// Configurar y abrir la conexión
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error fatal: No se pudo conectar a MongoDB: ", err)
	}

	// 4. "Ping" a la base de datos (para verificar que la conexión es real)
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Error fatal: No se pudo hacer 'ping' a MongoDB: ", err)
	}

	log.Println("¡Conexión a la base de datos (MongoDB) exitosa!")

	// 5. Asignar la instancia de la BD a nuestra variable global
	Mongo = client.Database(dbName)
}
