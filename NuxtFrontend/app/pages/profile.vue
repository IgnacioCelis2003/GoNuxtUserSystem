<script setup lang="ts">
// 1. Importamos las herramientas necesarias
const config = useRuntimeConfig();
const { token } = useAuth();

// 2. Protección de Ruta: Si no hay token, mandamos al usuario al login
if (!token.value) {
  await navigateTo('/login');
}

// 3. Obtener datos del usuario desde la API de Go
// Usamos 'key' para que Nuxt sepa identificar esta petición
const { data: userData, refresh, error: fetchError } = await useFetch<any>(`${config.public.apiBase}/api/users/profile`, {
  key: 'user-profile',
  headers: {
    Authorization: `Bearer ${token.value}`
  }
});

// Si el token expiró o es inválido, la API da error -> redirigimos
if (fetchError.value) {
  navigateTo('/login');
}

// 4. Lógica para la URL de la Imagen
const userImage = computed(() => {
  console.log('Valor de userData:', userData.value?.user?.profile_image_path);
  const path = userData.value?.user?.profile_image_path;

  if (!path) return null; // Si no hay imagen, retorna null

  // Si la ruta ya es absoluta (ej: http://...), la usamos tal cual
  if (path.startsWith('http')) return path;

  // Si es relativa (ej: /static/...), le pegamos el dominio de la API
  if (path.startsWith('/static')) {
    return `${config.public.apiBase}${path}`;
  }

  console.log('User image path:', path);

  return path;
});

// 5. Variables para la subida de archivo
const file = ref<File[] | null>(null); // Vuetify devuelve un array de archivos
const isUploading = ref(false);

// Variables para el Snackbar (Notificación)
const snackbar = reactive({
  show: false,
  text: '',
  color: 'success'
});

// 6. Función para subir la imagen
async function uploadImage() {
  

  // Verificamos si es null o undefined
  if (!file.value) {
    console.error("ERROR: file.value es nulo o indefinido");
    alert("Error: No se detecta el archivo seleccionado");
    return;
  }

  // 2. OBTENER EL ARCHIVO REAL
  let archivoReal;

  // Vuetify a veces devuelve un array y a veces el objeto, dependiendo de la versión
  if (Array.isArray(file.value)) {
     console.log("Es un array. Largo:", file.value.length);
     if (file.value.length === 0) {
        console.error("ERROR: El array está vacío");
        return;
     }
     archivoReal = file.value[0];
  } else {
     console.log("No es un array, es un objeto directo");
     archivoReal = file.value;
  }

  console.log("Archivo a enviar:", archivoReal);

  // 3. PREPARAR EL ENVÍO
  isUploading.value = true;
  const formData = new FormData();
  formData.append('profile_picture', archivoReal); 

  console.log("Enviando petición a la API...");

  // 4. PETICIÓN
  const { error, data } = await useFetch(`${config.public.apiBase}/api/users/profile/picture`, {
    method: 'POST',
    body: formData,
    headers: {
      Authorization: `Bearer ${token.value}`
    }
  });


  isUploading.value = false;

  if (error.value) {
    // ... manejo de error
    console.error("Hubo un error en la petición");
  } else {
    // ... éxito
    console.log("¡Éxito!");
    file.value = null; 
    refresh(); 
  }
}
</script>

<template>
  <v-container class="py-10">
    <h1 class="text-h3 font-weight-bold mb-8 text-grey-darken-3">Mi Perfil</h1>

    <v-row>
      <v-col cols="12" md="4">
        <v-card elevation="3" rounded="lg" class="text-center pa-6">
          
          <div class="mb-4">
            <v-avatar size="150" color="grey-lighten-3" class="elevation-2">
              <v-img 
                v-if="userImage" 
                :src="userImage" 
                alt="Foto de perfil" 
                cover
              ></v-img>
              
              <v-icon 
                v-else 
                icon="mdi-account-circle" 
                size="150" 
                color="grey-lighten-1"
              ></v-icon>
            </v-avatar>
          </div>

          <h2 class="text-h5 font-weight-bold mb-1">
            {{ userData?.user?.email || 'Usuario' }}
          </h2>
          
          <v-chip color="primary" variant="flat" size="small" class="mt-2 font-weight-bold">
            {{ userData?.user?.role?.toUpperCase() || 'USER' }}
          </v-chip>

          <v-divider class="my-4"></v-divider>

          <div class="d-flex justify-space-between px-4 text-caption text-grey">
            <span>ID de Usuario:</span>
            <span class="font-weight-bold">{{ userData?.user?.id }}</span>
          </div>
        </v-card>
      </v-col>

      <v-col cols="12" md="8">
        
        <v-card class="mb-6" rounded="lg" elevation="2">
          <v-card-item>
            <v-card-title class="text-h6">Cambiar Foto de Perfil</v-card-title>
            <v-card-subtitle>Sube una imagen JPG o PNG (máx 5MB)</v-card-subtitle>
          </v-card-item>

          <v-card-text class="mt-2">
            <v-file-input
              v-model="file"
              label="Seleccionar nueva imagen"
              prepend-icon="mdi-camera"
              variant="outlined"
              density="comfortable"
              accept="image/png, image/jpeg, image/jpg"
              show-size
              :loading="isUploading"
              color="primary"
            ></v-file-input>
          </v-card-text>

          <v-card-actions class="pa-4 pt-0">
            <v-spacer></v-spacer>
            <v-btn 
              color="primary" 
              variant="flat" 
              :disabled="!file || file.length === 0" 
              :loading="isUploading"
              @click="uploadImage"
              prepend-icon="mdi-cloud-upload"
            >
              Subir Foto
            </v-btn>
          </v-card-actions>
        </v-card>

        <v-card color="red-lighten-5" class="border-red" rounded="lg" elevation="0" border>
          <div class="d-flex flex-column flex-sm-row align-center pa-6">
            <v-icon icon="mdi-credit-card-outline" color="red-darken-2" size="48" class="mb-4 mb-sm-0 mr-sm-6"></v-icon>
            
            <div class="text-center text-sm-left flex-grow-1">
              <div class="text-h6 font-weight-bold text-red-darken-4">Zona de Pagos</div>
              <div class="text-body-2 text-red-darken-3">
                Aquí integraremos la pasarela de Webpay Plus para realizar transacciones seguras.
              </div>
            </div>

            <div class="mt-4 mt-sm-0 ml-sm-4">
              <v-btn color="red-darken-2" elevation="2" prepend-icon="mdi-cash">
                Ir a Pagar
              </v-btn>
            </div>
          </div>
        </v-card>

      </v-col>
    </v-row>

    <v-snackbar v-model="snackbar.show" :color="snackbar.color" timeout="3000" location="top">
      {{ snackbar.text }}
      <template v-slot:actions>
        <v-btn variant="text" @click="snackbar.show = false">Cerrar</v-btn>
      </template>
    </v-snackbar>

  </v-container>
</template>

<style scoped>
.border-red {
  border-color: #ffcdd2 !important;
}
</style>