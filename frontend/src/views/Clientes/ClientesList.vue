<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useToastStore } from '../../stores/toastStore'
import { Plus, Search, Calendar as CalendarIcon, ChevronDown, ChevronLeft, ChevronRight, ChevronsLeft, ChevronsRight, Loader2, Save } from 'lucide-vue-next'
import { clientService } from '../../services/clientService'
import { projectService } from '../../services/projectService'
import { analistaService } from '../../services/analistaService'
import type { Cliente, Projeto, Analista } from '../../types'

import SlideOver from '../../components/ui/SlideOver.vue'
import ClienteForm from '../../components/forms/ClienteForm.vue'

const toastStore = useToastStore()
const clients = ref<Cliente[]>([])
const projects = ref<Projeto[]>([])
const analysts = ref<Analista[]>([])
const isLoading = ref(true)

const showSlideOver = ref(false)
const selectedClient = ref<Cliente | undefined>(undefined)
const isSaving = ref(false)
const clienteFormRef = ref<any>(null)

const searchQuery = ref('')
// ... rest of state ...
const selectedAnalystId = ref('All Analysts')
const selectedStatus = ref('All Status')

const statuses = ['All Status', 'Go-Live', 'In Progress', 'Backlog']

const fetchData = async () => {
  isLoading.value = true
  try {
    const [clientsData, projectsData, analystsData] = await Promise.all([
      clientService.getAll(),
      projectService.getAll(),
      analistaService.getAll()
    ])
    clients.value = clientsData
    projects.value = projectsData
    analysts.value = analystsData
  } catch (error) {
    console.error('Error fetching data:', error)
    toastStore.error('Erro ao carregar dados.')
  } finally {
    isLoading.value = false
  }
}

const openNewClient = () => {
  selectedClient.value = undefined
  showSlideOver.value = true
}

const openEditClient = (client: Cliente) => {
  selectedClient.value = client
  showSlideOver.value = true
}

const handleSave = async (formData: any) => {
  isSaving.value = true
  try {
    if (selectedClient.value) {
      await clientService.update(selectedClient.value.id, formData)
      toastStore.success('Cliente atualizado com sucesso!')
    } else {
      await clientService.create(formData)
      toastStore.success('Cliente cadastrado com sucesso!')
    }
    await fetchData()
    showSlideOver.value = false
  } catch (error) {
    console.error('Error saving client:', error)
    toastStore.error('Erro ao salvar cliente.')
  } finally {
    isSaving.value = false
  }
}

const getClientProject = (clientId: string) => {
// ... existing getClientProject ...
  return projects.value.find(p => p.cliente_id === clientId)
}

const getAnalystName = (analystId: string) => {
  return analysts.value.find(a => a.id === analystId)?.nome || 'Not assigned'
}

const mapStatus = (status: string | undefined) => {
  if (!status) return 'No Project'
  switch (status) {
    case 'Concluido': return 'Go-Live'
    case 'Em_Andamento': return 'In Progress'
    case 'Backlog': return 'Backlog'
    default: return status
  }
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'Go-Live': return 'bg-emerald-50 text-emerald-600 border-emerald-100'
    case 'In Progress': return 'bg-blue-50 text-blue-600 border-blue-100'
    case 'Backlog': return 'bg-slate-100 text-slate-500 border-slate-200'
    default: return 'bg-slate-50 text-slate-500'
  }
}

const filteredClients = computed(() => {
  return clients.value.map(client => {
    const project = getClientProject(client.id)
    return {
      ...client,
      project,
      analystName: project ? getAnalystName(project.analista_id) : 'N/A',
      displayStatus: mapStatus(project?.status_projeto),
      // Mock progress since it's not in the model yet
      progress: project?.status_projeto === 'Concluido' ? 100 : project?.status_projeto === 'Em_Andamento' ? 45 : 0,
      logoInitials: client.nome.substring(0, 1).toUpperCase(),
      logoColor: 'bg-brand-blue' // Default color
    }
  }).filter(item => {
    const matchesSearch = item.nome.toLowerCase().includes(searchQuery.value.toLowerCase()) || 
                          item.cnpj.includes(searchQuery.value)
    
    const matchesAnalyst = selectedAnalystId.value === 'All Analysts' || 
                           (item.project && item.project.analista_id === selectedAnalystId.value)
    
    const matchesStatus = selectedStatus.value === 'All Status' || 
                          item.displayStatus === selectedStatus.value

    return matchesSearch && matchesAnalyst && matchesStatus
  })
})

onMounted(fetchData)
</script>

<template>
  <div class="space-y-8">
    <div class="flex justify-between items-center">
      <h1 class="text-[32px] font-bold text-slate-900">Clients & Projects</h1>
      <button 
        @click="openNewClient"
        class="bg-brand-blue text-white px-6 py-3 rounded-xl font-bold flex items-center gap-2 hover:bg-blue-600 transition-all shadow-lg shadow-blue-200 active:scale-95"
      >
        <Plus class="w-5 h-5" />
        New Client
      </button>
    </div>

    <!-- Filters Bar -->
    <div class="bg-white p-6 rounded-[24px] shadow-soft border border-slate-50 flex flex-wrap gap-6 items-end">
      <div class="space-y-2 flex-1 min-w-[200px]">
        <label class="text-sm font-bold text-slate-900 ml-1">Analyst</label>
        <div class="relative group">
          <select 
            v-model="selectedAnalystId"
            class="w-full appearance-none bg-slate-50 border border-slate-100 rounded-xl px-4 py-3 pr-10 text-slate-600 font-medium focus:outline-none focus:ring-2 focus:ring-brand-blue/20 focus:bg-white transition-all cursor-pointer"
          >
            <option value="All Analysts">All Analysts</option>
            <option v-for="a in analysts" :key="a.id" :value="a.id">{{ a.nome }}</option>
          </select>
          <ChevronDown class="w-4 h-4 absolute right-4 top-1/2 -translate-y-1/2 text-slate-400 pointer-events-none group-hover:text-slate-600 transition-colors" />
        </div>
      </div>

      <div class="space-y-2 flex-1 min-w-[200px]">
        <label class="text-sm font-bold text-slate-900 ml-1">Status</label>
        <div class="relative group">
          <select 
            v-model="selectedStatus"
            class="w-full appearance-none bg-slate-50 border border-slate-100 rounded-xl px-4 py-3 pr-10 text-slate-600 font-medium focus:outline-none focus:ring-2 focus:ring-brand-blue/20 focus:bg-white transition-all cursor-pointer"
          >
            <option v-for="s in statuses" :key="s">{{ s }}</option>
          </select>
          <ChevronDown class="w-4 h-4 absolute right-4 top-1/2 -translate-y-1/2 text-slate-400 pointer-events-none group-hover:text-slate-600 transition-colors" />
        </div>
      </div>

      <div class="space-y-2 flex-1 min-w-[240px]">
        <label class="text-sm font-bold text-slate-900 ml-1">Date Range</label>
        <div class="relative">
          <CalendarIcon class="w-5 h-5 absolute left-4 top-1/2 -translate-y-1/2 text-slate-400 pointer-events-none" />
          <input 
            type="text" 
            placeholder="Date range Date" 
            class="w-full bg-slate-50 border border-slate-100 rounded-xl pl-12 pr-4 py-3 text-slate-600 font-medium focus:outline-none focus:ring-2 focus:ring-brand-blue/20 focus:bg-white transition-all"
          />
        </div>
      </div>

      <div class="relative flex-[1.5] min-w-[300px]">
        <Search class="w-5 h-5 absolute left-4 top-1/2 -translate-y-1/2 text-slate-400 pointer-events-none" />
        <input 
          v-model="searchQuery"
          type="text" 
          placeholder="Search clients or CNPJ..." 
          class="w-full bg-white border border-slate-200 rounded-xl pl-12 pr-4 py-3 text-slate-600 font-medium shadow-sm focus:outline-none focus:ring-2 focus:ring-brand-blue/20 transition-all placeholder:text-slate-300"
        />
      </div>
    </div>

    <!-- Table Container -->
    <div class="bg-white rounded-[24px] shadow-soft border border-slate-50 overflow-hidden">
      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="bg-slate-50/50">
              <th class="px-8 py-5 text-sm font-bold text-slate-500">Client Name</th>
              <th class="px-8 py-5 text-sm font-bold text-slate-500">CNPJ</th>
              <th class="px-8 py-5 text-sm font-bold text-slate-500 text-center">Assigned Analyst</th>
              <th class="px-8 py-5 text-sm font-bold text-slate-500 text-center">Project Status</th>
              <th class="px-8 py-5 text-sm font-bold text-slate-500">Implementation Progress</th>
            </tr>
          </thead>
          <tbody v-if="isLoading" class="divide-y divide-slate-100">
            <tr>
              <td colspan="5" class="px-8 py-10 text-center">
                <div class="flex flex-col items-center gap-2">
                  <Loader2 class="w-8 h-8 text-brand-blue animate-spin" />
                  <span class="text-slate-500 font-medium">Loading clients...</span>
                </div>
              </td>
            </tr>
          </tbody>
          <tbody v-else-if="filteredClients.length === 0" class="divide-y divide-slate-100">
            <tr>
              <td colspan="5" class="px-8 py-10 text-center">
                <span class="text-slate-500 font-medium">No clients found.</span>
              </td>
            </tr>
          </tbody>
          <tbody v-else class="divide-y divide-slate-100">
            <tr 
              v-for="client in filteredClients" 
              :key="client.id" 
              @click="openEditClient(client)"
              class="hover:bg-slate-50/80 transition-colors cursor-pointer group"
            >
              <td class="px-8 py-5">
                <div class="flex items-center gap-4">
                  <div :class="[client.logoColor, 'w-10 h-10 rounded-lg flex items-center justify-center text-white font-bold text-sm shadow-sm']">
                    {{ client.logoInitials }}
                  </div>
                  <span class="font-bold text-slate-900">{{ client.nome }}</span>
                </div>
              </td>
              <td class="px-8 py-5 font-medium text-slate-500">{{ client.cnpj }}</td>
              <td class="px-8 py-5">
                <div class="flex items-center justify-center gap-3">
                  <div class="w-8 h-8 rounded-full bg-slate-200 overflow-hidden border border-white shadow-sm">
                    <img :src="`https://ui-avatars.com/api/?name=${client.analystName}&background=random`" alt="Analyst" />
                  </div>
                  <span class="font-bold text-slate-900 text-sm">{{ client.analystName }}</span>
                </div>
              </td>
              <td class="px-8 py-5 text-center">
                <span 
                  :class="[
                    getStatusClass(client.displayStatus),
                    'px-4 py-1.5 rounded-full text-[13px] font-bold border'
                  ]"
                >
                  {{ client.displayStatus }}
                </span>
              </td>
              <td class="px-8 py-5">
                <div class="flex items-center gap-4 min-w-[200px]">
                  <div class="flex-1 h-2 bg-slate-100 rounded-full overflow-hidden">
                    <div 
                      class="h-full bg-brand-blue rounded-full transition-all duration-1000" 
                      :style="{ width: client.progress + '%' }"
                    ></div>
                  </div>
                  <span class="text-sm font-bold text-slate-900">{{ client.progress }}%</span>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="p-6 border-t border-slate-50 flex items-center justify-center gap-4">
        <button class="p-2 text-slate-300 hover:text-slate-600 transition-colors"><ChevronsLeft class="w-5 h-5" /></button>
        <button class="p-2 text-slate-300 hover:text-slate-600 transition-colors"><ChevronLeft class="w-5 h-5" /></button>
        <div class="w-8 h-8 bg-blue-50 text-brand-blue flex items-center justify-center rounded-lg font-bold border border-blue-100">1</div>
        <button class="p-2 text-slate-400 hover:text-slate-600 transition-colors"><ChevronRight class="w-5 h-5" /></button>
        <button class="p-2 text-slate-400 hover:text-slate-600 transition-colors"><ChevronsRight class="w-5 h-5" /></button>
      </div>
    </div>

    <!-- SlideOver for Creating/Editing -->
    <SlideOver 
      :show="showSlideOver" 
      :title="selectedClient ? 'Editar Cliente' : 'Novo Cliente'"
      :description="selectedClient ? 'Atualize as informações cadastrais do cliente.' : 'Cadastre um novo cliente no sistema.'"
      @close="showSlideOver = false"
    >
      <ClienteForm 
        ref="clienteFormRef"
        :initial-data="selectedClient"
        @save="handleSave"
      />
      
      <template #footer>
        <button 
          @click="showSlideOver = false"
          class="px-5 py-2.5 rounded-xl font-bold text-slate-500 hover:bg-slate-100 transition-all"
        >
          Cancelar
        </button>
        <button 
          @click="clienteFormRef?.submit()"
          :disabled="isSaving"
          class="bg-brand-blue hover:bg-brand-blue/90 disabled:opacity-50 text-white px-8 py-2.5 rounded-xl font-bold flex items-center gap-2 transition-all active:scale-95 shadow-lg shadow-brand-blue/20"
        >
          <Loader2 v-if="isSaving" class="w-4 h-4 animate-spin" />
          <Save v-else class="w-4 h-4" />
          {{ selectedClient ? 'Salvar Alterações' : 'Cadastrar Cliente' }}
        </button>
      </template>
    </SlideOver>
  </div>
</template>
