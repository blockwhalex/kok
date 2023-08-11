import { createRouter, createWebHistory } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import type { Action } from 'element-plus'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Login.vue')
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('../views/Register.vue')
    },
    {
      path: '/home',
      name: 'home',
      component: () => import('../views/Home.vue'),
      meta: { requiresAuth:true }
    },
  ]
});

router.beforeEach((to, from, next) => {
  const isLoggedIn = localStorage.getItem('isLoggedIn');
  if (to.meta.requiresAuth && !isLoggedIn) {
    ElMessageBox.alert('请登录后访问该页面', '未登录', {
      type: 'warning',
      confirmButtonText: '确定',
      callback: (action: Action) => {
        next('/login'); // 重定向到登录页面
      },
    })
  } else {
    next();
  }
});

export default router
