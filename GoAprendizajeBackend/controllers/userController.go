package controllers

import (
	"go-aprendizaje/config"
	"go-aprendizaje/database"
	"go-aprendizaje/models"
	"go-aprendizaje/utils"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {

	// Estructura para enlazar los datos de entrada (el cuerpo que esperamos del JSON)
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	// Bindear (enlazar) el JSON de entrada a la estructura, mapea los datos del JSON a la estructura Go
	// Si el JSON no tiene la estructura esperada o falta un campo, devolver un error
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Verificar si el usuario ya existe
	var existingUser models.User
	result2 := database.DB.Where("email = ?", input.Email).First(&existingUser)
	if result2.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El usuario ya existe"})
		return
	}

	// Hashear la contraseña usando bcrypt
	// GenerateFromPassword toma la contraseña en bytes y un "costo" y devuelve la contraseña hasheada o un error.
	// El costo (DefaultCost) define qué tan lenta/segura será la encriptación.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hashear la contraseña"})
		return
	}

	// Crear un nuevo usuario con el email y la contraseña hasheada
	user := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	// Guardar el usuario en la base de datos
	result := database.DB.Create(&user)

	// Si hay un error al guardar el usuario, devolver un error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario: " + result.Error.Error()})
		return
	}

	// Enviar correo de bienvenida (de forma asíncrona para no bloquear la respuesta)
	go utils.SendWelcomeEmail(user.Email)

	// 7. Responder con éxito
	// Es una buena práctica no devolver la contraseña (ni el hash)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado exitosamente",
		"userID":  user.ID,
		"email":   user.Email,
	})

}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Buscar el usuario en la base de datos por email
	var user models.User

	// SELECT * FROM users WHERE email = input.Email LIMIT 1;
	result := database.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	// Comparar la contraseña hasheada con la contraseña proporcionada
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	// Generar un token JWT

	jwtSecret := config.GetEnv("JWT_SECRET_KEY", "fallback_secret")

	// Crear los claims del token
	claims := jwt.MapClaims{
		"userID": user.ID,
		"role":   user.Role,
		// "exp" (Expiration Time): Es OBLIGATORIO.
		// Define cuándo expira el token. (Ej. en 24 horas)
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		// "iat" (Issued At): Cuándo se emitió
		"iat": time.Now().Unix(),
	}

	// Crear el token con los claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token"})
		return
	}

	// Devolver el token al cliente
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func GetProfile(c *gin.Context) {
	// 1. Obtener los datos del usuario que el MIDDLEWARE puso en el contexto
	// Hacemos "type assertion" para convertir el tipo 'any' a lo que sabemos que es
	userID_any, _ := c.Get("userID")
	userID := userID_any.(float64) // JWT guarda números como float64

	role_any, _ := c.Get("role")
	role := role_any.(string)

	// 2. Buscar al usuario en la BD (opcional, pero buena práctica)
	// (En este punto ya sabemos que es válido, pero quizás queremos datos frescos)
	var user models.User
	if err := database.DB.First(&user, uint(userID)); err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// 3. Devolver los datos (sin la contraseña)
	c.JSON(http.StatusOK, gin.H{
		"message": "Perfil obtenido exitosamente",
		"user": gin.H{
			"id":    user.ID,
			"email": user.Email,
			"role":  role, // Podríamos usar 'user.Role' o el 'role' del token
		},
	})
}

// GetAllUsers es un controlador solo para administradores
func GetAllUsers(c *gin.Context) {
	var users []models.User

	// 1. Buscar todos los usuarios en la BD
	// (Usamos .Omit("Password") para no devolver los hashes)
	result := database.DB.Omit("Password").Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios"})
		return
	}

	// 2. Devolver la lista de usuarios
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// PgUploadProfilePicture maneja la subida de imágenes de perfil para usuarios en Postgres
func PgUploadProfilePicture(c *gin.Context) {
	// 1. Obtener el ID de usuario del token
	userID_any, _ := c.Get("userID")
	userID := uint(userID_any.(float64))

	// 2. Obtener el archivo del formulario
	file, err := c.FormFile("profile_picture")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se proporcionó ningún archivo"})
		return
	}

	// 3. Generar un nombre de archivo único
	extension := filepath.Ext(file.Filename) // Obtener la extensión del archivo
	uniqueFilename := uuid.New().String() + extension

	// Obtener la ruta de guardado desde la configuración o usar el valor por defecto
	filePathEnv := config.GetEnv("FILE_PATH", "./uploads")

	// 4. Definir la ruta de destino
	destinationPath := filepath.Join(filePathEnv, uniqueFilename) // Ruta completa donde se guardará el archivo (./uploads/nombre-archivo.ext)

	// 5. Guardar el archivo en ./uploads/
	if err := c.SaveUploadedFile(file, destinationPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar el archivo"})
		return
	}

	// 6. Generar la ruta de URL pública
	publicPath := filepath.ToSlash(filepath.Join("/static", uniqueFilename))

	// 7. Actualizar la base de datos (usando 'database.DB' global)
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Actualizamos el campo usando la BD global
	if err := database.DB.Model(&user).Update("profile_image_path", publicPath).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la ruta del archivo"})
		return
	}

	// 8. Devolver la respuesta
	c.JSON(http.StatusOK, gin.H{
		"message":   "Archivo subido exitosamente",
		"file_path": publicPath,
	})
}
