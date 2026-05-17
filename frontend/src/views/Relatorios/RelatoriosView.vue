<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { BarChart3, TrendingUp, Users, Clock, Award, Loader2 } from 'lucide-vue-next'
import { projectService } from '../../services/projectService'

interface Stats {
  total_projetos: number
  total_clientes: number
  reunioes_hoje: number
  por_status: Record<string, number>
  historico_mensal: Array<{ mes: string; total: number }>
  atividades_recentes: Array<{ tipo: string; descricao: string; status: string; data: string }>
}

const stats = ref<Stats | null>(null)
const isLoading = ref(true)

const fetchStats = async () => {
  try {
    const data = await projectService.getStats()
    stats.value = data
  } catch (error) {
    console.error('Error fetching stats:', error)
  } finally {
    isLoading.value = false
  }
}

onMounted(fetchStats)

const cards = computed(() => [
  { 
    name: 'Satisfação Média', 
    value: '4.9/5', 
    icon: Award, 
    color: 'text-amber-500', 
    bg: 'bg-amber-50',
    description: 'NPS'
  },
  { 
    name: 'Tempo Médio Go-Live', 
    value: '14 dias', 
    icon: Clock, 
    color: 'text-blue-500', 
    bg: 'bg-blue-50',
    description: 'Média global'
  },
  { 
    name: 'Projetos Ativos', 
    value: stats.value ? stats.value.total_projetos.toString() : '0', 
    icon: TrendingUp, 
    color: 'text-emerald-500', 
    bg: 'bg-emerald-50',
    description: 'Em andamento'
  },
  { 
    name: 'Total de Clientes', 
    value: stats.value ? stats.value.total_clientes.toString() : '0', 
    icon: Users, 
    color: 'text-purple-500', 
    bg: 'bg-purple-50',
    description: 'Base ativa'
  },
])

const maxMonthlyValue = computed(() => {
  if (!stats.value?.historico_mensal.length) return 1
  const max = Math.max(...stats.value.historico_mensal.map(m => m.total))
  return max === 0 ? 1 : max
})
</script>

<template>
  <div class="space-y-10">
    <div class="space-y-1">
      <h1 class="text-3xl font-bold tracking-tight text-[var(--color-text-primary)]">Relatórios e Performance</h1>
      <p class="text-[var(--color-text-secondary)] font-medium">Análise detalhada de entregas e eficiência do time.</p>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
      <div 
        v-for="card in cards" 
        :key="card.name"
        class="bg-[var(--color-surface)] p-6 rounded-[var(--radius-apple)] shadow-premium border border-[var(--color-border-soft)] relative overflow-hidden"
      >
        <div v-if="isLoading" class="absolute inset-0 bg-white/50 backdrop-blur-[1px] flex items-center justify-center z-10">
          <Loader2 class="w-5 h-5 text-[var(--color-primary)] animate-spin" />
        </div>
        
        <div class="flex items-center justify-between mb-4">
          <div :class="[card.bg, card.color, 'p-3 rounded-2xl']">
            <component :is="card.icon" class="w-6 h-6" />
          </div>
          <span class="text-[10px] font-black uppercase tracking-widest text-emerald-600 bg-emerald-50 px-2 py-0.5 rounded-full">Estável</span>
        </div>
        <p class="text-2xl font-black text-[var(--color-text-primary)]">{{ card.value }}</p>
        <p class="text-sm font-bold text-[var(--color-text-tertiary)] mt-1">{{ card.name }}</p>
      </div>
    </div>

    <!-- Charts Placeholder -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
      <!-- Volume Mensal -->
      <div class="bg-[var(--color-surface)] p-8 rounded-[var(--radius-apple)] shadow-premium border border-[var(--color-border-soft)] min-h-[400px] flex flex-col relative">
        <div class="flex items-center justify-between mb-8">
          <h3 class="text-lg font-bold text-[var(--color-text-primary)]">Volume de Projetos (Últimos 6 meses)</h3>
          <BarChart3 class="w-5 h-5 text-[var(--color-text-tertiary)]" />
        </div>

        <div v-if="isLoading" class="flex-1 flex items-center justify-center">
          <Loader2 class="w-8 h-8 text-[var(--color-primary)] animate-spin" />
        </div>
        
        <div v-else-if="stats?.historico_mensal.length" class="flex-1 flex items-end justify-between gap-4 px-2">
          <div v-for="item in stats.historico_mensal" :key="item.mes" class="flex-1 group relative">
            <div 
              class="w-full bg-[var(--color-primary)]/10 group-hover:bg-[var(--color-primary)] rounded-t-lg transition-all cursor-pointer relative"
              :style="{ height: (item.total / maxMonthlyValue * 100) + '%' }"
            >
              <div class="absolute -top-8 left-1/2 -translate-x-1/2 bg-zinc-900 text-white text-[10px] py-1 px-2 rounded opacity-0 group-hover:opacity-100 transition-opacity whitespace-nowrap">
                {{ item.total }} projetos
              </div>
            </div>
            <div class="absolute -bottom-6 left-1/2 -translate-x-1/2 text-[10px] font-bold text-[var(--color-text-tertiary)] uppercase tracking-tighter">
              {{ item.mes }}
            </div>
          </div>
        </div>
        <div v-else class="flex-1 flex items-center justify-center text-[var(--color-text-tertiary)] font-medium">
          Nenhum dado histórico disponível.
        </div>
      </div>

      <!-- Status das Implantações -->
      <div class="bg-[var(--color-surface)] p-8 rounded-[var(--radius-apple)] shadow-premium border border-[var(--color-border-soft)] min-h-[400px] flex flex-col">
        <div class="flex items-center justify-between mb-8">
          <h3 class="text-lg font-bold text-[var(--color-text-primary)]">Status das Implantações</h3>
          <div class="text-xs font-bold text-[var(--color-primary)]">Total: {{ stats?.total_projetos || 0 }}</div>
        </div>
        
        <div v-if="isLoading" class="flex-1 flex items-center justify-center">
          <Loader2 class="w-8 h-8 text-[var(--color-primary)] animate-spin" />
        </div>
        
        <div v-else-if="stats && Object.keys(stats.por_status).length" class="flex-1 space-y-6">
          <div v-for="(count, status) in stats.por_status" :key="status" class="space-y-2">
            <div class="flex justify-between text-sm font-bold">
              <span class="text-[var(--color-text-secondary)]">{{ status }}</span>
              <span class="text-[var(--color-text-primary)]">{{ count }}</span>
            </div>
            <div class="h-3 bg-zinc-100 rounded-full overflow-hidden">
              <div 
                class="h-full bg-[var(--color-primary)] rounded-full transition-all duration-1000"
                :style="{ width: (count / stats.total_projetos * 100) + '%' }"
              ></div>
            </div>
          </div>
        </div>
        <div v-else class="flex-1 flex items-center justify-center text-[var(--color-text-tertiary)] font-medium">
          Nenhum projeto registrado.
        </div>
      </div>
    </div>
  </div>
</template>
