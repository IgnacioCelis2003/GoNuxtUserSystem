package core

import (
	"go-aprendizaje/repositories"
	"log"
)

// --- VARIABLES GLOBALES (Singleton) ---
// Aquí vivirán todas las instancias de repositorios de Mongo.
// Comienzan como 'nil'.

var MongoUserRepo *repositories.MongoUserRepository

// (Aquí podrías añadir: var MongoProductRepo *repositories.MongoProductRepository)

// InitMongoRepositories es la función que llamará 'main.go'
func InitMongoRepositories() {
	// Llama al constructor del repositorio para instanciar la variable global
	MongoUserRepo = repositories.NewMongoUserRepository()

	log.Println("Registro global de repositorios MongoDB inicializado.")
}
