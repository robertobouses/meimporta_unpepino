<template>
  <div>
    <h1>Búsqueda avanzada de crops</h1>
    <h4>Utiliza el criterio que quieras para buscar los crops apropiados:</h4>
    <div>
      <label for="cropName">Name del crop:</label>
      <input type="text" id="cropName" v-model="nameCrop" />
    </div>

    <div>
      <label for="color">Color:</label>
      <input type="text" id="color" v-model="color" />
    </div>


    <div>
      <label for="PlantingDensity">Densidad de plantacion:</label>
      <input type="text" id="plantingDensity" v-model="plantingDensity" />
    </div>

    
    <div>
      <label for="water">Water:</label>
      <input type="text" id="water" v-model="water" />
    </div>

    <div>
      <label for="soil">Soil:</label>
      <input type="text" id="soil" v-model="soil" />
    </div>
    <div>
      <label for="nutrition">Nutrition:</label>
      <input type="text" id="nutrition" v-model="nutrition" />
    </div>
    <div>
      <label for="salinity">Salinity:</label>
      <input type="text" id="salinity" v-model="salinity" />
    </div>

    <div>
      <label for="cycle        ">Cycle        :</label>
      <input type="text" id="cycle        " v-model="cycle        " />
    </div>




    <button @click="buscarCrop">Buscar</button>

    <div v-if="resultados.length > 0">
      <h2>Resultados de la búsqueda:</h2>
      <ul>
        <li v-for="crop in resultados" :key="crop.IdCrop">
          {{ crop.CropInformation .Name }} - {{ crop.CropInformation .Family }}
        </li>
      </ul>
    </div>
    <div v-else>
      <p>No se encontraron resultados para la búsqueda.</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      nameCrop: "",
      // Agrega más propiedades según tus necesidades
      resultados: [],
    };
  },
  methods: {
    buscarCrop() {
      // Realiza la solicitud al backend para buscar crops
      // Aquí puedes utilizar la biblioteca axios o la función fetch
      // Ajusta la URL según la configuración de tu servidor
      const url = "`${apiUrl}/searchCrop`;
      const data = {
        Name: this.nameCrop,
        // Agrega más campos según tus necesidades
      };

      // Realiza la solicitud POST al backend
      fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      })
        .then((response) => response.json())
        .then((data) => {
          // Actualiza los resultados en la propiedad resultados
          this.resultados = data;
        })
        .catch((error) => {
          console.error("Error al buscar crop:", error);
        });
    },
  },
};
</script>