// plugins/vuetify.ts
import '@mdi/font/css/materialdesignicons.css' // Importamos los iconos
import 'vuetify/styles' // Importamos los estilos base
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

export default defineNuxtPlugin((app) => {
  const vuetify = createVuetify({
    components,
    directives,
    
    // Configuración del tema (Opcional, por defecto es claro)
    theme: {
      defaultTheme: 'light',
    },
  })
  app.vueApp.use(vuetify)
})