# --- Etapa 1: "Builder" ---
# Usamos una imagen oficial de Go (basada en Alpine Linux) como base
FROM golang:1.24-alpine AS builder

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos los archivos de dependencias primero
# (Esto se cachea, así que Docker no vuelve a descargar todo cada vez)
COPY go.mod go.sum ./
RUN go mod download

# Copiamos todo el resto del código fuente
COPY . .

# Compilamos la aplicación.
# CGO_ENABLED=0 y GOOS=linux crean un binario estático para Linux
# -o /app/main le dice que el binario final se llame 'main'
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

# --- Etapa 2: "Final" ---
# Empezamos desde una imagen base de Alpine, que es súper ligera
FROM alpine:latest

# (Opcional) Instala certificados CA para que Go pueda hacer llamadas HTTPS
# (ej. al enviar emails)
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copiamos SÓLO el binario compilado de la etapa "builder"
COPY --from=builder /app/main .

# Expone el puerto 8080 (el que lee tu app del .env)
# Esto es solo documentación; no publica el puerto.
EXPOSE 8080

# El comando para ejecutar la aplicación cuando el contenedor arranque
CMD ["/app/main"]