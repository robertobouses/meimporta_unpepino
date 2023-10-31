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
        
        <label for="asociaciones">Asociaciones:</label>
        <input type="text" id="asociaciones" v-model="cultivo.informacioncultivo.asociaciones" required>
        
        <label for="agua">Agua:</label>
        <input type="text" id="agua" v-model="cultivo.requisitoscultivo.agua" required>
        
        <label for="tierra">Tierra:</label>
        <input type="text" id="tierra" v-model="cultivo.requisitoscultivo.tierra" required>
        
        <label for="nutricion">Nutrición:</label>
        <input type="text" id="nutricion" v-model="cultivo.requisitoscultivo.nutricion" required>
       
        <label for="salinidad">Salinidad:</label>
        <input type="text" id="salinidad" v-model="cultivo.requisitoscultivo.salinidad" required>
       
        <label for="ph">Ph:</label>
        <input type="number" id="ph" v-model="cultivo.requisitoscultivo.ph" required>
       
        <label for="clima">Clima:</label>
        <input type="text" id="clima" v-model="cultivo.requisitoscultivo.clima" required>
       

        <label for="profundidad">Profundidad:</label>
        <input type="text" id="profundidad" v-model="cultivo.requisitoscultivo.profundidad" required>
       
        <label for="siembra">Siembra:</label>
        <input type="text" id="siembra" v-model="cultivo.fechascultivo.siembra" required>
       
        <label for="transplante">Transplante:</label>
        <input type="transplante" id="transplante" v-model="cultivo.fechascultivo.transplante" required>
      

        <label for="cosecha">Cosecha:</label>
        <input type="text" id="cosecha" v-model="cultivo.fechascultivo.cosecha" required>
      
        <label for="ciclo">Ciclo:</label>
        <input type="text" id="ciclo" v-model="cultivo.fechascultivo.ciclo" required>
      
        <label for="produccion">Produccion:</label>
        <input type="text" id="produccion" v-model="cultivo.frutocultivo.produccion" required>
      
        <label for="nutrientes">Nutrientes:</label>
        <input type="text" id="nutrientes" v-model="cultivo.frutocultivo.nutrientes" required>
      
        <label for="semilla">Semilla:</label>
        <input type="text" id="semilla" v-model="cultivo.semillacultivo.semilla" required>
       
        <label for="semillasgramo">SemillasGramo:</label>
        <input type="text" id="semillagramo" v-model="cultivo.semillacultivo.semillasgramo" required>
        

        <label for="vidasemilla">VidaSemilla:</label>
        <input type="text" id="vidasemilla" v-model="cultivo.semillacultivo.vidasemilla" required>
        
        <label for="plagas">Plagas:</label>
        <input type="text" id="plagas" v-model="cultivo.problemascultivo.plagas" required>

        <label for="dificultades">Dificultades:</label>
        <input type="text" id="dificultades" v-model="cultivo.problemascultivo.dificultades" required>

        <label for="cuidados">Cuidados:</label>
        <input type="text" id="cuidados" v-model="cultivo.problemascultivo.cuidados" required>

        <label for="miscelanea">Miscelanea:</label>
        <input type="text" id="miscelanea" v-model="cultivo.problemascultivo.miscelanea" required>


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
            color: '',
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
            ph: 0,
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
          }
      },
    },
    methods: {
    submitForm() {
      // Aquí puedes enviar los datos del cultivo al servidor (backend)
      axios.post('http://localhost:8080/cultivo', this.cultivo)
        .then(response => {
          // Maneja la respuesta del servidor, por ejemplo, muestra un mensaje de éxito.
          console.log('Cultivo creado exitosamente');
        })
        .catch(error => {
          // Maneja los errores, por ejemplo, muestra un mensaje de error.
          console.error('Error al crear el cultivo:', error);
        });
    },
  },
};
</script>