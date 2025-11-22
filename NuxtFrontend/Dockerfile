# --- ETAPA 1: CONSTRUCCIÓN (Builder) ---
# Usamos una imagen ligera de Node 18 o 20
FROM node:20-alpine AS builder

# Directorio de trabajo
WORKDIR /app

# Copiamos los archivos de dependencias
COPY package*.json ./

# Instalamos las dependencias (npm ci es más rápido y limpio para CI/CD)
RUN npm ci

# Copiamos el resto del código fuente
COPY . .

# Construimos la aplicación para producción
RUN npm run build

# --- ETAPA 2: EJECUCIÓN (Runner) ---
# Empezamos desde cero con una imagen limpia
FROM node:20-alpine

WORKDIR /app

# Copiamos solo la carpeta de salida (.output) desde la etapa anterior
# Nuxt 3 genera todo lo necesario dentro de .output
COPY --from=builder /app/.output /app/.output

# Exponemos el puerto 3000
EXPOSE 3000

# Comando para iniciar el servidor
# Nuxt usa 'node' para correr el archivo index.mjs generado
CMD ["node", ".output/server/index.mjs"]