package controllers

import (
	"go-aprendizaje/config"
	"go-aprendizaje/core"
	"go-aprendizaje/models"
	"go-aprendizaje/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// MongoRegister maneja el registro de un nuevo usuario en MongoDB
func MongoRegister(c *gin.Context) {
	// 1. Definir el "Input" (lo que esperamos del JSON)
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	// 2. Bindear el JSON del body a nuestro struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Comprobamos si el usuario ya existe
	existingUser, err := core.MongoUserRepo.GetUserByEmail(input.Email)
	if err == nil && existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El usuario ya existe"})
		return
	}

	// 3. Hashear la contraseña usando Bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo hashear la contraseña"})
		return
	}

	// 4. Crear la instancia del modelo MongoUser
	user := models.MongoUser{
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "user", // Mongo no tiene 'default' en el struct, lo definimos aquí
		// CreatedAt y UpdatedAt se definen en el repositorio
	}

	// 5. Guardar el usuario en la BD usando el REPOSITORIO GLOBAL
	newID, err := core.MongoUserRepo.CreateUser(&user)
	if err != nil {
		// (Esto podría ser un error de email duplicado)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se pudo crear el usuario. ¿Email ya existe? (Mongo)"})
		return
	}

	// 6. Enviar email de bienvenida en segundo plano
	go utils.SendWelcomeEmail(user.Email)

	// 7. Responder con éxito
	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado exitosamente (Mongo)",
		"userID":  newID.Hex(), // Devolvemos el ID de Mongo como string
		"email":   user.Email,
	})
}

// MongoLogin maneja el inicio de sesión de un usuario desde MongoDB
func MongoLogin(c *gin.Context) {
	// 1. Definir el "Input"
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	// 2. Bindear el JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// 3. Buscar al usuario en la BD usando el REPOSITORIO GLOBAL
	user, err := core.MongoUserRepo.GetUserByEmail(input.Email)
	if err != nil {
		// Chequeo específico de Mongo para "no encontrado"
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email o contraseña incorrectos (Mongo)"})
			return
		}
		// Otro error de base de datos
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al contactar la base de datos"})
		return
	}

	// 4. Comparar la contraseña del input con el hash guardado
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		// ¡Las contraseñas NO coinciden!
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email o contraseña incorrectos (Mongo)"})
		return
	}

	// 5. ¡Autenticación exitosa! Generar el Token JWT
	jwtSecret := config.GetEnv("JWT_SECRET_KEY", "fallback_secret")

	// 6. Crear los "Claims" (el payload del token)
	claims := jwt.MapClaims{
		"userID": user.ID.Hex(), // Importante: Usamos el ID de Mongo como string
		"role":   user.Role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
		"iat":    time.Now().Unix(),
	}

	// 7. Firmar el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo generar el token"})
		return
	}

	// 8. Devolver el token al cliente
	c.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso (Mongo)",
		"token":   tokenString,
	})
}
