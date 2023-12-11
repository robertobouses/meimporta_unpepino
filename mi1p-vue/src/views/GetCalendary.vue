<template>
    <div>
      <h1>Get Crop Calendary</h1>
      <form @submit.prevent="getCropsCalendary">
        <label for="month">Month:</label>
        <input v-model="month" type="text" id="month" required>
  
        <label for="city">City:</label>
        <input v-model="city" type="text" id="city" required>
  
        <button type="submit">Get Crop Calendary</button>
      </form>
  
      <div v-if="crops.length > 0">
        <h2>Crops Calendary</h2>
        <ul>
          <li v-for="crop in crops" :key="crop.id">
            {{ crop.name }} - {{ crop.family }} - {{ crop.plantingDensity }}
          </li>
        </ul>
      </div>
  
      <div v-if="error">
        <p>Error: {{ error }}</p>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        month: "",
        city: "",
        crops: [],
        error: null,
      };
    },
    methods: {
      async getCropsCalendary() {
        try {
          // Realiza la petición HTTP al backend para obtener el calendario de cultivos
          const response = await this.$axios.get(
            `/api/crops/calendary?month=${this.month}&city=${this.city}`
          );
  
          // Asigna la respuesta a la variable crops
          this.crops = response.data;
          this.error = null;
        } catch (error) {
          // Maneja cualquier error que pueda ocurrir durante la petición
          this.error = "Error fetching crops calendary";
          this.crops = [];
        }
      },
    },
  };
  </script>
  
  <style scoped>
  /* Estilos específicos para este componente */
  </style>
  