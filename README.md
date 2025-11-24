# GoNuxtUserSystem
Repositorio en donde cree un proyecto de un sistema de usuarios con go, Nuxt 4 y distintas herramientas como jwt, gomail, gorm, vuetify, etc...

## 📦 Dependencias del Backend (Go)

Aquí están todas las dependencias de Go necesarias para ejecutar el proyecto.

### 1. Framework y Núcleo

* **Gin (Framework Web):** El framework principal para manejar las rutas y peticiones HTTP.
    ```bash
    go get [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
    ```
* **Godotenv (Variables de Entorno):** Para cargar el archivo `.env` en desarrollo.
    ```bash
    go get [github.com/joho/godotenv](https://github.com/joho/godotenv)
    ```
* **Logrus (Logging):** Para un logging estructurado y profesional (escribir en `app.log`).
    ```bash
    go get [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
    ```
* **Gin CORS (Middleware):** Para permitir que tu frontend (Nuxt) se comunique con tu API (Go).
    ```bash
    go get [github.com/gin-contrib/cors](https://github.com/gin-contrib/cors)
    ```

### 2. Bases de Datos

* **GORM (ORM de Postgres):** El ORM que usamos para la lógica de PostgreSQL.
    ```bash
    go get gorm.io/gorm
    ```
* **Driver GORM para Postgres:** El conector específico que GORM necesita para hablar con Postgres.
    ```bash
    go get gorm.io/driver/postgres
    ```
* **Driver Oficial de MongoDB:** El driver oficial mantenido por Mongo para Go.
    ```bash
    go get go.mongodb.org/mongo-driver/mongo
    ```

### 3. Seguridad y Autenticación

* **Bcrypt (Hashing):** La librería estándar de Go para hashear (encriptar) contraseñas de forma segura.
    ```bash
    go get golang.org/x/crypto/bcrypt
    ```
* **JWT-Go (Tokens):** La librería más popular para crear y validar JSON Web Tokens (JWT).
    ```bash
    go get [github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt/v5)
    ```

### 4. Utilidades

* **Gomail (Envío de Emails):** Una librería simple y popular para enviar correos vía SMTP (Gmail).
    ```bash
    go get gopkg.in/gomail.v2
    ```
* **Google UUID (Nombres de Archivos):** Para generar nombres de archivo únicos (UUIDs) al subir imágenes.
    ```bash
    go get [github.com/google/uuid](https://github.com/google/uuid)
    ```

---

### Comando Único (Opcional)

Si estás en un proyecto nuevo, puedes instalar todas de una vez con este comando:

```bash
go get [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) [github.com/joho/godotenv](https://github.com/joho/godotenv) [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus) [github.com/gin-contrib/cors](https://github.com/gin-contrib/cors) gorm.io/gorm gorm.io/driver/postgres go.mongodb.org/mongo-driver/mongo golang.org/x/crypto/bcrypt [github.com/golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt/v5) gopkg.in/gomail.v2 [github.com/google/uuid](https://github.com/google/uuid)
```


## 🖥️ Dependencias del Frontend (Nuxt + Vuetify)

El frontend está construido con **Nuxt 4** y utiliza **Vuetify** como librería de componentes UI.

### 1. Instalación Inicial

Si clonas el repositorio por primera vez, instala todas las dependencias listadas en el `package.json`:

```bash
npm install
```
### 2. Librerías Principales
Estas son las librerias que utiliza el proyecto

* **vuetify: Framework de UI Material Design, su plugin para Vite (para tree-shaking) y los iconos.
    ```bash
    npm install vuetify vite-plugin-vuetify @mdi/font
    ```
* **Sass: Preprocesador CSS necesario para que Vuetify compile sus estilos.
    ```bash
    npm install sass
    ```
* **Tipos de Node: Necesario para evitar errores de TypeScript con variables globales como process.env.
    ```bash
    npm install -D @types/node
    ```

### Comando Único (Opcional)

Si estás en un proyecto nuevo, puedes instalar todas de una vez con este comando:

```bash
npm install vuetify vite-plugin-vuetify @mdi/font sass && npm install -D @types/node
```
    
