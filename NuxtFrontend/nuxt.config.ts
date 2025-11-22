// https://nuxt.com/docs/api/configuration/nuxt-config

import vuetify from 'vite-plugin-vuetify'
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: { enabled: true },

  // Configuración del Build para Vuetify
  build: {
    transpile: ['vuetify'],
  },

  modules: [
    // Vuetify module
    /*(_options, nuxt) => {
      nuxt.hooks.hook('vite:extendConfig', (config) => {
        // @ts-expect-error
        config.plugins.push(vuetify({ autoImport: true }))
      })
    },*/
    '@nuxt/eslint',
    '@nuxt/image',
    '@nuxt/ui',
    '@nuxt/hints'
  ],

  runtimeConfig: {
    public: {
      apiBase: process.env.BACKEND_BASE_URL || 'http://localhost:8080'
    }
  },
})