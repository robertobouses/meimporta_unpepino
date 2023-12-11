<template>
    <div>
      <h1>Crear Nuevo Crop</h1>
      <form @submit.prevent="submitForm">
        <label for="abbreviation">Abbreviation:</label>
        <input type="text" id="abbreviation" v-model="crop.abbreviation" required>
        
        <label for="name">Name:</label>
        <input type="text" id="name" v-model="crop.cropinformation .name" required>
        
        <label for="color">Color:</label>
        <input type="text" id="color" v-model="crop.cropinformation .color" required>
        
        <label for="family">Family:</label>
        <input type="text" id="family" v-model="crop.cropinformation .family" required>
        
        <label for="plantingdensity">Densidad de Plantación:</label>
        <input type="text" id="plantingdensity" v-model="crop.cropinformation .plantingdensity" required>
        
        <label for="literspottingsoil  ">Litros de Soil por Maceta:</label>
        <input type="number" id="literspottingsoil  " v-model="crop.cropinformation .literspottingsoil  " required>
        
        <label for="associations       ">Associations       :</label>
        <input type="text" id="associations       " v-model="crop.cropinformation .associations       " required>
        
        <label for="water">Water:</label>
        <input type="text" id="water" v-model="crop.croprequirements .water" required>
        
        <label for="soil">Soil:</label>
        <input type="text" id="soil" v-model="crop.croprequirements .soil" required>
        
        <label for="nutrition">Nutrición:</label>
        <input type="text" id="nutrition" v-model="crop.croprequirements .nutrition" required>
       
        <label for="salinity">Salinity:</label>
        <input type="text" id="salinity" v-model="crop.croprequirements .salinity" required>
       
        <label for="ph">Ph:</label>
        <input type="number" id="ph" v-model="crop.croprequirements .ph" required>
       
        <label for="climate">Climate:</label>
        <input type="text" id="climate" v-model="crop.croprequirements .climate" required>
       

        <label for="depth">Depth:</label>
        <input type="text" id="depth" v-model="crop.croprequirements .depth" required>
       
        <label for="planting     ">Planting     :</label>
        <input type="text" id="planting     " v-model="crop.cropdates.planting     " required>
       
        <label for="transplant   ">Transplant   :</label>
        <input type="transplant   " id="transplant   " v-model="crop.cropdates.transplant   " required>
      

        <label for="harvest      ">Harvest      :</label>
        <input type="text" id="harvest      " v-model="crop.cropdates.harvest      " required>
      
        <label for="cycle        ">Cycle        :</label>
        <input type="text" id="cycle        " v-model="crop.cropdates.cycle        " required>
      
        <label for="production">Production:</label>
        <input type="text" id="production" v-model="crop.cropfruit.production" required>
      
        <label for="nutrients  ">Nutrients  :</label>
        <input type="text" id="nutrients  " v-model="crop.cropfruit.nutrients  " required>
      
        <label for="seed">Seed:</label>
        <input type="text" id="seed" v-model="crop.cropseed.seed" required>
       
        <label for="seedspergram">SeedsPerGram  :</label>
        <input type="text" id="seedgramo" v-model="crop.cropseed.seedspergram  " required>
        

        <label for="seedlifespan">SeedLifespan:</label>
        <input type="text" id="seedlifespan" v-model="crop.cropseed.seedlifespan" required>
        
        <label for="pests">Pests:</label>
        <input type="text" id="pests" v-model="crop.cropissues.pests" required>

        <label for="difficulties ">Difficulties :</label>
        <input type="text" id="difficulties " v-model="crop.cropissues.difficulties " required>

        <label for="care">Care:</label>
        <input type="text" id="care" v-model="crop.cropissues.care" required>

        <label for="miscellaneous">Miscellaneous:</label>
        <input type="text" id="miscellaneous" v-model="crop.cropissues.miscellaneous" required>


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
            color: '',
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
            ph: 0,
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
          }
      },
    },
    methods: {
    submitForm() {
      // Aquí puedes enviar los datos del crop al servidor (backend)
      axios.post('http://localhost:8080/crop', this.crop)
        .then(response => {
          // Maneja la respuesta del servidor, por ejemplo, muestra un mensaje de éxito.
          console.log('Crop creado exitosamente');
        })
        .catch(error => {
          // Maneja los errores, por ejemplo, muestra un mensaje de error.
          console.error('Error al crear el crop:', error);
        });
    },
  },
};
</script>