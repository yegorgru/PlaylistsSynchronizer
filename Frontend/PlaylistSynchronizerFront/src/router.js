import { createRouter, createWebHistory } from 'vue-router';
import LoginComponent from '@/components/Login.vue';
import CreateGroupComponent from '@/components/CreateGroup.vue';
import GroupInfoComponent from '@/components/GroupInfo.vue';
import HomeComponent from '@/components/Home.vue';
import AppComponent from '@/App.vue';

const routes = [
  { path: '/login', component: LoginComponent },
  { path: '/create_group', component: CreateGroupComponent, meta: { requiresAuth: true } },
  { path: '/group_info/:group_id', component: GroupInfoComponent, meta: { requiresAuth: true } },
  { path: '/home', component: HomeComponent, meta: { requiresAuth: true } },
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
    }
    else if (to.path === '/') {
      next('/home');
    } else {
      next();
    }
  } else {
    if (isAuthenticated()) {
      next('/home');
    }
    else {
      next();
    }
  }
});

function isAuthenticated() {
  return localStorage.getItem('access_token') !== null;
}

export default router;
