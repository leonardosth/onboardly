<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useToastStore } from '../../stores/toastStore'
import { useAuthStore } from '../../stores/authStore'
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
  Search,
  Bell
} from 'lucide-vue-next'

const router = useRouter()
const toastStore = useToastStore()
const authStore = useAuthStore()

const isProfileOpen = ref(false)
const isMobileMenuOpen = ref(false)

const user = computed(() => authStore.user)
const userName = computed(() => user.value?.nome || 'Usuário')
const userEmail = computed(() => user.value?.email || '')
const userCargo = computed(() => user.value?.cargo || 'Analista')
const userAvatarUrl = computed(() => `https://ui-avatars.com/api/?name=${encodeURIComponent(userName.value)}&background=0066FF&color=fff`)

const menuItems = computed(() => {
  const baseItems = [
    { name: 'Dashboard', icon: LayoutGrid, path: '/' },
    { name: 'Clientes', icon: UserCircle, path: '/clientes' },
    { name: 'Projetos', icon: Package, path: '/projetos' },
    { name: 'Reuniões', icon: Calendar, path: '/reunioes' },
  ]

  if (authStore.user?.cargo === 'Admin') {
    baseItems.push({ name: 'Analistas', icon: User, path: '/analistas' })
  }

  baseItems.push({ name: 'Relatórios', icon: BarChart3, path: '/relatorios' })
  
  return baseItems
})

const toggleProfile = () => {
  isProfileOpen.value = !isProfileOpen.value
}

const toggleMobileMenu = () => {
  isMobileMenuOpen.value = !isMobileMenuOpen.value
}

const handleLogout = () => {
  authStore.logout()
  toastStore.info('Sessão encerrada.')
  router.push('/login')
}

const handleSettings = () => {
  toastStore.info('Configurações em desenvolvimento...')
}
</script>

<template>
  <div class="h-screen bg-[var(--color-background-app)] flex overflow-hidden">
    <!-- Mobile Menu Overlay -->
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
        class="fixed inset-0 bg-zinc-900/20 backdrop-blur-sm z-40 lg:hidden"
      ></div>
    </Transition>

    <!-- Navigation Sidebar -->
    <aside 
      class="fixed inset-y-0 left-0 z-50 w-72 transform transition-transform duration-300 ease-apple lg:translate-x-0 lg:static lg:inset-0 bg-[var(--color-surface)] border-r border-[var(--color-border-soft)]"
      :class="[isMobileMenuOpen ? 'translate-x-0 shadow-premium' : '-translate-x-full']"
    >
      <div class="h-full flex flex-col p-6">
        <!-- Logo Area -->
        <div class="flex items-center gap-3 px-2 mb-10 cursor-default select-none">
          <div class="w-9 h-9 bg-[var(--color-primary)] rounded-xl flex items-center justify-center shadow-lg shadow-blue-500/20">
            <div class="w-1.5 h-1.5 bg-white rounded-full"></div>
          </div>
          <span class="text-xl font-bold tracking-tight text-[var(--color-text-primary)]">Onboardly</span>
        </div>

        <!-- Navigation Links -->
        <nav class="flex-1 space-y-1 overflow-y-auto pr-2 custom-scrollbar">
          <router-link 
            v-for="item in menuItems" 
            :key="item.path"
            :to="item.path"
            @click="isMobileMenuOpen = false"
            class="flex items-center gap-3 px-4 py-3 rounded-[var(--radius-apple-sm)] transition-all duration-200 group relative font-medium"
            :class="[$route.path === item.path ? 'bg-[var(--color-primary)]/5 text-[var(--color-primary)]' : 'text-[var(--color-text-secondary)] hover:bg-zinc-50 hover:text-[var(--color-text-primary)]']"
          >
            <component :is="item.icon" class="w-5 h-5" />
            <span>{{ item.name }}</span>
            <div v-if="$route.path === item.path" class="absolute left-0 w-1 h-5 bg-[var(--color-primary)] rounded-full -translate-x-0.5"></div>
          </router-link>
        </nav>

        <!-- Sidebar Bottom Section -->
        <div class="mt-auto space-y-1 pt-4 border-t border-[var(--color-border-soft)]">
          <button 
            @click="handleSettings"
            class="w-full flex items-center gap-3 px-4 py-3 text-[var(--color-text-secondary)] hover:bg-zinc-50 hover:text-[var(--color-text-primary)] rounded-[var(--radius-apple-sm)] transition-all font-medium"
          >
            <Settings class="w-5 h-5" />
            <span>Ajustes</span>
          </button>
          <button 
            @click="handleLogout"
            class="w-full flex items-center gap-3 px-4 py-3 text-rose-500 hover:bg-rose-50 rounded-[var(--radius-apple-sm)] transition-all font-medium"
          >
            <LogOut class="w-5 h-5" />
            <span>Sair</span>
          </button>
        </div>
      </div>
    </aside>

    <!-- Main View Area -->
    <div class="flex-1 flex flex-col min-w-0">
      <!-- Top Navigation Bar -->
      <header class="h-18 glass sticky top-0 z-40 px-6 lg:px-10 flex items-center justify-between border-b border-[var(--color-border-soft)]">
        <div class="flex items-center gap-4">
          <button 
            @click="toggleMobileMenu"
            class="p-2 text-[var(--color-text-secondary)] hover:bg-zinc-100 rounded-lg lg:hidden"
          >
            <Menu v-if="!isMobileMenuOpen" class="w-6 h-6" />
            <X v-else class="w-6 h-6" />
          </button>
          
          <div class="hidden sm:flex items-center gap-2 bg-zinc-100/50 px-4 py-2 rounded-full border border-zinc-200/50 w-64 lg:w-96">
            <Search class="w-4 h-4 text-[var(--color-text-tertiary)]" />
            <input type="text" placeholder="Buscar..." class="bg-transparent border-none focus:ring-0 text-sm w-full placeholder:text-[var(--color-text-tertiary)] text-[var(--color-text-primary)]" />
          </div>
        </div>

        <div class="flex items-center gap-2 lg:gap-6">
          <button class="p-2 text-[var(--color-text-secondary)] hover:bg-zinc-100 rounded-full transition-colors relative">
            <Bell class="w-5 h-5" />
            <span class="absolute top-2 right-2.5 w-2 h-2 bg-rose-500 border-2 border-white rounded-full"></span>
          </button>

          <div class="h-8 w-px bg-[var(--color-border-divider)] hidden lg:block"></div>

          <div class="relative">
            <button 
              @click="toggleProfile"
              class="flex items-center gap-3 p-1 rounded-full hover:bg-zinc-100 transition-colors"
            >
              <div class="w-9 h-9 rounded-full overflow-hidden border border-[var(--color-border-soft)] shadow-sm">
                <img :src="userAvatarUrl" :alt="userName" class="w-full h-full object-cover" />
              </div>
              <div class="hidden lg:block text-left mr-1">
                <p class="text-sm font-semibold text-[var(--color-text-primary)] leading-tight">{{ userName }}</p>
                <p class="text-[11px] font-medium text-[var(--color-text-tertiary)]">{{ userCargo }}</p>
              </div>
            </button>

            <!-- Profile Dropdown -->
            <Transition
              enter-active-class="transition duration-200 ease-apple"
              enter-from-class="transform scale-95 opacity-0 translate-y-2"
              enter-to-class="transform scale-100 opacity-100 translate-y-0"
              leave-active-class="transition duration-150 ease-apple"
              leave-from-class="transform scale-100 opacity-100 translate-y-0"
              leave-to-class="transform scale-95 opacity-0 translate-y-2"
            >
              <div 
                v-if="isProfileOpen"
                class="absolute top-full right-0 mt-3 w-64 bg-[var(--color-surface)] rounded-[var(--radius-apple)] shadow-premium border border-[var(--color-border-soft)] py-2 z-50 overflow-hidden"
              >
                <div class="px-5 py-4 border-b border-[var(--color-border-soft)] mb-1">
                  <p class="text-sm font-bold text-[var(--color-text-primary)]">{{ userName }}</p>
                  <p class="text-xs text-[var(--color-text-tertiary)] truncate">{{ userEmail }}</p>
                </div>
                <div class="p-2 space-y-0.5">
                  <button class="w-full flex items-center gap-3 px-3 py-2 text-sm text-[var(--color-text-secondary)] hover:bg-zinc-50 rounded-xl transition-colors">
                    <User class="w-4 h-4" />
                    Perfil
                  </button>
                  <button class="w-full flex items-center gap-3 px-3 py-2 text-sm text-[var(--color-text-secondary)] hover:bg-zinc-50 rounded-xl transition-colors">
                    <Shield class="w-4 h-4" />
                    Segurança
                  </button>
                </div>
                <div class="mt-1 pt-1 border-t border-[var(--color-border-soft)] p-2">
                  <button @click="handleLogout" class="w-full flex items-center gap-3 px-3 py-2 text-sm text-rose-600 hover:bg-rose-50 rounded-xl transition-colors">
                    <LogOut class="w-4 h-4" />
                    Sair da Conta
                  </button>
                </div>
              </div>
            </Transition>
          </div>
        </div>
      </header>

      <!-- Content Area -->
      <main class="flex-1 overflow-y-auto p-6 lg:p-12">
        <div class="max-w-7xl mx-auto space-y-8">
          <router-view v-slot="{ Component }">
            <Transition name="fade" mode="out-in">
              <component :is="Component" />
            </Transition>
          </router-view>
        </div>
      </main>
    </div>
  </div>
</template>

<style>
.ease-apple {
  transition-timing-function: cubic-bezier(0.25, 0.1, 0.25, 1);
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s cubic-bezier(0.25, 0.1, 0.25, 1), transform 0.2s cubic-bezier(0.25, 0.1, 0.25, 1);
}

.fade-enter-from {
  opacity: 0;
  transform: translateY(4px);
}

.fade-leave-to {
  opacity: 0;
  transform: translateY(-4px);
}
</style>
