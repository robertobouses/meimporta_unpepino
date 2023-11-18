<template>
  <div>
    <h1>Crear Nuevo Cultivo</h1>
    <form @submit.prevent="submitForm">
      <label for="siglas">Siglas:</label>
      <input type="text" id="siglas" v-model="cultivo.siglas" required>

      <label for="nombre">Nombre:</label>
      <input type="text" id="nombre" v-model="cultivo.informacioncultivo.nombre" required>

      <label for="color">Color (separados por comas):</label>
      <input type="text" id="color" v-model="cultivo.informacioncultivo.color">

      <label for="familia">Familia:</label>
      <input type="text" id="familia" v-model="cultivo.informacioncultivo.familia" required>

      <label for="densidadplantacion">Densidad de Plantación:</label>
      <input type="text" id="densidadplantacion" v-model="cultivo.informacioncultivo.densidadplantacion">

      <label for="litrostierramaceta">Litros de Tierra por Maceta:</label>
      <input type="number" id="litrostierramaceta" v-model="cultivo.informacioncultivo.litrostierramaceta">

      <label for="asociaciones">Asociaciones (separadas por comas):</label>
      <input type="text" id="asociaciones" v-model="cultivo.informacioncultivo.asociaciones">

      <label for="agua">Agua:</label>
      <input type="text" id="agua" v-model="cultivo.requisitoscultivo.agua">

      <label for="tierra">Tierra:</label>
      <input type="text" id="tierra" v-model="cultivo.requisitoscultivo.tierra">

      <label for="nutricion">Nutrición:</label>
      <input type="text" id="nutricion" v-model="cultivo.requisitoscultivo.nutricion">

      <label for="salinidad">Salinidad:</label>
      <input type="text" id="salinidad" v-model="cultivo.requisitoscultivo.salinidad">

      <label for="ph">pH (separados por comas):</label>
      <input type="number" id="ph" v-model="cultivo.requisitoscultivo.ph">
      <!-- step="0.01" -->

      <label for="clima">Clima (separados por comas):</label>
      <input type="text" id="clima" v-model="cultivo.requisitoscultivo.clima">

      <label for="profundidad">Profundidad:</label>
      <input type="text" id="profundidad" v-model="cultivo.requisitoscultivo.profundidad">

      <label for="siembra">Siembra:</label>
      <input type="text" id="siembra" v-model="cultivo.fechascultivo.siembra">

      <label for="transplante">Transplante:</label>
      <input type="text" id="transplante" v-model="cultivo.fechascultivo.transplante">

      <label for="cosecha">Cosecha:</label>
      <input type="text" id="cosecha" v-model="cultivo.fechascultivo.cosecha">

      <label for="ciclo">Ciclo:</label>
      <input type="text" id="ciclo" v-model="cultivo.fechascultivo.ciclo">

      <label for="produccion">Produccion:</label>
      <input type="text" id="produccion" v-model="cultivo.frutocultivo.produccion">

      <label for="nutrientes">Nutrientes:</label>
      <input type="text" id="nutrientes" v-model="cultivo.frutocultivo.nutrientes">

      <label for="semilla">Semilla:</label>
      <input type="text" id="semilla" v-model="cultivo.semillacultivo.semilla">

      <label for="semillasgramo">SemillasGramo:</label>
      <input type="text" id="semillasgramo" v-model="cultivo.semillacultivo.semillasgramo">

      <label for="vidasemilla">VidaSemilla:</label>
      <input type="text" id="vidasemilla" v-model="cultivo.semillacultivo.vidasemilla">

      <label for="plagas">Plagas:</label>
      <input type="text" id="plagas" v-model="cultivo.problemascultivo.plagas">

      <label for="dificultades">Dificultades:</label>
      <input type="text" id="dificultades" v-model="cultivo.problemascultivo.dificultades">

      <label for="cuidados">Cuidados:</label>
      <input type="text" id="cuidados" v-model="cultivo.problemascultivo.cuidados">

      <label for="miscelanea">Miscelanea:</label>
      <input type="text" id="miscelanea" v-model="cultivo.problemascultivo.miscelanea">

      <button type="button" @click="debugData">Depurar Datos</button>

      <button type="submit">Guardar Cultivo</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      cultivo: {
        siglas: '',
        informacioncultivo: {
          nombre: '',
          color: [],
          familia: '',
          densidadplantacion: '',
          litrostierramaceta: 0,
          asociaciones: [],
        },
        requisitoscultivo: {
          agua: '',
          tierra: '',
          nutricion: '',
          salinidad: '',
          ph: [],
          clima: [],
          profundidad: '',
        },
        fechascultivo: {
          siembra: '',
          transplante: '',
          cosecha: '',
          ciclo: '',
        },
        frutocultivo: {
          produccion: '',
          nutrientes: '',
        },
        semillacultivo: {
          semilla: '',
          semillasgramo: '',
          vidasemilla: '',
        },
        problemascultivo: {
          plagas: '',
          dificultades: '',
          cuidados: '',
          miscelanea: '',
        },
      }, mensaje: 'aaaaaaaaaaaaabbbb',
    };
  },
  methods: {
    debugData() {
      console.log('Depurar Datos');
      console.log('Cultivo:', this.cultivo);
      console.log('Cultivo JSON:', JSON.stringify(this.cultivo));
    },
    async submitForm() {
      try {
        console.log('Submit Form Called');
        console.log('Cultivo:', this.cultivo);


// // Procesamiento del color
// if (typeof this.cultivo.informacioncultivo.color === 'string' && this.cultivo.informacioncultivo.color.trim() !== '') {
//   this.cultivo.informacioncultivo.color = this.cultivo.informacioncultivo.color.split(',').map(color => color.trim());
// } else {
//   // Set it to an empty array if it's an empty string
//   this.cultivo.informacioncultivo.color = [];
// }

// // Procesamiento de asociaciones
// if (typeof this.cultivo.informacioncultivo.asociaciones === 'string' && this.cultivo.informacioncultivo.asociaciones.trim() !== '') {
//   this.cultivo.informacioncultivo.asociaciones = this.cultivo.informacioncultivo.asociaciones.split(',').map(asociacion => asociacion.trim());
// } else {
//   // Set it to an empty array if it's an empty string
//   this.cultivo.informacioncultivo.asociaciones = [];
// }

// // Asegurarse de que ph sea un array
// if (!Array.isArray(this.cultivo.requisitoscultivo.ph)) {
//   this.cultivo.requisitoscultivo.ph = [this.cultivo.requisitoscultivo.ph];
// }

// // Convertir ph a números si es un array
// if (Array.isArray(this.cultivo.requisitoscultivo.ph)) {
//   this.cultivo.requisitoscultivo.ph = this.cultivo.requisitoscultivo.ph.map(item => parseFloat(item));
// }

// // Procesamiento de clima
// if (this.cultivo.requisitoscultivo.clima) {
//   console.log("this.cultivo.requisitoscultivo.clima",  this.cultivo.requisitoscultivo.clima)
//   this.cultivo.requisitoscultivo.clima = this.cultivo.requisitoscultivo.clima.split(',').map(clima => clima.trim());
// }


        // // Procesamiento del color
        // if (this.cultivo.informacioncultivo.color) {
        //   this.cultivo.informacioncultivo.color = this.cultivo.informacioncultivo.color.split(',').map(color => color.trim());
        // }

        // // Procesamiento de asociaciones
        // if (this.cultivo.informacioncultivo.asociaciones) {
        //   this.cultivo.informacioncultivo.asociaciones = this.cultivo.informacioncultivo.asociaciones.split(',').map(asociacion => asociacion.trim());
        // }

        // // Asegurarse de que ph sea un array
        // if (!Array.isArray(this.cultivo.requisitoscultivo.ph)) {
        //   this.cultivo.requisitoscultivo.ph = [this.cultivo.requisitoscultivo.ph];
        // }

        // // Convertir ph a números si es un array
        // if (Array.isArray(this.cultivo.requisitoscultivo.ph)) {
        //   this.cultivo.requisitoscultivo.ph = this.cultivo.requisitoscultivo.ph.map(item => parseFloat(item));
        // }

        console.log('Cultivo antes de enviar:', JSON.stringify(this.cultivo));

        // Enviar datos a la API
        const response = await axios.post('http://localhost:8080/createCultivo', this.cultivo);
        console.log('Respuesta del servidor:', response);

        this.mensaje = 'Cultivo creado exitosamente';
      } catch (error) {
        console.error('Error al crear el cultivo:', error);

        this.mensaje = 'Error al crear el cultivo';

        // Manejar errores de respuesta
        if (error.response) {
          this.mensaje = error.response.status === 400 ? 'Error de validación. Por favor, verifica los datos del formulario.' : 'Error del servidor. Inténtalo de nuevo más tarde.';
        }
      }
    },
  },
};
</script>