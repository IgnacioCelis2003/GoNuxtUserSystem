package routes

import (
	"go-aprendizaje/config"
	"go-aprendizaje/controllers"
	"go-aprendizaje/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configura todas las rutas de la aplicación
// (Debe ser Pública, con 'S' mayúscula En go las funciones publicas empiezan con mayúscula)
func SetupRoutes(router *gin.Engine) {

	// Obtener la ruta de archivos estáticos desde la configuración o usar el valor por defecto
	uploadDir := config.GetEnv("UPLOAD_PATH", "./uploads")

	// Servir archivos estáticos.
	// Cualquier petición GET a /static/[nombre-archivo]
	// servirá el archivo desde la carpeta ./uploads
	// Es un GET
	router.Static("/static", uploadDir)

	api := router.Group("/api")
	{

		api.GET("/ping", controllers.Ping)

		// Rutas para usuario
		userRoutes := api.Group("/users")
		{
			// Ruta para registrar un nuevo usuario
			userRoutes.POST("/register", controllers.RegisterUser)

			// Ruta para iniciar sesión
			userRoutes.POST("/login", controllers.Login)

			// Ruta protegida para obtener el perfil del usuario
			// Se añade el middleware de autenticación
			// Es como una cadena ejecución, primero el middleware y luego el controlador
			userRoutes.GET("/profile", middleware.AuthMiddleware(), controllers.GetProfile)

			// Rutas para usuario en MongoDB
			userRoutes.POST("/mongo/register", controllers.MongoRegister)
			userRoutes.POST("/mongo/login", controllers.MongoLogin)

			// Ruta para subir foto de perfil
			userRoutes.POST("/profile/picture",
				middleware.AuthMiddleware(),
				controllers.PgUploadProfilePicture,
			)

		}

		// Rutas para admin
		adminRoutes := api.Group("/admin")
		adminRoutes.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"))
		{
			// Ruta protegida para obtener usuarios (solo accesible por admin)
			adminRoutes.GET("/users", controllers.GetAllUsers)
		}

	}

}
