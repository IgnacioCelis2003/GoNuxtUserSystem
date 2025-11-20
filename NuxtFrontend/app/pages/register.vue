<script setup lang="ts">
const config = useRuntimeConfig();
const isLoading = ref(false);

// Estado para notificaciones
const snackbar = reactive({
  show: false,
  text: '',
  color: 'info'
});

async function handleRegister(formData: any) {
  isLoading.value = true;

  // 1. Petición al Backend (Go)
  const { error } = await useFetch(`${config.public.apiBase}/api/users/register`, {
    method: 'POST',
    body: formData
  });

  isLoading.value = false;

  // 2. Manejo de Errores (Ej: Email duplicado)
  if (error.value) {
    snackbar.text = error.value.data?.error || 'No se pudo crear la cuenta';
    snackbar.color = 'error';
    snackbar.show = true;
    return;
  }

  // 3. Éxito
  snackbar.text = '¡Cuenta creada con éxito! Por favor inicia sesión.';
  snackbar.color = 'success';
  snackbar.show = true;

  // Redirigir al login después de 1.5 segundos
  setTimeout(() => {
    navigateTo('/login');
  }, 1500);
}
</script>

<template>
  <v-container class="fill-height justify-center bg-grey-lighten-5">
    
    <CustomAuthForm 
      title="Crear Cuenta" 
      subtitle="Únete a nuestra plataforma gratis" 
      btn-label="Registrarse"
      :loading="isLoading"
      is-register
      @submit="handleRegister"
    />

    <v-snackbar v-model="snackbar.show" :color="snackbar.color" location="top" timeout="4000">
      {{ snackbar.text }}
      <template v-slot:actions>
        <v-btn variant="text" @click="snackbar.show = false">Cerrar</v-btn>
      </template>
    </v-snackbar>

  </v-container>
</template>