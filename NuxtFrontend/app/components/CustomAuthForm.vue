<script setup lang="ts">
defineProps({
  title: String,
  subtitle: String,
  btnLabel: String,
  loading: Boolean,
  isRegister: Boolean
});

const emit = defineEmits(['submit']); // Emitir eventos al componente padre esto hace defineEmits, crea un evento llamado 'submit' y el contendio se lo entrga al padre

// reactive: Crea un objeto que Vue "vigila". 
// Si escribes en el input, 'form.email' se actualiza en tiempo real.
const form = reactive({ email: '', password: '' });

// ref: Variable reactiva simple. Cambia entre true/false para mostrar/ocultar contraseña.
const showPassword = ref(false); // Para alternar ver/ocultar contraseña

// Reglas de validación simples (Vuetify las usa automáticamente)
const rules = {
  required: (v: any) => !!v || 'Este campo es obligatorio',
  email: (v: any) => /.+@.+\..+/.test(v) || 'El correo debe ser válido',
  min: (v: any) => v.length >= 6 || 'Mínimo 6 caracteres',
};
</script>

<template>
  <v-card class="mx-auto pa-4 elevation-4" max-width="400" rounded="lg">
    
    <div class="text-center mb-4">
      <v-avatar color="primary" size="50" class="mb-2">
        <v-icon icon="mdi-lock" color="white"></v-icon>
      </v-avatar>
      <h2 class="text-h5 font-weight-bold">{{ title }}</h2>
      <p class="text-body-2 text-grey">{{ subtitle }}</p>
    </div>

    <v-form @submit.prevent="emit('submit', form)">
      
      <v-text-field
        v-model="form.email"
        label="Correo Electrónico"
        prepend-inner-icon="mdi-email"
        variant="outlined"
        :rules="[rules.required, rules.email]"
        color="primary"
        class="mb-2"
      ></v-text-field>

      <v-text-field
        v-model="form.password"
        :type="showPassword ? 'text' : 'password'"
        label="Contraseña"
        prepend-inner-icon="mdi-lock"
        :append-inner-icon="showPassword ? 'mdi-eye' : 'mdi-eye-off'"
        @click:append-inner="showPassword = !showPassword"
        variant="outlined"
        :rules="[rules.required, rules.min]"
        color="primary"
      ></v-text-field>

      <v-btn
        type="submit"
        block
        color="primary"
        size="large"
        variant="flat"
        :loading="loading"
        class="mt-4"
      >
        {{ btnLabel }}
      </v-btn>
    </v-form>

    <v-divider class="my-4"></v-divider>

    <div class="text-center text-body-2">
      <template v-if="isRegister">
        ¿Ya tienes cuenta? 
        <NuxtLink to="/login" class="text-primary font-weight-bold text-decoration-none">Inicia sesión</NuxtLink>
      </template>
      <template v-else>
        ¿No tienes cuenta? 
        <NuxtLink to="/register" class="text-primary font-weight-bold text-decoration-none">Regístrate</NuxtLink>
      </template>
    </div>

  </v-card>
</template>