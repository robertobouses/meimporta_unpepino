<template>
  <div class="delete-cultivo-container">
    <h1 class="delete-cultivo-title">Eliminar Cultivo</h1>
    <label for="cultivoId" class="label">ID del Cultivo:</label>
    <input type="number" id="cultivoId" v-model="cultivoId" class="input" required>
    <button @click="eliminarCultivo" class="delete-button">Eliminar Cultivo</button>
  </div>
</template>

<script>
import axios from 'axios';
import 'sweetalert2/dist/sweetalert2.min.css';
import Swal from 'sweetalert2';
import { apiUrl } from '@/path-to-config/config.js';

export default {
  data() {
    return {
      cultivoId: '', 
    };
  },
  methods: {
    eliminarCultivo() {
      if (!this.cultivoId) {
        Swal.fire({
          icon: 'error',
          title: 'Oops...',
          text: 'Por favor, ingrese un ID de cultivo válido.',
        });
        return;
      }
      axios
        .delete(`${apiUrl}/deleteCultivo/${this.cultivoId}`)
        .then(response => {
          if (response.status === 200 || response.status === 204) {
            // Success message
            Swal.fire({
              icon: 'success',
              title: 'Éxito',
              text: 'Cultivo eliminado con éxito.',
            });
          } else {
            // Error message
            Swal.fire({
              icon: 'error',
              title: 'Error',
              text: 'No se pudo eliminar el cultivo. Verifique el ID.',
            });
          }
        })
        .catch(error => {
          // Error message
          Swal.fire({
            icon: 'error',
            title: 'Error',
            text: `Error al eliminar el cultivo: ${error}`,
          });
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
}

.delete-cultivo-title {
  font-size: 24px;
  color: #333;
  margin-bottom: 20px;
}

.label {
  display: block;
  margin-bottom: 8px;
  font-size: 14px;
  color: #555;
}

.input {
  width: 100%;
  padding: 8px;
  margin-bottom: 16px;
  font-size: 16px;
  border: 1px solid #ddd;
  border-radius: 3px;
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
</style>
