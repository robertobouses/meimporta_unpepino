<template>
  <div class="about">
    <h1>Selecciona una opción</h1>
    <router-link to="/createCultivo">
      <button type="button" class="btn btn-outline-success">Crear un Cultivo</button>
    </router-link>
    <div>
      <router-link to="/printCultivos">
        <button type="button" class="btn btn-outline-success">Ver todos los Cultivos</button>
      </router-link>
    </div>
    <div>
      <router-link to="/deleteCultivo">
        <button type="button" class="btn btn-outline-success">Eliminar un Cultivo</button>
      </router-link>
    </div>
    <div>
      <router-link to="/deleteAllCultivos">
        <button type="button" class="btn btn-outline-success">Eliminar todos los Cultivos</button>
      </router-link>
    </div>
    <div>
      <router-link to="/searchCultivo">
        <button type="button" class="btn btn-outline-success">Buscar un Cultivo</button>
      </router-link>
    </div>
    <div>
      <router-link to="/calculateCultivo">
        <button type="button" class="btn btn-outline-success">Calculador de Plantas</button>
      </router-link>
    </div>
    <div>
      <router-link to="/defineCultivo">
        <button type="button" class="btn btn-outline-success">Define tu propio Cultivo</button>
      </router-link>
    </div>
    <div>
      <router-link to="/problemsCultivo">
        <button type="button" class="btn btn-outline-success">Plagas y problemas</button>
      </router-link>
    </div>
    <div>
      <router-link to="/readCultivo">
        <button type="button" class="btn btn-outline-success">Lecturas</button>
      </router-link>
    </div>
  </div>
</template>
