import { createRouter, createWebHistory } from 'vue-router';
import Index from '../views/Index.vue'; // Asegúrate de importar correctamente

const routes = [
  // ...otras rutas
  {
    path: '/',
    name: 'Index',
    component: Index
  }
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;
