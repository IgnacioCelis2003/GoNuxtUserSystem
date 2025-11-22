<template>
  <v-app> // Componente principal de Vuetify que envuelve toda la aplicación, siempre debe estar presente y rodeando todo el contenido, sino Vuetify no funcionará correctamente
    <Header />

    <v-main class="bg-grey-lighten-5">
      <v-container class="py-10 h-100">
        <slot />
      </v-container>
    </v-main>

    <Footer />
    
  </v-app>
</template>