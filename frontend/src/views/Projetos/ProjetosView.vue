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
  AlertCircle,
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
const selectedStatus = ref('All')

const showSlideOver = ref(false)
const selectedProject = ref<Projeto | undefined>(undefined)
const isSaving = ref(false)
const projectFormRef = ref<any>(null)

const statusConfig = {
// ... existing statusConfig ...
  'Backlog': { 
    label: 'Backlog', 
    color: 'text-brand-blue bg-brand-blue/10', 
    dot: 'bg-brand-blue',
    icon: Clock
  },
  'Em_Andamento': { 
    label: 'Em Andamento', 
    color: 'text-brand-amber bg-brand-amber/10', 
    dot: 'bg-brand-amber',
    icon: ArrowUpRight
  },
  'Concluido': { 
    label: 'Concluído', 
    color: 'text-brand-emerald bg-brand-emerald/10', 
    dot: 'bg-brand-emerald',
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
    
    // Map clients and analistas for easy lookup
    clients.value = clientsData.reduce((acc, client) => {
      acc[client.id] = client
      return acc
    }, {} as Record<string, Cliente>)
    
    analistas.value = analistasData.reduce((acc, analista) => {
      acc[analista.id] = analista
      return acc
    }, {} as Record<string, Analista>)

  } catch (error) {
    console.error('Erro ao buscar dados:', error)
    toastStore.error('Erro ao carregar dados.')
  } finally {
    loading.value = false
  }
}

const filteredProjects = computed(() => {
  return projects.value.filter(p => {
    const matchesStatus = selectedStatus.value === 'All' || p.status_projeto === selectedStatus.value
    
    const query = searchQuery.value.toLowerCase()
    const clientName = clients.value[p.cliente_id]?.nome.toLowerCase() || ''
    const analistaName = analistas.value[p.analista_id]?.nome.toLowerCase() || ''
    const matchesSearch = !searchQuery.value || 
                          clientName.includes(query) || 
                          analistaName.includes(query) ||
                          p.id.includes(query)

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
  <div class="space-y-8">
    <!-- Header Action Bar -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-4">
      <div>
        <h1 class="text-3xl font-bold text-slate-900 tracking-tight">Projetos</h1>
        <p class="text-slate-500 font-medium mt-1">Gerencie o pipeline de implantação dos seus clientes.</p>
      </div>

      <div class="flex items-center gap-3">
        <div class="relative group">
          <Search class="absolute left-3 top-1/2 -translate-y-1/2 w-4 h-4 text-slate-400 group-focus-within:text-brand-blue transition-colors" />
          <input 
            v-model="searchQuery"
            type="text" 
            placeholder="Buscar projeto..." 
            class="pl-10 pr-4 py-2.5 bg-white border border-slate-200 rounded-xl w-64 focus:outline-none focus:ring-2 focus:ring-brand-blue/20 focus:border-brand-blue transition-all"
          />
        </div>
        
        <div class="relative min-w-[160px]">
          <select 
            v-model="selectedStatus"
            class="w-full appearance-none bg-white border border-slate-200 rounded-xl px-4 py-2.5 pr-10 text-slate-600 font-medium focus:outline-none focus:ring-2 focus:ring-brand-blue/20 transition-all cursor-pointer"
          >
            <option value="All">Todos Status</option>
            <option value="Backlog">Backlog</option>
            <option value="Em_Andamento">Em Andamento</option>
            <option value="Concluido">Concluído</option>
          </select>
          <ChevronDown class="w-4 h-4 absolute right-3 top-1/2 -translate-y-1/2 text-slate-400 pointer-events-none" />
        </div>

        <button 
          @click="openNewProject"
          class="bg-brand-blue hover:bg-brand-blue/90 text-white px-5 py-2.5 rounded-xl font-bold flex items-center gap-2 transition-all active:scale-95 shadow-lg shadow-brand-blue/20"
        >
          <Plus class="w-5 h-5" />
          Novo Projeto
        </button>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="loading" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div v-for="i in 6" :key="i" class="h-64 bg-white rounded-[24px] border border-slate-100 flex items-center justify-center">
        <Loader2 class="w-8 h-8 text-slate-200 animate-spin" />
      </div>
    </div>

    <!-- Empty State -->
    <div v-else-if="filteredProjects.length === 0" class="bg-white rounded-[32px] p-20 border border-slate-100 flex flex-col items-center text-center shadow-soft">
      <div class="w-20 h-20 bg-slate-50 rounded-full flex items-center justify-center mb-6">
        <Package class="w-10 h-10 text-slate-300" />
      </div>
      <h3 class="text-xl font-bold text-slate-900">Nenhum projeto encontrado</h3>
      <p class="text-slate-500 max-w-sm mt-2">Não encontramos nenhum projeto com os critérios de busca ou ainda não há projetos cadastrados.</p>
      <button 
        @click="searchQuery = ''; selectedStatus = 'All'" 
        v-if="searchQuery || selectedStatus !== 'All'" 
        class="mt-6 text-brand-blue font-bold hover:underline"
      >
        Limpar filtros
      </button>
    </div>

    <!-- Projects Grid -->
    <div v-if="!loading && filteredProjects.length > 0" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <article 
        v-for="project in filteredProjects" 
        :key="project.id"
        @click="openEditProject(project)"
        class="bg-white rounded-[24px] p-6 border border-slate-50 shadow-soft hover:shadow-xl hover:-translate-y-1 transition-all duration-300 group cursor-pointer"
      >
        <div class="flex justify-between items-start mb-6">
          <div 
            :class="[
              'px-3 py-1 rounded-full text-[10px] font-bold uppercase tracking-wider flex items-center gap-1.5',
              statusConfig[project.status_projeto]?.color || 'bg-slate-100 text-slate-500'
            ]"
          >
            <div :class="['w-1.5 h-1.5 rounded-full', statusConfig[project.status_projeto]?.dot || 'bg-slate-400']"></div>
            {{ statusConfig[project.status_projeto]?.label || project.status_projeto }}
          </div>
          <button class="p-1 text-slate-300 hover:text-slate-600 transition-colors">
            <MoreVertical class="w-5 h-5" />
          </button>
        </div>

        <div class="space-y-4">
          <div>
            <div class="flex items-center gap-2 text-slate-400 mb-1">
              <Building2 class="w-3.5 h-3.5" />
              <span class="text-[11px] font-bold uppercase tracking-widest">Cliente</span>
            </div>
            <h3 class="text-lg font-bold text-slate-900 group-hover:text-brand-blue transition-colors">
              {{ clients[project.cliente_id]?.nome || 'Carregando...' }}
            </h3>
          </div>

          <div class="grid grid-cols-2 gap-4 pt-4 border-t border-slate-50">
            <div>
              <div class="flex items-center gap-2 text-slate-400 mb-1">
                <User class="w-3.5 h-3.5" />
                <span class="text-[11px] font-bold uppercase tracking-widest">Analista</span>
              </div>
              <p class="text-sm font-semibold text-slate-700 truncate">
                {{ analistas[project.analista_id]?.nome || 'Não atribuído' }}
              </p>
            </div>
            <div>
              <div class="flex items-center gap-2 text-slate-400 mb-1">
                <Calendar class="w-3.5 h-3.5" />
                <span class="text-[11px] font-bold uppercase tracking-widest">Início</span>
              </div>
              <p class="text-sm font-mono font-bold text-slate-700">
                {{ formatDate(project.data_contratacao) }}
              </p>
            </div>
          </div>
        </div>

        <div class="mt-6 flex items-center justify-between">
          <div class="flex -space-x-2">
            <div class="w-8 h-8 rounded-full bg-brand-blue text-white flex items-center justify-center text-[10px] font-bold border-2 border-white">
              {{ clients[project.cliente_id]?.nome?.substring(0,2).toUpperCase() || '?' }}
            </div>
            <div class="w-8 h-8 rounded-full bg-slate-100 text-slate-400 flex items-center justify-center text-[10px] font-bold border-2 border-white">
              {{ analistas[project.analista_id]?.nome?.substring(0,1).toUpperCase() || '?' }}
            </div>
          </div>
          
          <router-link 
            :to="`/projetos/${project.id}`"
            @click.stop
            class="flex items-center gap-1 text-xs font-bold text-brand-blue hover:gap-2 transition-all"
          >
            Detalhes
            <ArrowUpRight class="w-3.5 h-3.5" />
          </router-link>
        </div>
      </article>
    </div>

    <!-- SlideOver for Creating/Editing -->
    <SlideOver 
      :show="showSlideOver" 
      :title="selectedProject ? 'Editar Projeto' : 'Novo Projeto'"
      :description="selectedProject ? 'Atualize as informações do projeto selecionado.' : 'Preencha os dados para iniciar uma nova implantação.'"
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
          class="px-5 py-2.5 rounded-xl font-bold text-slate-500 hover:bg-slate-100 transition-all"
        >
          Cancelar
        </button>
        <button 
          @click="projectFormRef?.submit()"
          :disabled="isSaving"
          class="bg-brand-blue hover:bg-brand-blue/90 disabled:opacity-50 text-white px-8 py-2.5 rounded-xl font-bold flex items-center gap-2 transition-all active:scale-95 shadow-lg shadow-brand-blue/20"
        >
          <Loader2 v-if="isSaving" class="w-4 h-4 animate-spin" />
          <Save v-else class="w-4 h-4" />
          {{ selectedProject ? 'Salvar Alterações' : 'Criar Projeto' }}
        </button>
      </template>
    </SlideOver>
  </div>
</template>

<style scoped>
.shadow-soft {
  box-shadow: 0 4px 20px -2px rgba(0, 0, 0, 0.05), 0 2px 10px -2px rgba(0, 0, 0, 0.03);
}
</style>
