package repositories

import (
	"context"
	"go-aprendizaje/database"
	"go-aprendizaje/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoUserRepository maneja la lógica de BD para usuarios en MongoDB
type MongoUserRepository struct {
	// Collection es la conexión a la colección en MongoDB
	collection *mongo.Collection // es un puntero porque siempre queremos apuntar a la misma colección
}

// NewMongoUserRepository es un "constructor" para crear el repositorio
func NewMongoUserRepository() *MongoUserRepository {
	// Obtenemos la colección "users" de nuestra BD "Mongo"
	return &MongoUserRepository{
		collection: database.Mongo.Collection("users"), // crea una conexión a la colección "users"
	}
}

// CreateUser inserta un nuevo usuario en Mongo
// (Estos métodos serán los que implementen nuestra interfaz en el Paso 11)
// func (r *MongoUserRepository) indica que este método es un metodo asociado a un puntero de MongoUserRepository, es decir,
// cada que tengamos un puntero a MongoUserRepository podremos llamar a este método CreateUser
func (r *MongoUserRepository) CreateUser(user *models.MongoUser) (primitive.ObjectID, error) {
	// Ponemos fechas de creación
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// 'InsertOne' es el comando de Mongo
	// 'context.Background()' es un contexto simple "vacío"
	result, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// Devolvemos el ID generado por Mongo
	return result.InsertedID.(primitive.ObjectID), nil
}

// GetUserByEmail busca un usuario por su email en Mongo
func (r *MongoUserRepository) GetUserByEmail(email string) (*models.MongoUser, error) {
	var user models.MongoUser

	// 'bson.M' es un 'map' para construir queries de Mongo
	// { "email": email }
	filter := bson.M{"email": email}

	// 'FindOne' es el comando de Mongo
	err := r.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		// (err puede ser 'mongo.ErrNoDocuments')
		return nil, err
	}

	return &user, nil
}

func (r *MongoUserRepository) GetUsers() ([]models.MongoUser, error) {
	var users []models.MongoUser

	// 'Find' es el comando de Mongo para múltiples documentos
	// cursor es como un "iterador" que nos permite recorrer los resultados, este apunta al primer resultado, pues no se devuelve todo de una vez
	// sino que apunta a los resultados (BSONs) y con el cursor se van obteniendo uno por uno
	cursor, err := r.collection.Find(context.Background(), bson.M{}) // bson.M{} significa "sin filtro", es decir, todos los documentos
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background()) // Asegura que el cursor se cierre al final

	// Existen otras formas de hacer for en GO, como la normal que es for i := 0; i < n; i++
	// O la forma de rango: for index, value := range collection {}, en esta ultima para evitar el index se usa el guion bajo _
	// cursor.Next() avanza al siguiente documento retornando true, y devuelve false cuando ya no hay más documentos, cuando se llama por
	// primera vez apunta al primer documento
	for cursor.Next(context.Background()) { // Este tipo de for funciona como un while, pues en GO solo existe el for, su sintaxis es "for true {}"
		var user models.MongoUser
		if err := cursor.Decode(&user); err != nil { // Decode decodifica el documento actual en la variable user, ya que se obtiene un BSON
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil { // Verifica si hubo errores durante la iteración Err devuelve el error si hubo alguno
		return nil, err
	}

	return users, nil
}
