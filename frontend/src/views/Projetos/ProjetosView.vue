<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useToastStore } from '../../stores/toastStore'
import { 
  Package, 
  Search, 
  Plus, 
  Calendar, 
  User, 
  Building2,
  MoreVertical,
  ArrowUpRight,
  Clock,
  CheckCircle2,
  Loader2,
  ChevronDown,
  Save
} from 'lucide-vue-next'
import { projectService } from '../../services/projectService'
import { clientService } from '../../services/clientService'
import { analistaService } from '../../services/analistaService'
import type { Projeto, Cliente, Analista } from '../../types'

import SlideOver from '../../components/ui/SlideOver.vue'
import ProjetoForm from '../../components/forms/ProjetoForm.vue'

const toastStore = useToastStore()
const projects = ref<Projeto[]>([])
const clients = ref<Record<string, Cliente>>({})
const analistas = ref<Record<string, Analista>>({})
const loading = ref(true)
const searchQuery = ref('')
const selectedStatus = ref('Todos Status')

const showSlideOver = ref(false)
const selectedProject = ref<Projeto | undefined>(undefined)
const isSaving = ref(false)
const projectFormRef = ref<any>(null)

const statusConfig = {
  'Backlog': { 
    label: 'Backlog', 
    color: 'text-zinc-500 bg-zinc-100', 
    dot: 'bg-zinc-400',
    icon: Clock
  },
  'Em_Andamento': { 
    label: 'Em Andamento', 
    color: 'text-blue-600 bg-blue-50', 
    dot: 'bg-blue-500',
    icon: ArrowUpRight
  },
  'Concluido': { 
    label: 'Concluído', 
    color: 'text-emerald-600 bg-emerald-50', 
    dot: 'bg-emerald-500',
    icon: CheckCircle2
  }
}

const fetchData = async () => {
  loading.value = true
  try {
    const [projectsData, clientsData, analistasData] = await Promise.all([
      projectService.getAll(),
      clientService.getAll(),
      analistaService.getAll()
    ])
    
    projects.value = projectsData
    clients.value = clientsData.reduce((acc, client) => { acc[client.id] = client; return acc }, {} as Record<string, Cliente>)
    analistas.value = analistasData.reduce((acc, analista) => { acc[analista.id] = analista; return acc }, {} as Record<string, Analista>)
  } catch (error) {
    console.error('Erro ao buscar dados:', error)
    toastStore.error('Erro ao carregar dados.')
  } finally {
    loading.value = false
  }
}

const filteredProjects = computed(() => {
  return projects.value.filter(p => {
    const matchesStatus = selectedStatus.value === 'Todos Status' || p.status_projeto === selectedStatus.value
    const query = searchQuery.value.toLowerCase()
    const clientName = clients.value[p.cliente_id]?.nome.toLowerCase() || ''
    const analistaName = analistas.value[p.analista_id]?.nome.toLowerCase() || ''
    const matchesSearch = !searchQuery.value || clientName.includes(query) || analistaName.includes(query)
    return matchesStatus && matchesSearch
  })
})

const formatDate = (dateString: string) => {
  return new Date(dateString).toLocaleDateString('pt-BR', {
    day: '2-digit',
    month: 'short',
    year: 'numeric'
  })
}

onMounted(fetchData)

const openNewProject = () => {
  selectedProject.value = undefined
  showSlideOver.value = true
}

const openEditProject = (project: Projeto) => {
  selectedProject.value = project
  showSlideOver.value = true
}

const handleSave = async (formData: any) => {
  isSaving.value = true
  try {
    if (selectedProject.value) {
      await projectService.update(selectedProject.value.id, formData)
      toastStore.success('Projeto atualizado com sucesso!')
    } else {
      await projectService.create(formData)
      toastStore.success('Projeto criado com sucesso!')
    }
    await fetchData()
    showSlideOver.value = false
  } catch (error) {
    console.error('Erro ao salvar projeto:', error)
    toastStore.error('Erro ao salvar projeto.')
  } finally {
    isSaving.value = false
  }
}
</script>

<template>
  <div class="space-y-10">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl font-bold tracking-tight text-[var(--color-text-primary)]">Projetos</h1>
        <p class="text-[var(--color-text-secondary)] font-medium">Acompanhe e gerencie cada etapa da jornada de implementação.</p>
      </div>
      <button 
        @click="openNewProject"
        class="bg-[var(--color-primary)] text-white px-6 py-3 rounded-full font-bold flex items-center gap-2 hover:bg-[var(--color-primary-hover)] transition-all shadow-lg shadow-blue-500/20 active:scale-95 text-sm"
      >
        <Plus class="w-5 h-5" />
        Novo Projeto
      </button>
    </div>

    <!-- Toolbar -->
    <div class="bg-[var(--color-surface)] p-2 rounded-2xl shadow-premium border border-[var(--color-border-soft)] flex flex-col md:flex-row items-center gap-2">
      <div class="relative flex-1 w-full md:w-auto">
        <Search class="w-4 h-4 absolute left-4 top-1/2 -translate-y-1/2 text-[var(--color-text-tertiary)]" />
        <input 
          v-model="searchQuery"
          type="text" 
          placeholder="Buscar projeto por cliente ou analista..." 
          class="w-full bg-transparent border-none focus:ring-0 pl-11 pr-4 py-3 text-sm text-[var(--color-text-primary)] placeholder:text-[var(--color-text-tertiary)]"
        />
      </div>

      <div class="h-8 w-px bg-[var(--color-border-divider)] hidden md:block"></div>

      <div class="flex items-center gap-2 w-full md:w-auto px-2">
        <div class="flex items-center gap-2 text-sm font-semibold text-[var(--color-text-secondary)] py-3 px-2">
          <span>Status:</span>
          <select 
            v-model="selectedStatus"
            class="bg-transparent border-none focus:ring-0 text-sm font-bold text-[var(--color-primary)] cursor-pointer p-0"
          >
            <option>Todos Status</option>
            <option value="Backlog">Backlog</option>
            <option value="Em_Andamento">Em Andamento</option>
            <option value="Concluido">Concluído</option>
          </select>
        </div>
      </div>
    </div>

    <!-- Grid Container -->
    <div v-if="loading" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
      <div v-for="i in 6" :key="i" class="h-64 bg-zinc-50 rounded-[var(--radius-apple)] animate-pulse border border-zinc-100"></div>
    </div>

    <div v-else-if="filteredProjects.length === 0" class="bg-[var(--color-surface)] rounded-[var(--radius-apple)] p-20 border border-[var(--color-border-soft)] shadow-premium flex flex-col items-center text-center">
      <div class="w-20 h-20 bg-zinc-50 rounded-full flex items-center justify-center mb-6">
        <Package class="w-10 h-10 text-zinc-300" />
      </div>
      <h3 class="text-xl font-bold text-[var(--color-text-primary)]">Nenhum projeto encontrado</h3>
      <p class="text-[var(--color-text-tertiary)] max-w-sm mt-2 font-medium">Não há projetos que correspondam aos seus filtros ou busca.</p>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
      <article 
        v-for="project in filteredProjects" 
        :key="project.id"
        @click="openEditProject(project)"
        class="bg-[var(--color-surface)] rounded-[var(--radius-apple)] p-8 border border-[var(--color-border-soft)] shadow-premium hover:shadow-premium-hover hover:-translate-y-1.5 transition-all duration-300 group cursor-pointer flex flex-col"
      >
        <div class="flex justify-between items-start mb-8">
          <div 
            :class="[
              'px-3 py-1.5 rounded-full text-[11px] font-bold uppercase tracking-wider flex items-center gap-2 border border-transparent',
              statusConfig[project.status_projeto]?.color || 'bg-zinc-100 text-zinc-500'
            ]"
          >
            <div :class="['w-1.5 h-1.5 rounded-full', statusConfig[project.status_projeto]?.dot || 'bg-zinc-400']"></div>
            {{ statusConfig[project.status_projeto]?.label || project.status_projeto }}
          </div>
          <button class="p-1.5 text-zinc-300 hover:text-zinc-600 transition-colors">
            <MoreVertical class="w-5 h-5" />
          </button>
        </div>

        <div class="space-y-6 flex-1">
          <div class="space-y-1.5">
            <div class="flex items-center gap-2 text-[var(--color-text-tertiary)]">
              <Building2 class="w-4 h-4" />
              <span class="text-[11px] font-bold uppercase tracking-widest">Cliente</span>
            </div>
            <h3 class="text-xl font-bold text-[var(--color-text-primary)] group-hover:text-[var(--color-primary)] transition-colors tracking-tight">
              {{ clients[project.cliente_id]?.nome || 'Carregando...' }}
            </h3>
          </div>

          <div class="grid grid-cols-2 gap-6 pt-6 border-t border-[var(--color-border-soft)]">
            <div class="space-y-1.5">
              <div class="flex items-center gap-2 text-[var(--color-text-tertiary)]">
                <User class="w-4 h-4" />
                <span class="text-[11px] font-bold uppercase tracking-widest">Analista</span>
              </div>
              <p class="text-sm font-bold text-[var(--color-text-primary)] truncate">
                {{ analistas[project.analista_id]?.nome || 'Não atribuído' }}
              </p>
            </div>
            <div class="space-y-1.5">
              <div class="flex items-center gap-2 text-[var(--color-text-tertiary)]">
                <Calendar class="w-4 h-4" />
                <span class="text-[11px] font-bold uppercase tracking-widest">Início</span>
              </div>
              <p class="text-sm font-bold text-[var(--color-text-primary)] font-mono">
                {{ formatDate(project.data_contratacao) }}
              </p>
            </div>
          </div>
        </div>

        <div class="mt-8 pt-6 border-t border-[var(--color-border-soft)] flex items-center justify-between">
          <div class="flex -space-x-2">
            <div class="w-8 h-8 rounded-lg bg-zinc-100 flex items-center justify-center text-[10px] font-bold text-[var(--color-text-primary)] border-2 border-white">
              {{ clients[project.cliente_id]?.nome?.substring(0,1).toUpperCase() || '?' }}
            </div>
            <div class="w-8 h-8 rounded-lg bg-[var(--color-primary)] text-white flex items-center justify-center text-[10px] font-bold border-2 border-white">
              {{ analistas[project.analista_id]?.nome?.substring(0,1).toUpperCase() || '?' }}
            </div>
          </div>
          
          <button class="flex items-center gap-1.5 text-[13px] font-bold text-[var(--color-primary)] group-hover:translate-x-1 transition-transform">
            Ver detalhes
            <ArrowUpRight class="w-4 h-4" />
          </button>
        </div>
      </article>
    </div>

    <!-- SlideOver -->
    <SlideOver 
      :show="showSlideOver" 
      :title="selectedProject ? 'Editar Projeto' : 'Novo Projeto'"
      :description="selectedProject ? 'Atualize o cronograma e atribuição do projeto.' : 'Configure os parâmetros iniciais da nova implantação.'"
      @close="showSlideOver = false"
    >
      <ProjetoForm 
        ref="projectFormRef"
        :initial-data="selectedProject"
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
          @click="projectFormRef?.submit()"
          :disabled="isSaving"
          class="bg-[var(--color-primary)] hover:bg-[var(--color-primary-hover)] disabled:opacity-50 text-white px-8 py-3 rounded-full font-bold flex items-center gap-2 transition-all active:scale-95 shadow-lg shadow-blue-500/20 text-sm"
        >
          <Loader2 v-if="isSaving" class="w-4 h-4 animate-spin" />
          <Save v-else class="w-4 h-4" />
          {{ selectedProject ? 'Salvar Alterações' : 'Criar Projeto' }}
        </button>
      </template>
    </SlideOver>
  </div>
</template>
