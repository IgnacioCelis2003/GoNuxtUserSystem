<template>
  <v-app>
    <Header />

    <v-main class="bg-grey-lighten-5">
      <v-container class="py-10 h-100">
        <slot />
      </v-container>
    </v-main>

    <Footer />
    
    </v-app>
</template>