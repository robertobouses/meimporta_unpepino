import { createRouter, createWebHistory } from 'vue-router';
import Home from './views/Home.vue';
import CreateCultivo from './views/CreateCultivo.vue';

const routes = [
  { path: '/', component: Home },
  { path: '/crear-cultivo', component: CreateCultivo },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
