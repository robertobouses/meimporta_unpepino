<template>
  <div>
    <div class="delete-cultivo-container">
      <h1 class="delete-cultivo-title">Eliminar todos los Cultivos</h1>
    </div>
    <button @click="confirmEliminarCultivos" class="delete-button">Eliminar todos los Cultivos</button>
    <div v-if="isLoading" class="loading-message">Eliminando cultivos...</div>
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
    confirmEliminarCultivos() {
      if (confirm('¿Estás seguro de que deseas eliminar todos los cultivos?')) {
        this.eliminarAllCultivos();
      }
    },
    eliminarAllCultivos() {
      this.isLoading = true;

      axios
        .delete(`${apiUrl}/deleteAllCultivos`)
        .then(response => {
          if (response.status === 200 || response.status === 204) {
            alert('Cultivos eliminados con éxito.');
          } else {
            alert('No se pudieron eliminar los cultivos.');
          }
        })
        .catch(error => {
          alert('Error al eliminar los cultivos:', error);
        })
        .finally(() => {
          this.isLoading = false;
        });
    },
  },
};
</script>

<style scoped>
.delete-cultivo-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  background-color: #f8f8f8;
  border: 1px solid #ccc;
  border-radius: 5px;
  margin-right: 20;
}


.delete-cultivo-title {
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
