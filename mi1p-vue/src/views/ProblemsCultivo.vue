<template>
  <div>
    <h2>Problemas de los cultivos</h2>
    <div>
      <input v-model="busqueda" placeholder="Nombre del cultivo">
      <button @click="buscarCultivo">Buscar</button>
    </div>

    <div v-if="cultivosFiltrados.length > 0">
      <p v-for="(cultivo, index) in cultivosFiltrados" :key="index">
        {{ Object.entries(cultivo).map(([key, value]) => `${key}: ${value}`).join(', ') }}
      </p>
    </div>

    <div v-else>
      No se encontraron resultados para la búsqueda.
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      busqueda: "",
      cultivosFiltrados: [],
    };
  },
  methods: {
    async buscarCultivo() {
      console.log("Iniciando búsqueda...");
      console.log("Búsqueda actual:", this.busqueda);

      try {
        const response = await axios.post('http://localhost:8080/problemsCultivo', {
          name: this.busqueda,
        });

        console.log("Respuesta del servidor:", response.data);

        const resultados = response.data;

        if (Array.isArray(resultados) && resultados.length > 0) {
          this.cultivosFiltrados = resultados;
          console.log("Cultivos después del filtro:", JSON.stringify(this.cultivosFiltrados));
        } else {
          console.log("No se encontraron resultados para la búsqueda.");
          this.cultivosFiltrados = [];
        }
      } catch (error) {
        console.error('Error al obtener datos del servidor:', error);
      }
    },
  },
  mounted() {
    console.log("Componente montado. Cultivos:", this.cultivosFiltrados);
  },
};
</script>
