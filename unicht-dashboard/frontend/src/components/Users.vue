<script setup>
import { computed, ref } from 'vue';

const searchEmail = ref('')
const filterStatus = ref('')

const filteredUsers = computed(() => {
    return users.value.filter(user => {
        const matchesEmail = user.email.toLowerCase().includes(searchEmail.value.toLowerCase())
        const matchesStatus = filterStatus.value === '' || user.status === filterStatus.value
        return matchesEmail && matchesStatus
    })
})

const users = ref([
    { name: 'Juan Pérez', email: 'juan@example.com', ip: '192.168.1.10', status: 'activo' },
    { name: 'Ana López', email: 'ana@example.com', ip: '192.168.1.11', status: 'ausente' },
    { name: 'Carlos Ruiz', email: 'carlos@example.com', ip: '192.168.1.12', status: 'activo' },
])

function handleAction(user) {
    alert(`Acción ejecutada a: ${user.name}`)
}
</script>

<template>
    <div class="user-panel">
        <h1 class="title">Panel de usuarios</h1>
        <div class="search-bar">
            <input type="text" v-model="searchEmail" placeholder="Buscar por email" />
            <select v-model="filterStatus">
                <option value="">Todos</option>
                <option value="activo">Activo</option>
                <option value="ausente">Ausente</option>
            </select>
        </div>
        <table class="user-table">
            <thead>
                <tr>
                    <th>Nombre</th>
                    <th>Email</th>
                    <th>IP</th>
                    <th>Estado</th>
                    <th>Acción</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="user in filteredUsers" :key="user.email">
                    <td>{{ user.name }}</td>
                    <td>{{ user.email }}</td>
                    <td>{{ user.ip }}</td>
                    <td :class="user.status">{{ user.status }}</td>
                    <td>
                        <button @click="handleAction(user)">Desconectar</button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>

<style scoped>
.user-panel {
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

.search-bar input,
.search-bar select {
    padding: 0.5rem;
    border-radius: 8px;
    border: 1px solid var(--purple);
    background-color: #1e1f29;
    color: var(--fg);
}

.search-bar select {
    appearance: none;           /* Remueve el estilo nativo */
    -webkit-appearance: none;
    -moz-appearance: none;
}
.user-table {
    width: 100%;
    border-collapse: collapse;
}

.user-table th,
.user-table td {
    border: 1px solid var(--purple);
    padding: 0.75rem;
    text-align: center;
}

.user-table th {
    background-color: #1e1f29;
}

.user-table td.activo {
    color: var(--green);
}

.user-table td.ausente {
    color: var(--pink);
}

button {
    padding: 0.5rem 1rem;
    background-color: var(--pink);
    color: var(--bg);
    border: none;
    border-radius: 6px;
    cursor: pointer;
    transition: background-color 0.3s;
}

button:hover {
    background-color: var(--green);
    color: var(--bg);
}
</style>
