import { createRouter, createWebHistory } from 'vue-router'

import HomeView from '../views/HomeView.vue'
import AboutView from '../views/AboutView.vue';
import CreateCrop from '../views/CreateCrop.vue';
import PrintCrops from '../views/PrintCrops.vue';
import DeleteCrop from '../views/DeleteCrop.vue';
import DeleteAllCrops from '../views/DeleteAllCrops.vue';
import SearchCrop from '../views/SearchCrop.vue';
import CalculateCrop from '../views/CalculateCrop.vue';
import DefineMyCrop from '../views/DefineMyCrop.vue';
import IssuesCrop from '../views/IssuesCrop.vue';


import GetCalendary from '../views/GetCalendary.vue';
import CreateProvinces from '../views/CreateProvinces.vue';
import DefineMyFields from '../views/DefineMyFields.vue';
import GetMyCrop from '../views/GetMyCrop.vue';
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
      path: '/issuesCrop',
      name: 'issuescrop',
      component: IssuesCrop
   },
   
    {
      path: '/calculateCrop',
      name: 'calculatecrop',
      component: CalculateCrop
    },
   
    
 {
      path: '/searchCrop',
      name: 'searchcrop',
      component: SearchCrop
    },
    {
      path: '/getCalendary',
      name: 'getcalendary',
      component: GetCalendary
    },


    {
      path: '/createProvinces',
      name: 'createprovinces',
      component: CreateProvinces
    },


    {
      path: '/defineMyFields',
      name: 'definemyfields',
      component: DefineMyFields
    },
 {
      path: '/defineMyCrop',
      name: 'definemycrop',
      component: DefineMyCrop
    },

    {
      path: '/getMyCrop',
      name: 'getmycrop',
      component: GetMyCrop
    },




    {
      path: '/readCrop',
      name: 'readcrop',
      component: ReadCrop
    }
  ]
})

export default router
