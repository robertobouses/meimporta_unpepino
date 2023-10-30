<template>
    <div>
      <h1>Crear Nuevo Cultivo</h1>
      <form @submit="submitForm">
        <!-- Aquí coloca los campos del formulario -->
        <label for="nombre">Nombre:</label>
        <input type="text" id="nombre" v-model="cultivo.nombre" required>
  
        <!-- Agrega más campos según tu estructura de datos -->
  
        <button type="submit">Guardar Cultivo</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        cultivo: {
          nombre: '',
          // Agrega más campos aquí según tu estructura de datos
        },
      };
    },
    methods: {
      submitForm() {
        // Aquí puedes enviar los datos del cultivo al servidor (backend)
        // Puedes usar Axios u otra librería para hacer la solicitud POST al servidor.
      },
    },
  };
  </script>
  