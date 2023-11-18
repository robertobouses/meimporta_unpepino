<template>
  <div>
    <h1>Calculadora de planta necesaria para tu cultivo</h1>

    <div>
      <label for="cultivoNombre">Nombre del cultivo:</label>
      <input type="text" id="cultivoNombre" v-model="nombreCultivo" />
    </div>

    <div>
      <label for="metrosCuadrados">Metros cuadrados:</label>
      <input type="number" id="metrosCuadrados" v-model="metrosCuadrados" />
    </div>

    <button @click="calcularCultivo">Calcular</button>

    <div v-if="resultado">
      <p>{{ resultado.frase }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'CalculateCultivo',
  data() {
    return {
      nombreCultivo: '',
      metrosCuadrados: 0,
      resultado: null,
    };
  },
  methods: {
    calcularCultivo() {
      
      axios.post('http://localhost:8080/calculateCultivo', {
        nombre: this.nombreCultivo,
        metros: this.metrosCuadrados,
      })
        .then(response => {
          
          const result = response.data.result;
          const frase = `Va a necesitar ${result} plantas de ${this.nombreCultivo} para plantar sus ${this.metrosCuadrados} metros cuadrados.`;

         
          this.resultado = { frase };
        })
        .catch(error => {
        
          console.error('Error al llamar al servidor:', error);
        });
    },
  },
};
</script>