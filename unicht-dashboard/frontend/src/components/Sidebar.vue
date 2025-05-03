<template>
  <div class="container">
    <h1 class="title">UniChat</h1>
    <div class="sidebar">
      <RouterLink to="/" class="nav-link">Inicio</RouterLink>
      <RouterLink to="/users" class="nav-link">Usuarios</RouterLink>
      <RouterLink to="/groups" class="nav-link">Grupos</RouterLink>
    </div>
  </div>
</template>

<style scoped>
.title {
  font-size: 2.5rem; 
  color: var(--fg);  
  margin: 1rem 0;    
  font-weight: bold; 
  text-align: center; 
  position: relative;
}

.title::after {
  content: ''; 
  position: absolute;
  bottom: -5px; 
  left: 0;
  right: 0;
  margin: 0 auto234;
  width: 70%;
  height: 3px; 
  background-color: var(--purple); 
}

.container {
  display: flex;
  flex-direction: column;
  border-right: 1px solid var(--purple);
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.2);
}
.sidebar {
  display: flex;
  flex-direction: column;
  width: 220px;
  height: 100%;
  background-color: var(--bg);
  padding: 1.5rem 1rem;
}

.nav-link {
  color: var(--fg);
  text-decoration: none;
  font-size: 1.1rem;
  margin: 0.5rem 0;
  padding: 0.75rem 1rem;
  border-radius: 8px;
  transition: background-color 0.2s ease, color 0.2s ease;
}

.nav-link:hover {
  background-color: var(--purple);
  color: var(--bg);
}

.router-link-active {
  background-color: var(--green);
  color: var(--bg);
}
</style>