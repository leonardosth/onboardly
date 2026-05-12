<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { Users, Briefcase, Video, Loader2 } from 'lucide-vue-next'
import { projectService } from '../../services/projectService'
import { clientService } from '../../services/clientService'
import { reuniaoService } from '../../services/reuniaoService'

const loading = ref(true)
const dashboardStats = ref({
  total_projetos: 0,
  por_status: {} as Record<string, number>
})
const totalClients = ref(0)
const meetingsToday = ref(0)

const stats = computed(() => [
  { 
    name: 'Total Clients', 
    value: totalClients.value.toString(), 
    icon: Users, 
    color: 'text-brand-blue', 
    bg: 'bg-blue-50',
    chartColor: '#3B82F6'
  },
  { 
    name: 'Projetos Ativos', 
    value: (dashboardStats.value.por_status['Em_Andamento'] || 0).toString(), 
    icon: Briefcase, 
    color: 'text-brand-emerald', 
    bg: 'bg-emerald-50',
    chartColor: '#10B981'
  },
  { 
    name: 'Reuniões Hoje', 
    value: meetingsToday.value.toString(), 
    icon: Video, 
    color: 'text-brand-amber', 
    bg: 'bg-amber-50',
    chartColor: '#F59E0B'
  },
])

const fetchData = async () => {
  loading.value = true
  try {
    const [statsData, clientsData, meetingsData] = await Promise.all([
      projectService.getStats(),
      clientService.getAll(),
      reuniaoService.getAll()
    ])
    
    dashboardStats.value = statsData
    totalClients.value = clientsData.length
    
    // Filter meetings for today
    const today = new Date().toISOString().split('T')[0]
    meetingsToday.value = meetingsData.filter(m => m.data_agendada.startsWith(today)).length

  } catch (error) {
    console.error('Error fetching dashboard data:', error)
  } finally {
    loading.value = false
  }
}

const activities = [
  { id: 1, title: 'Projeto Beta - Status Atualizado', status: 'Backlog', time: '1h atrás', color: 'bg-brand-blue' },
  { id: 2, title: 'Reunião de Kick-off - Concluída', status: 'Completed', time: '2h atrás', color: 'bg-brand-emerald' },
  { id: 3, title: 'Novo Cliente - Cadastrado', status: 'Active', time: '4h atrás', color: 'bg-brand-cyan' },
  { id: 4, title: 'Tarefa de Análise - Concluída', status: 'Completed', time: 'ontem', color: 'bg-brand-emerald' },
]

const months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun']
const chartData = computed(() => {
  // Mocking historical data for the chart as we don't have it in the backend yet
  return [
    { backlog: 12, completed: 8 },
    { backlog: 15, completed: 10 },
    { backlog: 18, completed: 11 },
    { backlog: 19, completed: 14 },
    { backlog: 20, completed: 16 },
    { backlog: dashboardStats.value.por_status['Backlog'] || 0, completed: dashboardStats.value.por_status['Concluido'] || 0 },
  ]
})

onMounted(fetchData)
</script>

<template>
  <div class="space-y-8">
    <!-- Stats Cards -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
      <div v-for="stat in stats" :key="stat.name" class="bg-white p-6 rounded-[24px] shadow-soft border border-slate-50 flex flex-col gap-4 relative overflow-hidden group">
        <div class="flex items-center gap-4">
          <div :class="[stat.bg, stat.color, 'w-12 h-12 rounded-2xl flex items-center justify-center transition-transform group-hover:scale-110']">
            <component :is="stat.icon" class="w-6 h-6" />
          </div>
          <div>
            <p class="text-[15px] font-medium text-slate-500">{{ stat.name }}</p>
            <div v-if="loading" class="h-8 w-16 bg-slate-100 animate-pulse rounded mt-1"></div>
            <p v-else class="text-3xl font-bold text-slate-900">{{ stat.value }}</p>
          </div>
        </div>
        
        <!-- Sparkline SVG -->
        <div class="h-16 w-full mt-2">
          <svg viewBox="0 0 400 100" class="w-full h-full preserve-3d">
            <defs>
              <linearGradient :id="'grad-' + stat.name" x1="0%" y1="0%" x2="0%" y2="100%">
                <stop offset="0%" :style="{ stopColor: stat.chartColor, stopOpacity: 0.2 }" />
                <stop offset="100%" :style="{ stopColor: stat.chartColor, stopOpacity: 0 }" />
              </linearGradient>
            </defs>
            <path 
              d="M0,80 Q50,40 100,70 T200,30 T300,60 T400,20 L400,100 L0,100 Z" 
              :fill="'url(#grad-' + stat.name + ')'" 
            />
            <path 
              d="M0,80 Q50,40 100,70 T200,30 T300,60 T400,20" 
              fill="none" 
              :stroke="stat.chartColor" 
              stroke-width="4" 
              stroke-linecap="round"
            />
          </svg>
        </div>
      </div>
    </div>

    <!-- Main Section -->
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <!-- Chart Column -->
      <div class="lg:col-span-2 bg-white p-8 rounded-[24px] shadow-soft border border-slate-50">
        <div class="flex justify-between items-center mb-10">
          <h2 class="text-xl font-bold text-slate-900">Implementation Progress</h2>
          <div class="flex items-center gap-6">
            <div class="flex items-center gap-2">
              <div class="w-8 h-3 bg-brand-blue rounded-full"></div>
              <span class="text-sm font-medium text-slate-500">Backlog</span>
            </div>
            <div class="flex items-center gap-2">
              <div class="w-8 h-3 bg-brand-cyan rounded-full"></div>
              <span class="text-sm font-medium text-slate-500">Completed Projects</span>
            </div>
          </div>
        </div>

        <!-- Custom CSS Bar Chart -->
        <div class="relative h-[300px] flex items-end justify-between px-4 pb-8 border-b border-slate-100">
          <!-- Grid Lines -->
          <div class="absolute inset-0 flex flex-col justify-between pointer-events-none pb-8">
            <div v-for="i in 5" :key="i" class="w-full border-t border-slate-100 flex justify-between items-center">
              <span class="text-[10px] text-slate-300 -mt-2">{{ (5 - i) * 5 + 5 }}</span>
            </div>
          </div>

          <!-- Bars -->
          <div v-for="(data, idx) in chartData" :key="idx" class="flex items-end gap-1.5 z-10 w-16 group">
            <div 
              class="w-full bg-brand-blue rounded-t-lg transition-all duration-500 hover:brightness-110 relative" 
              :style="{ height: Math.max(data.backlog * 4, 10) + 'px' }"
            >
              <div class="absolute -top-8 left-1/2 -translate-x-1/2 bg-slate-900 text-white text-[10px] px-2 py-1 rounded opacity-0 group-hover:opacity-100 transition-opacity">
                {{ data.backlog }}
              </div>
            </div>
            <div 
              class="w-full bg-brand-cyan rounded-t-lg transition-all duration-500 hover:brightness-110 relative" 
              :style="{ height: Math.max(data.completed * 4, 10) + 'px' }"
            >
               <div class="absolute -top-8 left-1/2 -translate-x-1/2 bg-slate-900 text-white text-[10px] px-2 py-1 rounded opacity-0 group-hover:opacity-100 transition-opacity">
                {{ data.completed }}
              </div>
            </div>
          </div>
        </div>
        
        <!-- X Axis Labels -->
        <div class="flex justify-between px-4 mt-4">
          <span v-for="month in months" :key="month" class="text-xs font-semibold text-slate-400 w-16 text-center">{{ month }}</span>
        </div>
      </div>

      <!-- Recent Activities -->
      <div class="bg-white p-8 rounded-[24px] shadow-soft border border-slate-50">
        <h2 class="text-xl font-bold text-slate-900 mb-8">Recent Activities</h2>
        
        <div class="relative space-y-10 pl-8">
          <!-- Timeline Line -->
          <div class="absolute left-[11px] top-2 bottom-2 w-0.5 bg-slate-100"></div>

          <div v-for="activity in activities" :key="activity.id" class="relative group">
            <!-- Dot -->
            <div :class="[activity.color, 'absolute -left-[30px] top-1.5 w-4 h-4 rounded-full border-4 border-white ring-2 ring-slate-100 z-10 group-hover:scale-125 transition-transform']"></div>
            
            <div class="space-y-1">
              <p class="text-[15px] font-bold text-slate-900 leading-snug">{{ activity.title }}</p>
              <div class="flex items-center gap-2">
                <span 
                  :class="[
                    'text-[10px] font-bold uppercase tracking-wider px-2 py-0.5 rounded-full text-white',
                    activity.color
                  ]"
                >
                  {{ activity.status }}
                </span>
                <span class="text-xs font-medium text-slate-400">{{ activity.time }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
