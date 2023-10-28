<template>
    <div>
      <h1>Me Importa un Pepino</h1>
      <button @click="navigateTo('/cultivo')">Crear Cultivo</button>
      <button @click="navigateTo('/cultivos')">Ver Cultivos</button>
      <button @click="navigateTo('/borrartodo')">Borrar Todos los Cultivos</button>
    </div>
  </template>
  
  <script>
  export default {
    methods: {
      navigateTo(route) {
        this.$router.push(route);
      }
    }
  }
  </script>
  