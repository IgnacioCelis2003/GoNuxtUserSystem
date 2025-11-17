package logging

import (
	"go-aprendizaje/config"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

// Log es nuestra instancia global del logger
var Log *logrus.Logger

func InitLogging() {
	// 1. Crear una nueva instancia de Logrus
	Log = logrus.New()

	// 2. Configurar el formato
	// Queremos formato JSON para que sea "machine-readable"
	Log.SetFormatter(&logrus.JSONFormatter{})

	// 3. Configurar el nivel de log
	// (En producción, podrías leer esto desde el .env)
	// INFO: "Información general"
	// DEBUG: "Información para depurar" (muy verboso)
	// ERROR: "Algo falló"
	Log.SetLevel(logrus.InfoLevel)

	loggingPath := config.GetEnv("LOGGING_PATH", "/app/logging/app.log")

	// 4. Configurar la salida (¡la parte clave!)
	// Abrimos un archivo 'app.log'.
	// os.O_CREATE: Crea el archivo si no existe.
	// os.O_WRONLY: Solo escritura.
	// os.O_APPEND: Escribe al final del archivo (no lo sobrescribe).
	file, err := os.OpenFile(loggingPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// Si no podemos abrir el archivo de log, fallamos rápido
		log.Fatal("Error fatal: No se pudo abrir el archivo de log: ", err)
	}

	// Le decimos a Logrus que escriba en ese archivo
	Log.SetOutput(file)

	// (Opcional) También puedes hacer que escriba en la terminal Y en el archivo
	// mw := io.MultiWriter(os.Stdout, file)
	// Log.SetOutput(mw)

	Log.Info("Logging inicializado correctamente. Escribiendo en app.log")
}
