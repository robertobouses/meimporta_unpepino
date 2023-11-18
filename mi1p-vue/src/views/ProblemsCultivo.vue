<template>
  <div>
    <h2>Problemas de los crops</h2>
    <div>
      <input v-model="busqueda" placeholder="Name del crop">
      <button @click="buscarCrop">Buscar</button>
    </div>

    <div v-if="cropsFiltrados.length > 0">
      <p v-for="(crop, index) in cropsFiltrados" :key="index">
        {{ Object.entries(crop).map(([key, value]) => `${key}: ${value}`).join(', ') }}
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
      cropsFiltrados: [],
    };
  },
  methods: {
    async buscarCrop() {
      console.log("Iniciando búsqueda...");
      console.log("Búsqueda actual:", this.busqueda);

      try {
        const response = await axios.post(`${apiUrl}/problemsCrop`, {
          name: this.busqueda,
        });

        console.log("Respuesta del servidor:", response.data);

        const resultados = response.data;

        if (Array.isArray(resultados) && resultados.length > 0) {
          this.cropsFiltrados = resultados;
          console.log("Crops después del filtro:", JSON.stringify(this.cropsFiltrados));
        } else {
          console.log("No se encontraron resultados para la búsqueda.");
          this.cropsFiltrados = [];
        }
      } catch (error) {
        console.error('Error al obtener datos del servidor:', error);
      }
    },
  },
  mounted() {
    console.log("Componente montado. Crops:", this.cropsFiltrados);
  },
};
</script>
