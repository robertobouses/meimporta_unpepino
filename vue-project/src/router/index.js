import { createRouter, createWebHistory } from 'vue-router';
import CreateCultivo from '../views/CreateCultivo.vue';

const routes = [
  {
    path: '/create-cultivo',
    name: 'CreateCultivo',
    component: CreateCultivo,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
