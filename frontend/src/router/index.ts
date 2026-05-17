import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/authStore'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/Auth/LoginView.vue'),
      meta: { title: 'Login', layout: 'empty' }
    },
    {
      path: '/',
      name: 'dashboard',
      component: () => import('../views/Dashboard/DashboardView.vue'),
      meta: { title: 'Dashboard' }
    },
    {
      path: '/clientes',
      name: 'clientes',
      component: () => import('../views/Clientes/ClientesList.vue'),
      meta: { title: 'Clientes' }
    },
    {
      path: '/analistas',
      name: 'analistas',
      component: () => import('../views/Analistas/AnalistasView.vue'),
      meta: { title: 'Analistas', requiresAdmin: true }
    },
    {
      path: '/projetos',
      name: 'projetos',
      component: () => import('../views/Projetos/ProjetosView.vue'),
      meta: { title: 'Projetos' }
    },
    {
      path: '/reunioes',
      name: 'reunioes',
      component: () => import('../views/Reunioes/ReunioesView.vue'),
      meta: { title: 'Reuniões' }
    },
    {
      path: '/relatorios',
      name: 'relatorios',
      component: () => import('../views/Relatorios/RelatoriosView.vue'),
      meta: { title: 'Relatórios' }
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  const isAuthenticated = authStore.isAuthenticated
  const isAdmin = authStore.user?.cargo === 'Admin'
  
  if (to.name !== 'login' && !isAuthenticated) {
    next({ name: 'login' })
  } else if (to.name === 'login' && isAuthenticated) {
    next({ name: 'dashboard' })
  } else if (to.meta.requiresAdmin && !isAdmin) {
    next({ name: 'dashboard' }) // Redireciona se não for admin
  } else {
    next()
  }
})

export default router
