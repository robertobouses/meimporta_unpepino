<template>
    <div>
      <h1>Crear Nuevo Cultivo</h1>
      <form @submit.prevent="submitForm">
        <label for="siglas">Siglas:</label>
        <input type="text" id="siglas" v-model="cultivo.siglas" required>
        
        <label for="nombre">Nombre:</label>
        <input type="text" id="nombre" v-model="cultivo.informacioncultivo.nombre" required>
        
        <label for="color">Color:</label>
        <input type="text" id="color" v-model="cultivo.informacioncultivo.color" required>
        
        <label for="familia">Familia:</label>
        <input type="text" id="familia" v-model="cultivo.informacioncultivo.familia" required>
        
        <label for="densidadplantacion">Densidad de Plantación:</label>
        <input type="text" id="densidadplantacion" v-model="cultivo.informacioncultivo.densidadplantacion" required>
        
        <label for="litrostierramaceta">Litros de Tierra por Maceta:</label>
        <input type="number" id="litrostierramaceta" v-model="cultivo.informacioncultivo.litrostierramaceta" required>
        
        <!-- Agrega más campos según la estructura de datos de InformacionCultivo -->
        
        <label for="agua">Agua:</label>
        <input type="text" id="agua" v-model="cultivo.requisitoscultivo.agua" required>
        
        <label for="tierra">Tierra:</label>
        <input type="text" id="tierra" v-model="cultivo.requisitoscultivo.tierra" required>
        
        <label for="nutricion">Nutrición:</label>
        <input type="text" id="nutricion" v-model="cultivo.requisitoscultivo.nutricion" required>
        
        <!-- Continúa agregando campos según la estructura de datos de RequisitosCultivo y otras secciones -->
        
        <button type="submit">Guardar Cultivo</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        cultivo: {
          siglas: '',
          informacioncultivo: {
            nombre: '',
            color: '',
            familia: '',
            densidadplantacion: '',
            litrostierramaceta: 0,
            // Agrega más campos según la estructura de datos de InformacionCultivo
          },
          requisitoscultivo: {
            agua: '',
            tierra: '',
            nutricion: '',
            // Agrega campos de RequisitosCultivo
          },
          fechascultivo: {
            // Agrega campos de FechasCultivo
          },
          frutocultivo: {
            // Agrega campos de FrutoCultivo
          },
          semillacultivo: {
            // Agrega campos de SemillaCultivo
          },
          problemascultivo: {
            // Agrega campos de ProblemasCultivo
          }
        };
      },
      methods: {
        submitForm() {
          // Aquí puedes enviar los datos del cultivo al servidor (backend)
          // Puedes usar Axios u otra librería para hacer la solicitud POST al servidor.
          // Por ejemplo:
          // axios.post('/api/crear-cultivo', this.cultivo)
          //   .then(response => {
          //     // Maneja la respuesta del servidor, por ejemplo, muestra un mensaje de éxito.
          //   })
          //   .catch(error => {
          //     // Maneja los errores, por ejemplo, muestra un mensaje de error.
          //   });
        },
      },
    };
  </script>
  