<template>
  <v-container class="text-center py-16">
    <v-row justify="center">
      <v-col cols="12" md="8">
        <v-chip color="primary" class="mb-4" label>Versión 1.0</v-chip>
        
        <h1 class="text-h2 font-weight-bold mb-4 text-grey-darken-4">
          Desarrollo Web Moderno con <span class="text-primary">Go</span> & <span class="text-success">Vuetify</span>
        </h1>
        
        <p class="text-h6 text-grey text-regular mb-8">
          Una arquitectura completa, escalable y hermosa. Backend potente, Frontend elegante.
        </p>

        <div class="d-flex justify-center gap-4">
          <v-btn color="primary" size="x-large" to="/register" elevation="2">
            Comenzar Gratis
          </v-btn>
          <v-btn variant="outlined" size="x-large" to="/login">
            Ingresar
          </v-btn>
        </div>
      </v-col>
    </v-row>

    <v-row class="mt-16">
      <v-col cols="12" md="4">
        <v-card class="pa-6" elevation="0" >
          <v-icon icon="mdi-speedometer" color="primary" size="48" class="mb-4"></v-icon>
          <h3 class="text-h5 font-weight-bold mb-2">Rápido</h3>
          <p class="text-body-1 text-grey">Backend en Go compilado y Frontend optimizado por Nuxt.</p>
        </v-card>
      </v-col>
      <v-col cols="12" md="4">
        <v-card class="pa-6" elevation="0" >
          <v-icon icon="mdi-shield-check" color="success" size="48" class="mb-4"></v-icon>
          <h3 class="text-h5 font-weight-bold mb-2">Seguro</h3>
          <p class="text-body-1 text-grey">Autenticación JWT robusta y protección de rutas.</p>
        </v-card>
      </v-col>
      <v-col cols="12" md="4">
        <v-card class="pa-6" elevation="0">
          <v-icon icon="mdi-palette" color="purple" size="48" class="mb-4"></v-icon>
          <h3 class="text-h5 font-weight-bold mb-2">Hermoso</h3>
          <p class="text-body-1 text-grey">Componentes Material Design listos para usar con Vuetify.</p>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>