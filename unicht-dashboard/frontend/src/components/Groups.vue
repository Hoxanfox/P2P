<script setup>
import { ref, computed } from 'vue'

const groups = ref([
  {
    id: 1,
    name: "Grupo 1",
    quantity: 2,
    creator: "Juan",
    members: [
      { id: 1, name: "Miembro 1" },
      { id: 2, name: "Miembro 2" },
    ],
  },
  {
    id: 2,
    name: "Grupo 2",
    quantity: 2,
    creator: "Ana",
    members: [
      { id: 3, name: "Miembro 3" },
      { id: 4, name: "Miembro 4" },
    ],
  },
])

const selectedGroup = ref(null)
const searchGroup = ref('')

const filteredGroups = computed(() =>
  groups.value.filter(group =>
    group.name.toLowerCase().includes(searchGroup.value.toLowerCase())
  )
)

const openModal = (group) => {
  selectedGroup.value = group
}

const closeModal = () => {
  selectedGroup.value = null
}
</script>

<template>
  <div class="groups-panel">
    <h1 class="title">Panel de Grupos</h1>
    <div class="search-bar">
      <input type="text" v-model="searchGroup" placeholder="Buscar por email" />
    </div>
    <table class="group-table">
      <thead>
        <tr>
          <th>Nombre</th>
          <th>Cantidad</th>
          <th>Creador</th>
          <th>Acción</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="group in filteredGroups" :key="group.id">
          <td>{{ group.name }}</td>
          <td>{{ group.quantity }}</td>
          <td>{{ group.creator }}</td>
          <td>
            <button class="action-button" @click="openModal(group)">Ver Miembros</button>
          </td>
        </tr>
      </tbody>
    </table>

    <div v-if="selectedGroup" class="modal">
      <div class="modal-content">
        <h2>Miembros de {{ selectedGroup.name }}</h2>
        <ul>
          <li v-for="member in selectedGroup.members" :key="member.id">
            {{ member.name }}
          </li>
        </ul>
        <button class="close-button" @click="closeModal">Cerrar</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.groups-panel {
  max-width: 100%;
  height: 100%;
  max-height: 920px;
  margin: 2rem auto;
  padding: 1rem;
  background-color: var(--bg);
  color: var(--fg);
  border: 1px solid var(--purple);
  border-radius: 12px;
  box-shadow: 0 0 10px rgba(189, 147, 249, 0.2);
}


.title {
  font-size: 2rem;
  color: var(--pink);
  margin-bottom: 1rem;
  border-bottom: 2px solid var(--purple);
  display: inline-block;
  padding-bottom: 0.5rem;
}

.search-bar {
  display: flex;
  gap: 1rem;
  margin-bottom: 1rem;
  justify-content: center;
}

.search-bar input {
  padding: 0.5rem;
  border-radius: 8px;
  border: 1px solid var(--purple);
  background-color: #1e1f29;
  color: var(--fg);
}

.group-table {
  width: 100%;
  border-collapse: collapse;
}

.group-table th,
.group-table td {
  border: 1px solid var(--purple);
  padding: 0.75rem;
  text-align: center;
}

.group-table th {
  background-color: #1e1f29;
}

.group-table td {
  background-color: var(--bg);
}

.action-button {
  background-color: var(--pink);
  color: var(--bg);
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 8px;
  cursor: pointer;
  font-weight: bold;
  transition: background 0.3s;
}

.action-button:hover {
  background-color: var(--green);
}

.modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(40, 42, 54, 0.9);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 10;
}

.modal-content {
  background-color: #1e1f29;
  padding: 2rem;
  border-radius: 12px;
  width: 90%;
  max-width: 400px;
  box-shadow: 0 0 10px var(--purple);
  text-align: center;
  color: var(--fg);
}

.modal-content ul {
  list-style: none;
  text-align: center;
}

.close-button {
  margin-top: 1rem;
  background-color: var(--cyan);
  color: var(--bg);
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 8px;
  cursor: pointer;
  font-weight: bold;
  transition: background 0.3s;
}

.close-button:hover {
  background-color: var(--pink);
}
</style>
