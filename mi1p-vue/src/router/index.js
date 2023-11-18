import { createRouter, createWebHistory } from 'vue-router'

import HomeView from '../views/HomeView.vue'
import AboutView from '../views/AboutView.vue';
import CreateCultivo from '../views/CreateCultivo.vue';
import PrintCultivos from '../views/PrintCultivos.vue';
import DeleteCultivo from '../views/DeleteCultivo.vue';
import DeleteAllCultivos from '../views/DeleteAllCultivos.vue';
import SearchCultivo from '../views/SearchCultivo.vue';
import CalculateCultivo from '../views/CalculateCultivo.vue';
import DefineCultivo from '../views/DefineCultivo.vue';
import ProblemsCultivo from '../views/ProblemsCultivo.vue';
import ReadCultivo from '../views/ReadCultivo.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about',
      name: 'about',
      component: AboutView
    },
    {
      path: '/createCultivo',
      name: 'createcultivo',
      component: CreateCultivo
    },
    {
      path: '/printCultivos',
      name: 'printcultivos',
      component: PrintCultivos
    },
    {
      path: '/deleteCultivo',
      name: 'deletecultivo',
      component: DeleteCultivo
    },
    {
      path: '/deleteAllCultivos',
      name: 'deleteallcultivos',
      component: DeleteAllCultivos
    },

    {
      path: '/searchCultivo',
      name: 'searchcultivo',
      component: SearchCultivo
    },
    {
      path: '/calculateCultivo',
      name: 'calculatecultivo',
      component: CalculateCultivo
    },
    {
      path: '/defineCultivo',
      name: 'definecultivo',
      component: DefineCultivo
    },
    {
      path: '/problemsCultivo',
      name: 'problemscultivo',
      component: ProblemsCultivo
    },
    {
      path: '/readCultivo',
      name: 'readcultivo',
      component: ReadCultivo
    }
  ]
})

export default router
