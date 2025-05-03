<template>
  <SplashScreen v-if="loading" />
  <div class="layout" v-else>
    <Sidebar />
    <div class="content">
      <RouterView />
    </div>
  </div>

</template>

<script setup>
import { ref, onMounted } from 'vue';
import Sidebar from './components/Sidebar.vue'
import SplashScreen from './components/SplashScreen.vue'
import Dashboard from './components/Dashboard.vue'

const loading = ref(true);

onMounted(() => {
  setTimeout(() => {
    loading.value = false;
  }, 3000); // Muestra el splash screen durante 3 segundos
});
</script>

<style scoped>
.layout {
  display: flex;
  height: 100%;
}
.content {
  flex: 1;
  padding: 1rem;
}
</style>
