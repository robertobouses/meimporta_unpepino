<template>
  <div>
    <h1>Define My Crop</h1>
    <form @submit.prevent="defineMyCrop">
      <label for="name">Name:</label>
      <input type="text" id="name" v-model="myCrop.name" required />

      <label for="planting">Planting Date:</label>
      <input type="date" id="planting" v-model="myCrop.planting" required />

      <label for="fieldId">Field ID:</label>
      <input type="text" id="fieldId" v-model="myCrop.fieldId" required />

      <button type="submit">Define My Crop</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';
import { apiUrl } from '@/path-to-config/config.js';

export default {
  data() {
    return {
      myCrop: {
        name: '',
        planting: '', // Utilizamos un formato de cadena para la fecha
        fieldId: '',
      },
    };
  },
  methods: {
    async defineMyCrop() {
  try {
    // Formatear la fecha en el formato yyyy-MM-dd
    const fechaFormateada = new Date(this.myCrop.planting).toISOString().split('T')[0];

    const response = await axios.post(`${apiUrl}/mycrops/define`, {
      name: this.myCrop.name,
      planting: fechaFormateada,
      fieldid: this.myCrop.fieldId,
    });

    if (response.status === 200) {
      console.log('Cultivo creado exitosamente');
      const mensaje = response.data.mensaje;
      console.log('Mensaje:', mensaje);
    } else {
      console.error('Error al crear el cultivo:', response.statusText);
    }
  } catch (error) {
    console.error('Error al enviar la solicitud:', error);

    // Agrega esta línea para mostrar detalles específicos del error en la consola
    console.error('Detalles del error:', error.response.data);
  }
},


  },
};
</script>
