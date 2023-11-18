<template>
  <div>
    <h1>Búsqueda avanzada de cultivos</h1>
    <h4>Utiliza el criterio que quieras para buscar los cultivos apropiados:</h4>
    <div>
      <label for="cultivoNombre">Nombre del cultivo:</label>
      <input type="text" id="cultivoNombre" v-model="nombreCultivo" />
    </div>

    <div>
      <label for="color">Color:</label>
      <input type="text" id="color" v-model="color" />
    </div>


    <div>
      <label for="DensidadPlantacion">Densidad de plantacion:</label>
      <input type="text" id="densidadPlantacion" v-model="densidadPlantacion" />
    </div>

    
    <div>
      <label for="agua">Agua:</label>
      <input type="text" id="agua" v-model="agua" />
    </div>

    <div>
      <label for="tierra">Tierra:</label>
      <input type="text" id="tierra" v-model="tierra" />
    </div>
    <div>
      <label for="nutricion">Nutricion:</label>
      <input type="text" id="nutricion" v-model="nutricion" />
    </div>
    <div>
      <label for="salinidad">Salinidad:</label>
      <input type="text" id="salinidad" v-model="salinidad" />
    </div>

    <div>
      <label for="ciclo">Ciclo:</label>
      <input type="text" id="ciclo" v-model="ciclo" />
    </div>




    <button @click="buscarCultivo">Buscar</button>

    <div v-if="resultados.length > 0">
      <h2>Resultados de la búsqueda:</h2>
      <ul>
        <li v-for="cultivo in resultados" :key="cultivo.IdCultivo">
          {{ cultivo.InformacionCultivo.Nombre }} - {{ cultivo.InformacionCultivo.Familia }}
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
      nombreCultivo: "",
      // Agrega más propiedades según tus necesidades
      resultados: [],
    };
  },
  methods: {
    buscarCultivo() {
      // Realiza la solicitud al backend para buscar cultivos
      // Aquí puedes utilizar la biblioteca axios o la función fetch
      // Ajusta la URL según la configuración de tu servidor
      const url = "http://localhost:8080/searchCultivo";
      const data = {
        Nombre: this.nombreCultivo,
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
          console.error("Error al buscar cultivo:", error);
        });
    },
  },
};
</script>