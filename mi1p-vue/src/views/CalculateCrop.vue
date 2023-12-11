<template>
  <div>
    <h1>Calculadora de planta necesaria para tu crop</h1>

    <div>
      <label for="cropName">Name del crop:</label>
      <input type="text" id="cropName" v-model="nameCrop" />
    </div>

    <div>
      <label for="squareMeters">Meters cuadrados:</label>
      <input type="number" id="squareMeters" v-model="squareMeters" />
    </div>

    <button @click="calcularCrop">Calcular</button>

    <div v-if="resultado">
      <p>{{ resultado.frase }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'CalculateCrop',
  data() {
    return {
      nameCrop: '',
      squareMeters: 0,
      resultado: null,
    };
  },
  methods: {
    calcularCrop() {
      
      axios.post(`${apiUrl}/crops/calculate`, {
        name: this.nameCrop,
        meters: this.squareMeters,
      })
        .then(response => {
          
          const result = response.data.result;
          const frase = `Va a necesitar ${result} plantas de ${this.nameCrop} para plantar sus ${this.squareMeters} meters cuadrados.`;

         
          this.resultado = { frase };
        })
        .catch(error => {
        
          console.error('Error al llamar al servidor:', error);
        });
    },
  },
};
</script>