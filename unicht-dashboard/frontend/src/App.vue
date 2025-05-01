<template>
  <SplashScreen v-if="loading" />
  <Dashboard v-else />
</template>

<script setup>
import { ref, onMounted } from 'vue';
import SplashScreen from './components/SplashScreen.vue';
import Dashboard from './components/Dashboard.vue';

const loading = ref(true);

onMounted(() => {
  setTimeout(() => {
    loading.value = false;
  }, 3000); // Muestra el splash screen durante 3 segundos
});
</script>
