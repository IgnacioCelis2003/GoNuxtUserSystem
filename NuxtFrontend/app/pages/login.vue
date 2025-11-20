<script setup lang="ts">
// Importaciones necesarias
const config = useRuntimeConfig();
const { setToken } = useAuth();
const router = useRouter();

// Estado local
const isLoading = ref(false);
const snackbar = reactive({
  show: false,
  text: '',
  color: 'error'
});

// Función que recibe los datos del componente AuthForm ({ email, password })
async function handleLogin(formData: any) {
  isLoading.value = true;

  // 1. Petición al Backend (Go)
  const { data, error } = await useFetch<any>(`${config.public.apiBase}/api/users/login`, {
    method: 'POST',
    body: formData
  });

  isLoading.value = false;

  // 2. Manejo de Errores
  if (error.value) {
    snackbar.text = error.value.data?.error || 'Error al iniciar sesión';
    snackbar.color = 'error';
    snackbar.show = true;
    return;
  }

  // 3. Éxito: Guardar token y redirigir
  if (data.value?.token) {
    setToken(data.value.token);
    snackbar.text = '¡Bienvenido de nuevo!';
    snackbar.color = 'success';
    snackbar.show = true;
    
    // Esperamos un milisegundo para que se vea el mensaje y redirigimos
    setTimeout(() => {
      navigateTo('/profile');
    }, 500);
  }
}
</script>

<template>
  <v-container class="fill-height justify-center bg-grey-lighten-5">
    
    <CustomAuthForm 
      title="Bienvenido" 
      subtitle="Ingresa tus credenciales para continuar" 
      btn-label="Iniciar Sesión"
      :loading="isLoading"
      @submit="handleLogin"
    />

    <v-snackbar v-model="snackbar.show" :color="snackbar.color" location="top" timeout="3000">
      {{ snackbar.text }}
      <template v-slot:actions>
        <v-btn variant="text" @click="snackbar.show = false">Cerrar</v-btn>
      </template>
    </v-snackbar>

  </v-container>
</template>