<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { Users, Briefcase, Video, ArrowUpRight, Clock, Loader2 } from 'lucide-vue-next'
import { projectService } from '../../services/projectService'
import { useAuthStore } from '../../stores/authStore'

const authStore = useAuthStore()
const loading = ref(true)
const dashboardStats = ref({
  total_projetos: 0,
  total_clientes: 0,
  reunioes_hoje: 0,
  por_status: {} as Record<string, number>,
  historico_mensal: [] as any[],
  atividades_recentes: [] as any[]
})

const stats = computed(() => [
  { 
    name: 'Clientes Totais', 
    value: dashboardStats.value.total_clientes.toString(), 
    icon: Users, 
    trend: '+12%',
    chartColor: '#0066FF'
  },
  { 
    name: 'Projetos Ativos', 
    value: (dashboardStats.value.por_status['Em_Andamento'] || 0).toString(), 
    icon: Briefcase, 
    trend: '+5%',
    chartColor: '#10B981'
  },
  { 
    name: 'Reuniões Hoje', 
    value: dashboardStats.value.reunioes_hoje.toString(), 
    icon: Video, 
    trend: 'Estável',
    chartColor: '#F59E0B'
  },
])

const fetchData = async () => {
  loading.value = true
  try {
    const data = await projectService.getStats()
    dashboardStats.value = data
  } catch (error) {
    console.error('Error fetching dashboard data:', error)
  } finally {
    loading.value = false
  }
}

const formatRelativeTime = (dateString: string) => {
  if (!dateString) return ''
  const date = new Date(dateString)
  const now = new Date()
  const diffInMinutes = Math.floor((now.getTime() - date.getTime()) / 60000)
  
  if (diffInMinutes < 1) return 'agora'
  if (diffInMinutes < 60) return `${diffInMinutes}m atrás`
  if (diffInMinutes < 1440) return `${Math.floor(diffInMinutes / 60)}h atrás`
  return date.toLocaleDateString('pt-BR')
}

const getDotColor = (tipo: string) => {
  switch (tipo) {
    case 'Cliente': return 'bg-blue-400'
    case 'Projeto': return 'bg-[var(--color-primary)]'
    case 'Reuniao': return 'bg-emerald-500'
    default: return 'bg-zinc-400'
  }
}

const getActivityTitle = (activity: any) => {
  if (activity.tipo === 'Cliente') return `Novo Cliente: ${activity.descricao}`
  return activity.descricao
}

onMounted(fetchData)
</script>

<template>
  <div class="space-y-10">
    <!-- Header Section -->
    <div class="flex flex-col gap-2">
      <h2 class="text-3xl font-bold tracking-tight text-[var(--color-text-primary)]">Resumo Geral</h2>
      <p class="text-[var(--color-text-secondary)] font-medium">Bem-vindo de volta, {{ authStore.user?.nome.split(' ')[0] }}. Veja o que mudou hoje.</p>
    </div>

    <!-- Stats Cards -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
      <div v-for="stat in stats" :key="stat.name" class="bg-[var(--color-surface)] p-8 rounded-[var(--radius-apple)] shadow-premium border border-[var(--color-border-soft)] group hover:shadow-premium-hover transition-all duration-300">
        <div class="flex justify-between items-start mb-4">
          <div class="w-12 h-12 bg-zinc-100 rounded-2xl flex items-center justify-center text-[var(--color-text-primary)] group-hover:scale-110 transition-transform">
            <component :is="stat.icon" class="w-6 h-6" />
          </div>
          <div class="flex items-center gap-1 text-[11px] font-bold text-emerald-600 bg-emerald-50 px-2.5 py-1 rounded-full">
            <ArrowUpRight class="w-3 h-3" />
            {{ stat.trend }}
          </div>
        </div>
        
        <div>
          <p class="text-[14px] font-semibold text-[var(--color-text-tertiary)]">{{ stat.name }}</p>
          <div v-if="loading" class="h-10 w-24 bg-zinc-100 animate-pulse rounded-lg mt-2"></div>
          <p v-else class="text-4xl font-bold text-[var(--color-text-primary)] mt-1 tracking-tight">{{ stat.value }}</p>
        </div>

        <!-- Mini Chart (Apple Style Line) -->
        <div class="h-12 w-full mt-6 opacity-30 group-hover:opacity-100 transition-opacity">
          <svg viewBox="0 0 400 100" class="w-full h-full">
            <path 
              d="M0,80 Q50,40 100,70 T200,30 T300,60 T400,20" 
              fill="none" 
              :stroke="stat.chartColor" 
              stroke-width="6" 
              stroke-linecap="round"
            />
          </svg>
        </div>
      </div>
    </div>

    <!-- Main Section -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <!-- Chart Column -->
      <div class="lg:col-span-2 bg-[var(--color-surface)] p-10 rounded-[var(--radius-apple)] shadow-premium border border-[var(--color-border-soft)]">
        <div class="flex justify-between items-center mb-12">
          <div>
            <h3 class="text-xl font-bold text-[var(--color-text-primary)] tracking-tight">Progresso de Implementação</h3>
            <p class="text-sm text-[var(--color-text-tertiary)] mt-1 font-medium">Comparativo mensal de projetos</p>
          </div>
          <select class="bg-zinc-100 border-none rounded-full px-4 py-2 text-xs font-bold text-[var(--color-text-secondary)] focus:ring-0">
            <option>Últimos 6 meses</option>
          </select>
        </div>

        <!-- Bar Chart (Real Data) -->
        <div class="h-64 flex items-end justify-between gap-4 px-2">
          <div v-if="loading" class="w-full flex items-center justify-center">
             <Loader2 class="w-8 h-8 text-[var(--color-primary)] animate-spin" />
          </div>
          <div v-else-if="dashboardStats.historico_mensal.length === 0" class="w-full text-center text-[var(--color-text-tertiary)] font-medium">
             Sem dados históricos para exibir.
          </div>
          <div v-else v-for="item in dashboardStats.historico_mensal" :key="item.mes" class="flex-1 flex flex-col items-center gap-4 group">
            <div class="w-full max-w-[40px] flex items-end gap-1 h-full">
              <div 
                class="flex-1 bg-zinc-100 rounded-t-lg transition-all duration-500 group-hover:bg-[var(--color-primary)]/20" 
                :style="{ height: '20%' }"
              ></div>
              <div 
                class="flex-1 bg-[var(--color-primary)] rounded-t-lg transition-all duration-500 group-hover:shadow-lg group-hover:shadow-blue-500/20" 
                :style="{ height: (Math.min(item.total * 20, 100)) + '%' }"
              ></div>
            </div>
            <span class="text-[11px] font-bold text-[var(--color-text-tertiary)] uppercase">{{ item.mes }}</span>
          </div>
        </div>
      </div>

      <!-- Recent Activities -->
      <div class="bg-[var(--color-surface)] p-10 rounded-[var(--radius-apple)] shadow-premium border border-[var(--color-border-soft)]">
        <div class="flex items-center justify-between mb-10">
          <h3 class="text-xl font-bold text-[var(--color-text-primary)] tracking-tight">Atividade Recente</h3>
          <Clock class="w-5 h-5 text-[var(--color-text-tertiary)]" />
        </div>
        
        <div class="space-y-10 relative">
          <!-- Timeline Vertical Line -->
          <div class="absolute left-1.5 top-2 bottom-2 w-px bg-zinc-100"></div>

          <div v-if="loading" class="flex justify-center">
             <Loader2 class="w-6 h-6 text-[var(--color-primary)] animate-spin" />
          </div>
          <div v-else-if="dashboardStats.atividades_recentes.length === 0" class="text-center text-[var(--color-text-tertiary)] text-sm font-medium py-10">
            Nenhuma atividade registrada.
          </div>
          <div v-else v-for="(activity, idx) in dashboardStats.atividades_recentes" :key="idx" class="flex gap-6 relative group">
            <div :class="[getDotColor(activity.tipo), 'w-3 h-3 rounded-full mt-1.5 flex-shrink-0 z-10 border-2 border-white ring-4 ring-zinc-50']"></div>
            
            <div class="space-y-1">
              <p class="text-[14px] font-bold text-[var(--color-text-primary)] leading-tight group-hover:text-[var(--color-primary)] transition-colors line-clamp-1">{{ getActivityTitle(activity) }}</p>
              <div class="flex items-center gap-3">
                <span class="text-[11px] font-bold text-[var(--color-text-tertiary)] uppercase tracking-wider">{{ activity.status }}</span>
                <span class="w-1 h-1 bg-zinc-300 rounded-full"></span>
                <span class="text-[12px] font-medium text-[var(--color-text-tertiary)]">{{ formatRelativeTime(activity.data) }}</span>
              </div>
            </div>
          </div>

          <button v-if="dashboardStats.atividades_recentes.length > 0" class="w-full py-3 mt-4 text-[13px] font-bold text-[var(--color-primary)] hover:bg-blue-50 rounded-xl transition-colors">
            Ver todas as atividades
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
