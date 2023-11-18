<template>
  <main>
    <TheWelcome />
  </main>
</template>

<script setup>
import TheWelcome from '../components/TheWelcome.vue'
</script>

<style scoped>
.crops-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f8f8f8;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.crops-title {
  font-size: 24px;
  color: #333;
  margin-bottom: 20px;
}

.crops-list {
  list-style: none;
  padding: 0;
}

.crop-item {
  margin-bottom: 10px;
}

.crop-name {
  font-size: 16px;
  color: #555;
}

.loading-message {
  font-size: 16px;
  color: #555;
}
</style>
