<template>
    <div>
      <h1>Crear Nuevo Cultivo</h1>
      <form @submit.prevent="submitForm">
        <!-- Agrega los campos del formulario para recopilar los datos del cultivo -->
        <label for="siglas">Siglas:</label>
        <input type="text" id="siglas" v-model="cultivo.siglas" />
  
        <!-- Agrega otros campos aquí -->
  
        <button type="submit">Crear Cultivo</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    name: 'CreateCultivo',
    data() {
      return {
        cultivo: {
          siglas: '',
          // Agrega otros campos aquí
        }
      };
    },
    methods: {
    submitForm() {
      axios
        .post('http://tu-servidor-golang/api/crear-cultivo', this.cultivo)
        .then((response) => {
          // Maneja la respuesta del servidor, por ejemplo, muestra un mensaje de éxito.
        })
        .catch((error) => {
          // Maneja los errores, por ejemplo, muestra un mensaje de error.
        });
    }
  }
};
</script>
  
  <style>
  </style>
  