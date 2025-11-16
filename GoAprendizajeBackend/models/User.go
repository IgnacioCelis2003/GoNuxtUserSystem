package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

// User representa el modelo de usuario en la base de datos SQL (Postgres)
type User struct {
	gorm.Model              // Esto le dice a GORM que incluya los campos ID, CreatedAt, UpdatedAt, DeletedAt y que es un modelo de GORM
	Email            string `json:"email" gorm:"unique;not null"`
	Password         string `json:"password" gorm:"not null"`
	Role             string `json:"role" gorm:"default:'user';not null"`
	ProfileImagePath string `json:"profile_image_path" gorm:"default:null"`
}

// MongoUser representa el modelo de usuario en la base de datos MongoDB
type MongoUser struct {
	// ID de Mongo (_id).
	// 'primitive.ObjectID' es el tipo de dato de ID de Mongo
	// 'bson:"_id,omitempty"' le dice a Mongo:
	// - Usa esto como el campo '_id'
	// - Si está vacío, genéralo (omitempty)
	ID primitive.ObjectID `bson:"_id,omitempty"`

	// Usamos 'bson' en lugar de 'gorm' y 'json' porque es para MongoDB y mongo usa BSON
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"` // bson es el nombre de la variable en MongoDB
	Role     string `bson:"role" json:"role"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
