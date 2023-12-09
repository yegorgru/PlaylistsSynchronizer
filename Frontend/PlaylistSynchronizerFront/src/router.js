import { createRouter, createWebHistory } from 'vue-router';
import LoginComponent from '@/components/Login.vue';
import AppComponent from '@/App.vue';

const routes = [
  { path: '/login', component: LoginComponent },
  { path: '/', component: AppComponent, meta: { requiresAuth: true } },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    if (!isAuthenticated()) {
      next('/login');
    } else {
      next();
    }
  } else {
    next();
  }
});

function isAuthenticated() {
  return localStorage.getItem('access_token') !== null;
}

export default router;
