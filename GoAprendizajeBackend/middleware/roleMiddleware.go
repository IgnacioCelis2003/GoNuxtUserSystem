package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtener el rol del usuario desde el contexto (establecido por AuthMiddleware)
		role, exists := c.Get("role")
		// Verificar si el rol existe en el contexto
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Acceso denegado: Rol de usuario no encontrado"})
			return
		}

		// Verificar si el rol tiene el formato correcto
		userRole, ok := role.(string)

		// Verificar si el rol del usuario coincide con el rol requerido
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Acceso denegado: Formato de rol inválido"})
			return
		}

		// Comparar roles
		if userRole != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Acceso denegado: No tienes permisos de " + requiredRole})
			return
		}
		c.Next()
	}
}
