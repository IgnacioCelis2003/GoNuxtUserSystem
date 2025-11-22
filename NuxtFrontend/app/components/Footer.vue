<script setup lang="ts">
// función para navegar 
const goTo = (path: string) => window.open(path, '_blank'); // Abrir en una nueva pestaña, __blank es para eso
</script>


<template>
  <v-footer app class="bg-grey-lighten-4">
    <v-container>
      <v-row justify="center" no-gutters>
        
        <v-col cols="12" md="6" class="text-center text-md-left">
          <strong class="text-h6">GoApp</strong>
          <div class="text-caption text-grey-darken-1 mt-1">
            Aprendiendo Nuxt 3 + Golang + Vuetify
          </div>
        </v-col>

        <v-col cols="12" md="6" class="text-center text-md-right mt-4 mt-md-0">
          <v-btn icon="mdi-github" variant="text" size="small" @click="goTo('https://github.com/IgnacioCelis2003')"></v-btn>
          <v-btn icon="mdi-linkedin" variant="text" size="small" @click="goTo('https://cl.linkedin.com/in/ignacio-celis-castro-a2a246259?trk=people-guest_people_search-card')"></v-btn>
        </v-col>

        <v-col cols="12" class="text-center mt-4">
          <v-divider class="mb-4"></v-divider>
          <span class="text-caption text-grey">
            &copy; {{ new Date().getFullYear() }} — <strong>GoApp Inc.</strong>
          </span>
        </v-col>

      </v-row>
    </v-container>
  </v-footer>
</template>

