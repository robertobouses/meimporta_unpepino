<template>
  <div>
    <h1>Lecturas</h1>
    <ul>
      <li v-for="lectura in lecturas" :key="lectura.data.id">{{ lectura.data.title }}</li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'ReadCrop',
  data() {
    return {
      lecturas: [], // Aquí almacenarás las lecturas
    };
  },
  mounted() {
    axios.get('https://www.reddit.com/r/pics.json')
      .then(response => {
        // Extraer las lecturas de la respuesta
        this.lecturas = response.data.data.children;
      })
      .catch(error => {
        console.error(error);
      });
  },
};
</script>
