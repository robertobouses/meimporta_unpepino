<template>
  <main>
    <TheWelcome />
  </main>
</template>

<script setup>
import TheWelcome from '../components/TheWelcome.vue'
</script>

<style scoped>
.cultivos-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f8f8f8;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.cultivos-title {
  font-size: 24px;
  color: #333;
  margin-bottom: 20px;
}

.cultivos-list {
  list-style: none;
  padding: 0;
}

.cultivo-item {
  margin-bottom: 10px;
}

.cultivo-nombre {
  font-size: 16px;
  color: #555;
}

.loading-message {
  font-size: 16px;
  color: #555;
}
</style>
