<template>
  <div class="splash-screen">
    <h1>Unichat <span class="loading-dots">...</span></h1>
    <div class="loading-bar">
      <div class="loading-progress"></div>
    </div>
  </div>
</template>

<script setup>
</script>


<style scoped>
.splash-screen {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background-color: var(--bg);
  color: var(--purple);
  font-size: 3rem;
  gap: 20px;
}

.loading-bar {
  width: 200px;
  height: 4px;
  background-color: rgba(128, 128, 128, 0.2);
  border-radius: 2px;
  overflow: hidden;
}

.loading-progress {
  width: 100%;
  height: 100%;
  background-color: var(--purple);
  animation: loading 1.5s infinite ease-in-out;
  transform-origin: left;
}

.loading-dots {
  animation: dots 1.5s infinite;
}

@keyframes dots {
  0%, 20% { opacity: 0; }
  40% { opacity: 0.3; }
  60% { opacity: 0.6; }
  80% { opacity: 0.9; }
  100% { opacity: 1; }
}

@keyframes loading {
  0% {
    transform: scaleX(0);
  }
  50% {
    transform: scaleX(1);
  }
  100% {
    transform: scaleX(0);
  }
}
</style>
  