import { createRouter, createWebHistory } from 'vue-router'

import HomeView from '../views/HomeView.vue'
import AboutView from '../views/AboutView.vue';
import CreateCrop from '../views/CreateCrop.vue';
import PrintCrops from '../views/PrintCrops.vue';
import DeleteCrop from '../views/DeleteCrop.vue';
import DeleteAllCrops from '../views/DeleteAllCrops.vue';
import SearchCrop from '../views/SearchCrop.vue';
import CalculateCrop from '../views/CalculateCrop.vue';
import DefineCrop from '../views/DefineCrop.vue';
import ProblemsCrop from '../views/ProblemsCrop.vue';
import ReadCrop from '../views/ReadCrop.vue';

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
      path: '/createCrop',
      name: 'createcrop',
      component: CreateCrop
    },
    {
      path: '/printCrops',
      name: 'printcrops',
      component: PrintCrops
    },
    {
      path: '/deleteCrop',
      name: 'deletecrop',
      component: DeleteCrop
    },
    {
      path: '/deleteAllCrops',
      name: 'deleteallcrops',
      component: DeleteAllCrops
    },

    {
      path: '/searchCrop',
      name: 'searchcrop',
      component: SearchCrop
    },
    {
      path: '/calculateCrop',
      name: 'calculatecrop',
      component: CalculateCrop
    },
    {
      path: '/defineCrop',
      name: 'definecrop',
      component: DefineCrop
    },
    {
      path: '/problemsCrop',
      name: 'problemscrop',
      component: ProblemsCrop
    },
    {
      path: '/readCrop',
      name: 'readcrop',
      component: ReadCrop
    }
  ]
})

export default router
