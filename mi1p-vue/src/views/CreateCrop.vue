<template>
  <div>
    <h1>Crear Nuevo Crop</h1>
    <form @submit.prevent="submitForm">
      <label for="abbreviation">Abbreviation:</label>
      <input type="text" id="abbreviation" v-model="crop.abbreviation" required>

      <label for="name">Name:</label>
      <input type="text" id="name" v-model="crop.cropinformation .name" required>

      <label for="color">Color (separados por comas):</label>
      <input type="text" id="color" v-model="crop.cropinformation .color">

      <label for="family">Family:</label>
      <input type="text" id="family" v-model="crop.cropinformation .family" required>

      <label for="plantingdensity">Densidad de Plantación:</label>
      <input type="text" id="plantingdensity" v-model="crop.cropinformation .plantingdensity">

      <label for="literspottingsoil  ">Litros de Soil por Maceta:</label>
      <input type="number" id="literspottingsoil  " v-model="crop.cropinformation .literspottingsoil  ">

      <label for="associations       ">Associations        (separadas por comas):</label>
      <input type="text" id="associations       " v-model="crop.cropinformation .associations       ">

      <label for="water">Water:</label>
      <input type="text" id="water" v-model="crop.croprequirements .water">

      <label for="soil">Soil:</label>
      <input type="text" id="soil" v-model="crop.croprequirements .soil">

      <label for="nutrition">Nutrición:</label>
      <input type="text" id="nutrition" v-model="crop.croprequirements .nutrition">

      <label for="salinity">Salinity:</label>
      <input type="text" id="salinity" v-model="crop.croprequirements .salinity">

      <label for="ph">pH (separados por comas):</label>
      <input type="number" id="ph" v-model="crop.croprequirements .ph">
      <!-- step="0.01" -->

      <label for="climate">Climate (separados por comas):</label>
      <input type="text" id="climate" v-model="crop.croprequirements .climate">

      <label for="depth">Depth:</label>
      <input type="text" id="depth" v-model="crop.croprequirements .depth">

      <label for="planting     ">Planting     :</label>
      <input type="text" id="planting     " v-model="crop.cropdates.planting     ">

      <label for="transplant   ">Transplant   :</label>
      <input type="text" id="transplant   " v-model="crop.cropdates.transplant   ">

      <label for="harvest      ">Harvest      :</label>
      <input type="text" id="harvest      " v-model="crop.cropdates.harvest      ">

      <label for="cycle        ">Cycle        :</label>
      <input type="text" id="cycle        " v-model="crop.cropdates.cycle        ">

      <label for="production">Production:</label>
      <input type="text" id="production" v-model="crop.cropfruit.production">

      <label for="nutrients  ">Nutrients  :</label>
      <input type="text" id="nutrients  " v-model="crop.cropfruit.nutrients  ">

      <label for="seed">Seed:</label>
      <input type="text" id="seed" v-model="crop.cropseed.seed">

      <label for="seedspergram  ">SeedsPerGram  :</label>
      <input type="text" id="seedspergram  " v-model="crop.cropseed.seedspergram  ">

      <label for="seedlifespan">SeedLifespan:</label>
      <input type="text" id="seedlifespan" v-model="crop.cropseed.seedlifespan">

      <label for="pests       ">Pests       :</label>
      <input type="text" id="pests       " v-model="crop.cropissues.pests       ">

      <label for="difficulties ">Difficulties :</label>
      <input type="text" id="difficulties " v-model="crop.cropissues.difficulties ">

      <label for="care">Care:</label>
      <input type="text" id="care" v-model="crop.cropissues.care">

      <label for="miscellaneous">Miscellaneous:</label>
      <input type="text" id="miscellaneous" v-model="crop.cropissues.miscellaneous">

      <button type="button" @click="debugData">Depurar Datos</button>

      <button type="submit">Guardar Crop</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      crop: {
        abbreviation: '',
        cropinformation : {
          name: '',
          color: [],
          family: '',
          plantingdensity: '',
          literspottingsoil  : 0,
          associations       : [],
        },
        croprequirements : {
          water: '',
          soil: '',
          nutrition: '',
          salinity: '',
          ph: [],
          climate: [],
          depth: '',
        },
        cropdates: {
          planting     : '',
          transplant   : '',
          harvest      : '',
          cycle        : '',
        },
        cropfruit: {
          production: '',
          nutrients  : '',
        },
        cropseed: {
          seed: '',
          seedspergram  : '',
          seedlifespan: '',
        },
        cropissues: {
          pests       : '',
          difficulties : '',
          care: '',
          miscellaneous: '',
        },
      }, mensaje: 'aaaaaaaaaaaaabbbb',
    };
  },
  methods: {
    debugData() {
      console.log('Depurar Datos');
      console.log('Crop:', this.crop);
      console.log('Crop JSON:', JSON.stringify(this.crop));
    },
    async submitForm() {
      try {
        console.log('Submit Form Called');
        console.log('Crop:', this.crop);


// // Procesamiento del color
// if (typeof this.crop.cropinformation .color === 'string' && this.crop.cropinformation .color.trim() !== '') {
//   this.crop.cropinformation .color = this.crop.cropinformation .color.split(',').map(color => color.trim());
// } else {
//   // Set it to an empty array if it's an empty string
//   this.crop.cropinformation .color = [];
// }

// // Procesamiento de associations       
// if (typeof this.crop.cropinformation .associations        === 'string' && this.crop.cropinformation .associations       .trim() !== '') {
//   this.crop.cropinformation .associations        = this.crop.cropinformation .associations       .split(',').map(asociacion => asociacion.trim());
// } else {
//   // Set it to an empty array if it's an empty string
//   this.crop.cropinformation .associations        = [];
// }

// // Asegurarse de que ph sea un array
// if (!Array.isArray(this.crop.croprequirements .ph)) {
//   this.crop.croprequirements .ph = [this.crop.croprequirements .ph];
// }

// // Convertir ph a números si es un array
// if (Array.isArray(this.crop.croprequirements .ph)) {
//   this.crop.croprequirements .ph = this.crop.croprequirements .ph.map(item => parseFloat(item));
// }

// // Procesamiento de climate
// if (this.crop.croprequirements .climate) {
//   console.log("this.crop.croprequirements .climate",  this.crop.croprequirements .climate)
//   this.crop.croprequirements .climate = this.crop.croprequirements .climate.split(',').map(climate => climate.trim());
// }


        // // Procesamiento del color
        // if (this.crop.cropinformation .color) {
        //   this.crop.cropinformation .color = this.crop.cropinformation .color.split(',').map(color => color.trim());
        // }

        // // Procesamiento de associations       
        // if (this.crop.cropinformation .associations       ) {
        //   this.crop.cropinformation .associations        = this.crop.cropinformation .associations       .split(',').map(asociacion => asociacion.trim());
        // }

        // // Asegurarse de que ph sea un array
        // if (!Array.isArray(this.crop.croprequirements .ph)) {
        //   this.crop.croprequirements .ph = [this.crop.croprequirements .ph];
        // }

        // // Convertir ph a números si es un array
        // if (Array.isArray(this.crop.croprequirements .ph)) {
        //   this.crop.croprequirements .ph = this.crop.croprequirements .ph.map(item => parseFloat(item));
        // }

        console.log('Crop antes de enviar:', JSON.stringify(this.crop));

        // Enviar datos a la API
        const response = await axios.post(`${apiUrl}/crops`, this.crop);
        console.log('Respuesta del servidor:', response);

        this.mensaje = 'Crop creado exitosamente';
      } catch (error) {
        console.error('Error al crear el crop:', error);

        this.mensaje = 'Error al crear el crop';

        // Manejar errores de respuesta
        if (error.response) {
          this.mensaje = error.response.status === 400 ? 'Error de validación. Por favor, verifica los datos del formulario.' : 'Error del servidor. Inténtalo de nuevo más tarde.';
        }
      }
    },
  },
};
</script>