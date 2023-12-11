<template>
  <div class="about">
    <h1>Selecciona una opción</h1>
    <router-link to="/createCrop">
      <button type="button" class="btn btn-outline-success">Crear un Crop</button>
    </router-link>
    <div>
      <router-link to="/printCrops">
        <button type="button" class="btn btn-outline-success">Ver todos los Crops</button>
      </router-link>
    </div>
    <div>
      <router-link to="/deleteCrop">
        <button type="button" class="btn btn-outline-success">Eliminar un Crop</button>
      </router-link>
    </div>
    <div>
      <router-link to="/deleteAllCrops">
        <button type="button" class="btn btn-outline-success">Eliminar todos los Crops</button>
      </router-link>
    </div>
    <div>
      <router-link to="/searchCrop">
        <button type="button" class="btn btn-outline-success">Buscar un Crop</button>
      </router-link>
    </div>
    <div>
      <router-link to="/calculateCrop">
        <button type="button" class="btn btn-outline-success">Calculador de Plantas</button>
      </router-link>
    </div>
    <div>
      <router-link to="/defineCrop">
        <button type="button" class="btn btn-outline-success">Define tu propio Crop</button>
      </router-link>
    </div>
    <div>
      <router-link to="/problemsCrop">
        <button type="button" class="btn btn-outline-success">Pests y problemas</button>
      </router-link>
    </div>
    <div>
      <router-link to="/readCrop">
        <button type="button" class="btn btn-outline-success">Lecturas</button>
      </router-link>
    </div>
  </div>
</template>
