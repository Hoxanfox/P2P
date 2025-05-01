<template>
    <div class="dashboard">
      <div class="card" v-for="card in cards" :key="card.title">
        <h2>{{ card.title }}</h2>
        <p>{{ card.value }}</p>
      </div>
    </div>
    <div class="log-area">
      <h2>Logs del Servidor</h2>
      <textarea readonly :value="logs"></textarea>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue';

  import { GetActiveUsers, GetGroupCount, GetServerIP, GetOtherServers, GetLogs, GenerateTestLogs } from '../../wailsjs/go/main/App'
  
  const cards = ref([
    { title: 'Usuarios Activos', value: 'Cargando...' },
    { title: 'Total de Grupos', value: 'Cargando...' },
    { title: 'IP del Servidor', value: 'Cargando...' },
    { title: 'Otros Servidores', value: 'Cargando...' },
  ]);
  
  const logs = ref('Cargando logs...');

  onMounted(async () => {
    cards.value[0].value = await GetActiveUsers();
    cards.value[1].value = await GetGroupCount();
    cards.value[2].value = await GetServerIP();
    cards.value[3].value = (await GetOtherServers()).join(', ');
    loadLogs();

  });

  async function loadLogs() {
    console.log("Cargando logs...");
    await GenerateTestLogs();
    const result = await GetLogs();
    console.log("Logs recibidos:", result);
    logs.value = result;
  }

  </script>
  
  <style scoped>
  .dashboard {
    display: flex;
    flex-wrap: wrap;
    gap: 1rem;
    padding: 2rem;
    margin-bottom:0;
  }
  
  .card {
    background-color: var(--purple);
    color: var(--fg);
    padding: 1rem;
    border-radius: 8px;
    flex: 1 1 200px;
  }

  .log-area {
    display: flex;
    flex-direction: column;
    padding-left: 2rem;
    padding-right: 4rem;
  }

  textarea {
    width: 100%;
    min-height: 200px;
    background-color: #1e1f29;
    color: white;
    border: none;
    padding: 1rem;
    font-family: monospace;
    resize: none;
    border-radius: 8px;
  }
</style>
