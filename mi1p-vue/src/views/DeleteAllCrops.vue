<template>
  <div>
    <div class="delete-crop-container">
      <h1 class="delete-crop-title">Eliminar todos los Crops</h1>
    </div>
    <button @click="confirmEliminarCrops" class="delete-button">Eliminar todos los Crops</button>
    <div v-if="isLoading" class="loading-message">Eliminando crops...</div>
  </div>
</template>

<script>
import axios from 'axios';
import { apiUrl } from '@/path-to-config/config.js';

export default {
  data() {
    return {
      isLoading: false,
    };
  },
  methods: {
    confirmEliminarCrops() {
      if (confirm('¿Estás seguro de que deseas eliminar todos los crops?')) {
        this.eliminarAllCrops();
      }
    },
    eliminarAllCrops() {
      this.isLoading = true;

      axios
        .delete(`${apiUrl}/crops/all`)
        .then(response => {
          if (response.status === 200 || response.status === 204) {
            alert('Crops eliminados con éxito.');
          } else {
            alert('No se pudieron eliminar los crops.');
          }
        })
        .catch(error => {
          alert('Error al eliminar los crops:', error);
        })
        .finally(() => {
          this.isLoading = false;
        });
    },
  },
};
</script>

<style scoped>
.delete-crop-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f8f8f8;
  border: 1px solid #ccc;
  border-radius: 5px;
  margin-right: 20;
}


.delete-crop-title {
  font-size: 24px;
  color: #333;
  margin-bottom: 20px;
}

.delete-button {
  background-color: #e44d26;
  color: #fff;
  padding: 10px 20px;
  font-size: 16px;
  border: none;
  border-radius: 3px;
  cursor: pointer;
  
}

.loading-message {
  margin-top: 10px;
  font-size: 14px;
  color: #555;
}
</style>
