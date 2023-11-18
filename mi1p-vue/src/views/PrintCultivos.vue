<template>
  <div class="cultivos-container">
    <div v-if="cultivos.length > 0">
      <h2 class="cultivos-title">Cultivos</h2>
      <ul class="cultivos-list">
        <li v-for="cultivo in cultivos" :key="cultivo.idcultivo" class="cultivo-item">
          <p class="cultivo-nombre">{{ formatNombre(cultivo.informacioncultivo?.nombre) }}</p>
        </li>
      </ul>
    </div>
    <div v-else>
      <p class="loading-message">Cargando cultivos...</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import { apiUrl } from '@/path-to-config/config.js';

export default {
  data() {
    return {
      cultivos: [],
    };
  },
  created() {
    console.log('created se ejecutó');
    axios.get(`${apiUrl}/printCultivos`)
      .then(response => {
        console.log('Respuesta de Axios:', response);
        this.cultivos = response.data;
        console.log('Cultivos actualizados:', this.cultivos);
      })
      .catch(error => {
        console.error('Error al obtener los cultivos:', error);
      });
  },
  methods: {
    formatNombre(nombre) {
      return nombre ? nombre.charAt(0).toUpperCase() + nombre.slice(1).toLowerCase() : 'Nombre no disponible';
    },
  },
};
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
