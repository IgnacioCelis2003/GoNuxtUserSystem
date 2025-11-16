package middleware

import (
	"errors"
	"go-aprendizaje/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el token del encabezado Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token no autorizado"})
			c.Abort()
			return
		}

		// Validar el formato del token (Bearer <token>)
		// Usamos strings.Split para separar "Bearer" del token en sí
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No autorizado: Formato de token inválido"})
			return
		}

		// Sacamos el token
		tokenString := parts[1]

		// Parsear y validar el token JWT

		// Obtener la clave secreta desde la configuración
		jwtSecret := config.GetEnv("JWT_SECRET_KEY", "fallback_secret")

		// Parsear el token
		// jwt.Parse toma el token y una función para obtener la llave secreta
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// (Validación extra: chequea que el método de firma sea el que esperamos)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("método de firma inesperado")
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			// (Si 'err' no es nil, el token es inválido, expiró, o la firma no coincide)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No autorizado: Token inválido o expirado"})
			return
		}

		// Extraer los "claims" (datos) del token y chequear que el token sea valido
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			// ¡ÉXITO! Guardar los datos del usuario en el "contexto" de Gin
			// Esto permite que el *siguiente* handler (el controlador)
			// pueda saber qué usuario está haciendo la petición.
			c.Set("userID", claims["userID"])
			c.Set("role", claims["role"])

			// 6. Permitir que la petición continúe
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No autorizado: Claims de token inválidos"})
			return
		}

	}
}
