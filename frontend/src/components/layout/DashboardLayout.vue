<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useToastStore } from '../../stores/toastStore'
import { 
  LayoutGrid, 
  UserCircle, 
  Package, 
  Calendar, 
  BarChart3,
  Settings, 
  LogOut,
  Menu,
  X,
  User,
  Shield,
  CreditCard
} from 'lucide-vue-next'

const router = useRouter()
const toastStore = useToastStore()
const isProfileOpen = ref(false)
const isMobileMenuOpen = ref(false)

const menuItems = [
  { name: 'Dashboard', icon: LayoutGrid, path: '/' },
  { name: 'Clientes', icon: UserCircle, path: '/clientes' },
  { name: 'Projetos', icon: Package, path: '/projetos' },
  { name: 'Reuniões', icon: Calendar, path: '/reunioes' },
  { name: 'Relatórios', icon: BarChart3, path: '/relatorios' },
]

const toggleProfile = () => {
  isProfileOpen.value = !isProfileOpen.value
}

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}

const handleLogout = () => {
  localStorage.removeItem('isAuthenticated')
  toastStore.info('Sessão encerrada.')
  router.push('/login')
}

const handleSettings = () => {
  toastStore.info('Configurações em desenvolvimento...')
}
</script>

<template>
  <div class="min-h-screen bg-[#F8FAFC] flex">
    <!-- Overlay for mobile menu -->
    <Transition
      enter-active-class="transition duration-300 ease-out"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition duration-200 ease-in"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div 
        v-if="isMobileMenuOpen" 
        @click="toggleMobileMenu"
        class="fixed inset-0 bg-slate-900/50 backdrop-blur-sm z-40 lg:hidden"
      ></div>
    </Transition>

    <!-- Sidebar -->
    <aside 
      class="bg-brand-sidebar flex flex-col items-center py-6 fixed h-full z-50 transition-all duration-300 lg:translate-x-0 w-[72px]"
      :class="[isMobileMenuOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0']"
    >
      <!-- Logo -->
      <div class="mb-10">
        <div class="w-10 h-10 border-4 border-brand-cyan rounded-full flex items-center justify-center">
          <div class="w-2 h-2 bg-brand-cyan rounded-full"></div>
        </div>
      </div>

      <!-- Nav Items -->
      <nav class="flex-1 flex flex-col gap-6">
        <router-link 
          v-for="item in menuItems" 
          :key="item.path"
          :to="item.path"
          @click="isMobileMenuOpen = false"
          class="p-3 rounded-xl transition-all duration-200 group relative flex items-center justify-center"
          :class="[$route.path === item.path ? 'bg-brand-blue/20 text-brand-blue shadow-[0_0_15px_rgba(59,130,246,0.3)]' : 'text-slate-400 hover:text-white hover:bg-white/5']"
        >
          <component :is="item.icon" class="w-6 h-6" />
          
          <!-- Tooltip / Label lateral -->
          <div class="absolute left-full ml-4 px-3 py-1 bg-slate-900 text-white text-xs rounded opacity-0 group-hover:opacity-100 pointer-events-none transition-opacity whitespace-nowrap z-50">
            {{ item.name }}
          </div>
        </router-link>
      </nav>

      <!-- Bottom Items -->
      <div class="mt-auto flex flex-col gap-6">
        <button 
          @click="handleSettings"
          class="p-3 text-slate-400 hover:text-white hover:bg-white/5 rounded-xl transition-all"
        >
          <Settings class="w-6 h-6" />
        </button>
        <button 
          @click="handleLogout"
          class="p-3 text-rose-500 hover:text-rose-400 hover:bg-rose-500/10 rounded-xl transition-all"
          title="Sair"
        >
          <LogOut class="w-6 h-6" />
        </button>
      </div>
    </aside>

    <!-- Main Content -->
    <div class="flex-1 lg:ml-[72px] flex flex-col min-w-0">
      <header class="h-20 bg-white/80 backdrop-blur-md border-b border-slate-100 flex items-center justify-between px-6 lg:px-10 sticky top-0 z-40">
        <button 
          @click="toggleMobileMenu"
          class="p-2 text-slate-500 hover:bg-slate-50 rounded-lg lg:hidden"
        >
          <Menu v-if="!isMobileMenuOpen" class="w-6 h-6" />
          <X v-else class="w-6 h-6" />
        </button>
        
        <div class="hidden lg:block">
          <h1 class="text-xl font-bold text-slate-900">{{ $route.meta.title || 'Onboardly' }}</h1>
        </div>

        <div class="flex items-center gap-4 relative">
          <!-- Profile Toggle Overlay -->
          <div v-if="isProfileOpen" @click="isProfileOpen = false" class="fixed inset-0 z-40"></div>
          
          <button 
            @click="toggleProfile"
            class="flex items-center gap-4 hover:bg-slate-50 p-2 rounded-xl transition-colors relative z-50"
          >
            <div class="text-right hidden sm:block">
              <p class="text-[15px] font-bold text-slate-900 leading-tight">Leonardo</p>
              <p class="text-xs font-medium text-slate-400">Analista de Implantação</p>
            </div>
            <div class="w-11 h-11 bg-brand-blue text-white rounded-full flex items-center justify-center font-bold text-sm border-2 border-white shadow-sm overflow-hidden">
              <img src="https://ui-avatars.com/api/?name=Leonardo+E&background=3B82F6&color=fff" alt="Avatar" class="w-full h-full object-cover" />
            </div>
          </button>

          <!-- Profile Dropdown -->
          <Transition
            enter-active-class="transition duration-200 ease-out"
            enter-from-class="transform scale-95 opacity-0 -translate-y-2"
            enter-to-class="transform scale-100 opacity-100 translate-y-0"
            leave-active-class="transition duration-150 ease-in"
            leave-from-class="transform scale-100 opacity-100 translate-y-0"
            leave-to-class="transform scale-95 opacity-0 -translate-y-2"
          >
            <div 
              v-if="isProfileOpen"
              class="absolute top-full right-0 mt-2 w-64 bg-white rounded-2xl shadow-xl border border-slate-100 py-3 z-50 overflow-hidden"
            >
              <div class="px-4 py-3 border-b border-slate-50 mb-2">
                <p class="text-sm font-bold text-slate-900">Leonardo E.</p>
                <p class="text-xs text-slate-500 truncate">leonardo@onboardly.com</p>
              </div>

              <div class="px-2">
                <a href="#" class="flex items-center gap-3 px-3 py-2 text-sm text-slate-700 hover:bg-slate-50 rounded-lg transition-colors">
                  <User class="w-4 h-4 text-slate-400" />
                  Meu Perfil
                </a>
                <a href="#" class="flex items-center gap-3 px-3 py-2 text-sm text-slate-700 hover:bg-slate-50 rounded-lg transition-colors">
                  <Shield class="w-4 h-4 text-slate-400" />
                  Segurança
                </a>
                <a href="#" @click.prevent="handleSettings" class="flex items-center gap-3 px-3 py-2 text-sm text-slate-700 hover:bg-slate-50 rounded-lg transition-colors">
                  <Settings class="w-4 h-4 text-slate-400" />
                  Configurações
                </a>
              </div>

              <div class="mt-2 pt-2 border-t border-slate-50 px-2">
                <button 
                  @click="handleLogout"
                  class="w-full flex items-center gap-3 px-3 py-2 text-sm text-rose-600 hover:bg-rose-50 rounded-lg transition-colors"
                >
                  <LogOut class="w-4 h-4" />
                  Sair da Conta
                </button>
              </div>
            </div>
          </Transition>
        </div>
      </header>

      <main class="p-6 lg:p-10 max-w-[1600px] mx-auto w-full">
        <router-view />
      </main>
    </div>
  </div>
</template>


<style scoped>
@reference "../../style.css";

.router-link-active {
  @apply bg-brand-blue/20 text-brand-blue shadow-[0_0_15px_rgba(59,130,246,0.3)];
}
</style>
