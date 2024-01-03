<template>
  <div>
    <h1>Create Province</h1>
    <form @submit.prevent="crearProvincia">
      <label for="name">Name:</label>
      <input type="text" id="name" v-model="province.name" required />

      <label for="climate">Climate:</label>
      <input type="text" id="climate" v-model="province.climate" required />

      <button type="submit">Crear Provincia</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios'; 
import { apiUrl } from '@/path-to-config/config.js';
export default {
  data() {
    return {
      province: {
        name: '',
        climate: '',
      },
    };
  },
  methods: {
    async crearProvincia() {
  try {
    const response = await axios.post(`${apiUrl}/provinces`, this.province, {
      headers: {
        'Content-Type': 'application/json',
      },
    });

    console.log(response); // Muestra la respuesta completa del servidor

    if (response.status === 200) {
      console.log('Provincia creada exitosamente');

      // Accede a los datos específicos de la provincia creada y al mensaje
      const { province, mensaje } = response.data;
      console.log('Provincia creada:', province);
      console.log('Mensaje:', mensaje);
    } else {
      console.error('Error al crear la provincia:', response.statusText);
    }
  } catch (error) {
    console.error('Error al enviar la solicitud:', error);
  }
}
,



  },
};
</script>
