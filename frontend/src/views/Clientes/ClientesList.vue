<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useToastStore } from '../../stores/toastStore'
import { Plus, Search, ChevronDown, ChevronLeft, ChevronRight, ChevronsLeft, ChevronsRight, Loader2, Save } from 'lucide-vue-next'
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
const selectedAnalystId = ref('Todos os Analistas')
const selectedStatus = ref('Todos os Status')

const statuses = ['Todos os Status', 'Go-Live', 'Em Andamento', 'Backlog']

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
  return projects.value.find(p => p.cliente_id === clientId)
}

const getAnalystName = (analystId: string) => {
  return analysts.value.find(a => a.id === analystId)?.nome || 'Não atribuído'
}

const mapStatus = (status: string | undefined) => {
  if (!status) return 'Sem Projeto'
  switch (status) {
    case 'Concluido': return 'Go-Live'
    case 'Em_Andamento': return 'Em Andamento'
    case 'Backlog': return 'Backlog'
    default: return status
  }
}

const getStatusClass = (status: string) => {
  switch (status) {
    case 'Go-Live': return 'bg-emerald-50 text-emerald-600 border-emerald-100'
    case 'Em Andamento': return 'bg-blue-50 text-blue-600 border-blue-100'
    case 'Backlog': return 'bg-zinc-100 text-zinc-500 border-zinc-200'
    default: return 'bg-zinc-50 text-zinc-500'
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
      progress: project?.status_projeto === 'Concluido' ? 100 : project?.status_projeto === 'Em_Andamento' ? 45 : 0,
      logoInitials: client.nome.substring(0, 1).toUpperCase(),
    }
  }).filter(item => {
    const matchesSearch = item.nome.toLowerCase().includes(searchQuery.value.toLowerCase()) || 
                          item.cnpj.includes(searchQuery.value)
    
    const matchesAnalyst = selectedAnalystId.value === 'Todos os Analistas' || 
                           (item.project && item.project.analista_id === selectedAnalystId.value)
    
    const matchesStatus = selectedStatus.value === 'Todos os Status' || 
                          item.displayStatus === selectedStatus.value

    return matchesSearch && matchesAnalyst && matchesStatus
  })
})

onMounted(fetchData)
</script>

<template>
  <div class="space-y-10">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl font-bold tracking-tight text-[var(--color-text-primary)]">Clientes e Projetos</h1>
        <p class="text-[var(--color-text-secondary)] font-medium">Gerencie sua carteira de clientes e acompanhe cada estágio.</p>
      </div>
      <button 
        @click="openNewClient"
        class="bg-[var(--color-primary)] text-white px-6 py-3 rounded-full font-bold flex items-center gap-2 hover:bg-[var(--color-primary-hover)] transition-all shadow-lg shadow-blue-500/20 active:scale-95 text-sm"
      >
        <Plus class="w-5 h-5" />
        Novo Cliente
      </button>
    </div>

    <!-- Filters & Search Bar -->
    <div class="bg-[var(--color-surface)] p-2 rounded-2xl shadow-premium border border-[var(--color-border-soft)] flex flex-col md:flex-row items-center gap-2">
      <div class="relative flex-1 w-full md:w-auto">
        <Search class="w-4 h-4 absolute left-4 top-1/2 -translate-y-1/2 text-[var(--color-text-tertiary)]" />
        <input 
          v-model="searchQuery"
          type="text" 
          placeholder="Buscar por nome ou CNPJ..." 
          class="w-full bg-transparent border-none focus:ring-0 pl-11 pr-4 py-3 text-sm text-[var(--color-text-primary)] placeholder:text-[var(--color-text-tertiary)]"
        />
      </div>

      <div class="h-8 w-px bg-[var(--color-border-divider)] hidden md:block"></div>

      <div class="flex items-center gap-2 w-full md:w-auto px-2">
        <select 
          v-model="selectedAnalystId"
          class="bg-transparent border-none focus:ring-0 text-sm font-semibold text-[var(--color-text-secondary)] cursor-pointer py-3"
        >
          <option>Todos os Analistas</option>
          <option v-for="a in analysts" :key="a.id" :value="a.id">{{ a.nome }}</option>
        </select>

        <div class="h-6 w-px bg-[var(--color-border-divider)]"></div>

        <select 
          v-model="selectedStatus"
          class="bg-transparent border-none focus:ring-0 text-sm font-semibold text-[var(--color-text-secondary)] cursor-pointer py-3"
        >
          <option v-for="s in statuses" :key="s">{{ s }}</option>
        </select>
      </div>
    </div>

    <!-- Table Container -->
    <div class="bg-[var(--color-surface)] rounded-[var(--radius-apple)] shadow-premium border border-[var(--color-border-soft)] overflow-hidden">
      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="border-b border-[var(--color-border-soft)]">
              <th class="px-8 py-5 text-[12px] font-bold text-[var(--color-text-tertiary)] uppercase tracking-wider">Cliente</th>
              <th class="px-8 py-5 text-[12px] font-bold text-[var(--color-text-tertiary)] uppercase tracking-wider">CNPJ</th>
              <th class="px-8 py-5 text-[12px] font-bold text-[var(--color-text-tertiary)] uppercase tracking-wider">Analista</th>
              <th class="px-8 py-5 text-[12px] font-bold text-[var(--color-text-tertiary)] uppercase tracking-wider text-center">Status</th>
              <th class="px-8 py-5 text-[12px] font-bold text-[var(--color-text-tertiary)] uppercase tracking-wider">Progresso</th>
            </tr>
          </thead>
          <tbody v-if="isLoading">
            <tr>
              <td colspan="5" class="px-8 py-20 text-center">
                <div class="flex flex-col items-center gap-3">
                  <Loader2 class="w-10 h-10 text-[var(--color-primary)] animate-spin" />
                  <span class="text-[var(--color-text-tertiary)] font-medium">Carregando carteira...</span>
                </div>
              </td>
            </tr>
          </tbody>
          <tbody v-else-if="filteredClients.length === 0">
            <tr>
              <td colspan="5" class="px-8 py-20 text-center">
                <span class="text-[var(--color-text-tertiary)] font-medium">Nenhum cliente encontrado com os filtros atuais.</span>
              </td>
            </tr>
          </tbody>
          <tbody v-else class="divide-y divide-[var(--color-border-soft)]">
            <tr 
              v-for="client in filteredClients" 
              :key="client.id" 
              @click="openEditClient(client)"
              class="hover:bg-zinc-50 transition-colors cursor-pointer group"
            >
              <td class="px-8 py-6">
                <div class="flex items-center gap-4">
                  <div class="w-10 h-10 rounded-xl bg-[var(--color-primary)]/5 flex items-center justify-center text-[var(--color-primary)] font-bold text-sm">
                    {{ client.logoInitials }}
                  </div>
                  <span class="font-bold text-[var(--color-text-primary)] text-[15px]">{{ client.nome }}</span>
                </div>
              </td>
              <td class="px-8 py-6 text-sm font-medium text-[var(--color-text-secondary)] font-mono">{{ client.cnpj }}</td>
              <td class="px-8 py-6">
                <div class="flex items-center gap-3">
                  <div class="w-7 h-7 rounded-full bg-zinc-200 overflow-hidden">
                    <img :src="`https://ui-avatars.com/api/?name=${client.analystName}&background=E5E5EA&color=1C1C1E`" alt="Analyst" />
                  </div>
                  <span class="font-semibold text-[var(--color-text-primary)] text-sm">{{ client.analystName }}</span>
                </div>
              </td>
              <td class="px-8 py-6 text-center">
                <span 
                  :class="[
                    getStatusClass(client.displayStatus),
                    'px-3 py-1 rounded-full text-[12px] font-bold border'
                  ]"
                >
                  {{ client.displayStatus }}
                </span>
              </td>
              <td class="px-8 py-6">
                <div class="flex items-center gap-4 min-w-[180px]">
                  <div class="flex-1 h-1.5 bg-zinc-100 rounded-full overflow-hidden">
                    <div 
                      class="h-full bg-[var(--color-primary)] rounded-full transition-all duration-1000" 
                      :style="{ width: client.progress + '%' }"
                    ></div>
                  </div>
                  <span class="text-[13px] font-bold text-[var(--color-text-primary)]">{{ client.progress }}%</span>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="px-8 py-6 border-t border-[var(--color-border-soft)] flex items-center justify-between">
        <span class="text-sm text-[var(--color-text-tertiary)] font-medium">Exibindo {{ filteredClients.length }} clientes</span>
        <div class="flex items-center gap-2">
          <button class="p-2 text-[var(--color-text-tertiary)] hover:bg-zinc-100 rounded-full transition-colors"><ChevronLeft class="w-5 h-5" /></button>
          <div class="flex items-center gap-1">
            <button class="w-8 h-8 flex items-center justify-center rounded-lg bg-[var(--color-primary)] text-white text-sm font-bold shadow-sm">1</button>
            <button class="w-8 h-8 flex items-center justify-center rounded-lg hover:bg-zinc-100 text-[var(--color-text-secondary)] text-sm font-bold">2</button>
          </div>
          <button class="p-2 text-[var(--color-text-tertiary)] hover:bg-zinc-100 rounded-full transition-colors"><ChevronRight class="w-5 h-5" /></button>
        </div>
      </div>
    </div>

    <!-- SlideOver for Creating/Editing -->
    <SlideOver 
      :show="showSlideOver" 
      :title="selectedClient ? 'Editar Cliente' : 'Novo Cliente'"
      :description="selectedClient ? 'Atualize as informações cadastrais do cliente.' : 'Preencha os dados abaixo para cadastrar um novo cliente.'"
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
          class="px-6 py-3 rounded-full font-bold text-[var(--color-text-secondary)] hover:bg-zinc-100 transition-all text-sm"
        >
          Cancelar
        </button>
        <button 
          @click="clienteFormRef?.submit()"
          :disabled="isSaving"
          class="bg-[var(--color-primary)] hover:bg-[var(--color-primary-hover)] disabled:opacity-50 text-white px-8 py-3 rounded-full font-bold flex items-center gap-2 transition-all active:scale-95 shadow-lg shadow-blue-500/20 text-sm"
        >
          <Loader2 v-if="isSaving" class="w-4 h-4 animate-spin" />
          <Save v-else class="w-4 h-4" />
          {{ selectedClient ? 'Salvar Alterações' : 'Cadastrar Cliente' }}
        </button>
      </template>
    </SlideOver>
  </div>
</template>
