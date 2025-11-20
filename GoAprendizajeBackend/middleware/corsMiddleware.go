// middleware/cors_middleware.go
package middleware

import (
	myConfig "go-aprendizaje/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupCorsConfig configura y devuelve el middleware de CORS
func SetupCorsConfig() gin.HandlerFunc {
	config := cors.DefaultConfig()

	url := myConfig.GetEnv("FRONTEND_URL", "http://localhost:3000")

	// A) Permitir origen específico
	config.AllowOrigins = []string{url}

	// B) Permitir métodos
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}

	// C) Permitir encabezados (Authorization es clave para el Token)
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	// D) Exponer encabezados (opcional, útil si necesitas leer headers en el front)
	config.ExposeHeaders = []string{"Content-Length"}

	// E) Permitir credenciales (cookies, auth headers)
	config.AllowCredentials = true

	return cors.New(config)
}
