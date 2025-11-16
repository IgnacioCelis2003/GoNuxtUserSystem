package utils

import (
	"go-aprendizaje/config"
	"log"
	"strconv"

	"gopkg.in/gomail.v2"
)

// SendWelcomeEmail envía un correo de bienvenida al nuevo usuario
// Si falla solo lo loguea
func SendWelcomeEmail(toEmail string) {
	// 1. Leer la configuración del .env
	host := config.GetEnv("EMAIL_HOST", "smtp.gmail.com")
	portStr := config.GetEnv("EMAIL_PORT", "587")
	user := config.GetEnv("EMAIL_USER", "")
	password := config.GetEnv("EMAIL_PASSWORD", "")

	// 2. Convertir el puerto de string a int
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Printf("Error al convertir el puerto de email a int: %v", err)
		return // No podemos continuar si el puerto es inválido
	}

	// 3. Crear el mensaje (el correo)
	m := gomail.NewMessage()
	m.SetHeader("From", user)  // De: tu-correo@gmail.com
	m.SetHeader("To", toEmail) // Para: el-nuevo-usuario@dominio.com
	m.SetHeader("Subject", "¡Bienvenido a Mi API con Go!")
	m.SetBody("text/html", "¡Hola! <br><br>Gracias por registrarte en nuestra plataforma. Estamos felices de tenerte.<br><br>Saludos,<br>El equipo de Mi API")

	// 4. Configurar el "Dialer" (el que se conecta al servidor SMTP)
	// (host, puerto, usuario, contraseña)
	d := gomail.NewDialer(host, port, user, password)

	// 5. Enviar el correo
	log.Printf("Intentando enviar email de bienvenida a: %s", toEmail)
	if err := d.DialAndSend(m); err != nil {
		// Si hay un error, solo lo logueamos.
		// No queremos que esto detenga la ejecución principal.
		log.Printf("Error al enviar el email de bienvenida a %s: %v", toEmail, err)
	} else {
		log.Printf("Email de bienvenida enviado exitosamente a: %s", toEmail)
	}
}
