import { createRouter, createWebHashHistory } from 'vue-router';
import Home from './views/Home.vue';
import About from './views/About.vue';
import { HomeIcon, UserIcon } from '@heroicons/vue/24/solid';

export interface RouteConfig {
  name: string;
  path: string;
  component: any;
  icon: any;
}

export const routes: RouteConfig[] = [
  { path: '/', component: Home, name: 'Home', icon: HomeIcon },
  { path: '/about', component: About, name: 'About', icon: UserIcon },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
