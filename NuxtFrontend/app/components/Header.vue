<script setup lang="ts">
const { token, logout } = useAuth();

// Función para navegar (Vuetify usa router.push o nuxt link)
const goTo = (path: string) => navigateTo(path); // navigateTo es para navegar entre paginas del proyecto
</script>

<template>
  <v-app-bar app color="white" elevation="1">
    
    <v-container class="d-flex align-center py-0">
      
      <v-app-bar-title class="font-weight-bold text-h5 cursor-pointer" @click="goTo('/')">
        Go<span class="text-primary">App</span>
      </v-app-bar-title>

      <v-spacer></v-spacer> <div class="d-none d-md-flex align-center gap-2">
        
        <v-btn variant="text" @click="goTo('/')">Inicio</v-btn>

        <template v-if="!token">
          <v-btn variant="text" @click="goTo('/login')">Ingresar</v-btn>
          <v-btn color="primary" variant="flat" @click="goTo('/register')">Registrarse</v-btn>
        </template>

        <template v-else>
          <v-btn variant="text" prepend-icon="mdi-account" @click="goTo('/profile')">
            Mi Perfil
          </v-btn>
          <v-btn color="error" variant="text" prepend-icon="mdi-logout" @click="logout">
            Salir
          </v-btn>
        </template>
      </div>

    </v-container>
  </v-app-bar>
</template>