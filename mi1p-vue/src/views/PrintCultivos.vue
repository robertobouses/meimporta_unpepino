<template>
  <div class="crops-container">
    <div v-if="crops.length > 0">
      <h2 class="crops-title">Crops</h2>
      <ul class="crops-list">
        <li v-for="crop in crops" :key="crop.idcrop" class="crop-item">
          <p class="crop-name">{{ formatName(crop.cropinformation ?.name) }}</p>
        </li>
      </ul>
    </div>
    <div v-else>
      <p class="loading-message">Cargando crops...</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import { apiUrl } from '@/path-to-config/config.js';

export default {
  data() {
    return {
      crops: [],
    };
  },
  created() {
    console.log('created se ejecutó');
    axios.get(`${apiUrl}/printCrops`)
      .then(response => {
        console.log('Respuesta de Axios:', response);
        this.crops = response.data;
        console.log('Crops actualizados:', this.crops);
      })
      .catch(error => {
        console.error('Error al obtener los crops:', error);
      });
  },
  methods: {
    formatName(name) {
      return name ? name.charAt(0).toUpperCase() + name.slice(1).toLowerCase() : 'Name no disponible';
    },
  },
};
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
