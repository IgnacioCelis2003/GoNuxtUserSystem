pipeline {
    agent any // Usa cualquier "agente" (ejecutor) disponible

    stages {
        // Etapa 1: Obtener el código
        stage('Checkout') {
            steps {
                // Muestra un mensaje
                echo 'Clonando el repositorio...'
                // Git checkout se hace automático por la configuración del SCM,
                // pero este paso valida que estamos en la etapa correcta.
                checkout scm
            }
        }

        // Etapa 2: Verificar Variables de Entorno
        stage('Check Environment') {
            steps {
                echo 'Verificando archivos...'
                // En Linux/Docker usamos 'ls -la', en Windows 'dir'
                // sh es para shell de linux. Si tu jenkins corre en linux usa sh.
                sh 'ls -la' 
            }
        }

        // Etapa 3: Construcción del Backend (Simulación por ahora)
        stage('Build Backend') {
            steps {
                echo 'Construyendo Backend en Go...'
                // Aquí irían comandos como 'go build', pero
                // requeriría que Jenkins tenga Go instalado.
                // Por ahora solo imprimimos para probar el webhook.
            }
        }

        // Etapa 4: Construcción del Frontend (Simulación por ahora)
        stage('Build Frontend') {
            steps {
                echo 'Construyendo Frontend en Nuxt...'
            }
        }
    }
    
    post {
        success {
            echo '¡El Pipeline terminó exitosamente!'
        }
        failure {
            echo 'Algo falló en el pipeline.'
        }
    }
}