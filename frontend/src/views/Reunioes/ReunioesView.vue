<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useToastStore } from '../../stores/toastStore'
import { 
  Bold, Italic, Underline, Strikethrough, Code, Link as LinkIcon, 
  List, ListOrdered, Quote, Image as ImageIcon, Table as TableIcon,
  CheckCircle2, Clock, XCircle, Loader2, Plus, Save, CalendarDays, Package, User
} from 'lucide-vue-next'
import { reuniaoService } from '../../services/reuniaoService'
import { projectService } from '../../services/projectService'
import { clientService } from '../../services/clientService'
import { analistaService } from '../../services/analistaService'
import type { Reuniao, Projeto, Cliente, Analista } from '../../types'

import SlideOver from '../../components/ui/SlideOver.vue'
import ReuniaoForm from '../../components/forms/ReuniaoForm.vue'

const toastStore = useToastStore()
const meetings = ref<Reuniao[]>([])
const projects = ref<Record<string, Projeto>>({})
const clients = ref<Record<string, Cliente>>({})
const analistas = ref<Record<string, Analista>>({})
const isLoading = ref(true)
const selectedMeetingId = ref<string | null>(null)

const showSlideOver = ref(false)
const editingMeeting = ref<Reuniao | undefined>(undefined)
const isSaving = ref(false)
const reuniaoFormRef = ref<any>(null)

const selectedMeeting = computed(() => {
  return meetings.value.find(m => m.id === selectedMeetingId.value) || null
})

const fetchData = async () => {
  isLoading.value = true
  try {
    const [meetingsData, projectsData, clientsData, analistasData] = await Promise.all([
      reuniaoService.getAll(),
      projectService.getAll(),
      clientService.getAll(),
      analistaService.getAll()
    ])
    
    meetings.value = meetingsData
    projects.value = projectsData.reduce((acc, p) => { acc[p.id] = p; return acc }, {} as Record<string, Projeto>)
    clients.value = clientsData.reduce((acc, c) => { acc[c.id] = c; return acc }, {} as Record<string, Cliente>)
    analistas.value = analistasData.reduce((acc, a) => { acc[a.id] = a; return acc }, {} as Record<string, Analista>)

    if (meetingsData.length > 0 && !selectedMeetingId.value) {
      selectedMeetingId.value = meetingsData[0].id
    }
  } catch (error) {
    console.error('Error fetching meetings data:', error)
    toastStore.error('Erro ao carregar dados.')
  } finally {
    isLoading.value = false
  }
}

const openNewMeeting = () => {
  editingMeeting.value = undefined
  showSlideOver.value = true
}

const openEditMeeting = (meeting: Reuniao) => {
  editingMeeting.value = meeting
  showSlideOver.value = true
}

const handleSave = async (formData: any) => {
  isSaving.value = true
  try {
    if (editingMeeting.value) {
      await reuniaoService.update(editingMeeting.value.id, formData)
      toastStore.success('Reunião atualizada com sucesso!')
    } else {
      await reuniaoService.create(formData)
      toastStore.success('Reunião agendada com sucesso!')
    }
    await fetchData()
    showSlideOver.value = false
  } catch (error) {
    console.error('Error saving meeting:', error)
    toastStore.error('Erro ao salvar reunião.')
  } finally {
    isSaving.value = false
  }
}

const getClientNameByProjectId = (projectId: string) => {
  const project = projects.value[projectId]
  if (!project) return 'Projeto não encontrado'
  return clients.value[project.cliente_id]?.nome || 'Cliente não encontrado'
}

const formatTime = (dateString: string) => {
  return new Date(dateString).toLocaleTimeString('pt-BR', { hour: '2-digit', minute: '2-digit' })
}

const getStatusLabel = (status: string) => {
  switch (status) {
    case 'Agendada': return 'Agendada'
    case 'Realizada': return 'Realizada'
    case 'No_Show': return 'No Show'
    default: return status
  }
}

onMounted(fetchData)
</script>

<template>
  <div class="space-y-10">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl font-bold tracking-tight text-[var(--color-text-primary)]">Reuniões Técnicas</h1>
        <p class="text-[var(--color-text-secondary)] font-medium">Acompanhe seus agendamentos e registre notas de implementação.</p>
      </div>
      <button 
        @click="openNewMeeting"
        class="bg-[var(--color-primary)] text-white px-6 py-3 rounded-full font-bold flex items-center gap-2 hover:bg-[var(--color-primary-hover)] transition-all shadow-lg shadow-blue-500/20 active:scale-95 text-sm"
      >
        <Plus class="w-5 h-5" />
        Agendar Reunião
      </button>
    </div>

    <div v-if="isLoading" class="flex justify-center items-center h-96">
      <div class="flex flex-col items-center gap-4">
        <Loader2 class="w-12 h-12 text-[var(--color-primary)] animate-spin" />
        <p class="text-[var(--color-text-tertiary)] font-medium text-sm">Carregando cronograma...</p>
      </div>
    </div>

    <div v-else-if="meetings.length === 0" class="bg-[var(--color-surface)] p-20 rounded-[var(--radius-apple)] border border-[var(--color-border-soft)] shadow-premium text-center">
      <div class="w-20 h-20 bg-zinc-50 rounded-full flex items-center justify-center mb-6 mx-auto">
        <CalendarDays class="w-10 h-10 text-zinc-300" />
      </div>
      <h3 class="text-xl font-bold text-[var(--color-text-primary)]">Sem reuniões</h3>
      <p class="text-[var(--color-text-tertiary)] font-medium mt-2">Agende uma reunião para começar a documentar.</p>
    </div>

    <div v-else class="grid grid-cols-1 lg:grid-cols-5 gap-10 items-start">
      <!-- Meetings List -->
      <div class="lg:col-span-2 space-y-6">
        <div 
          v-for="meeting in meetings" 
          :key="meeting.id"
          @click="selectedMeetingId = meeting.id"
          :class="[
            selectedMeetingId === meeting.id ? 'ring-2 ring-[var(--color-primary)] shadow-lg' : 'hover:bg-zinc-50 border-[var(--color-border-soft)]',
            'bg-[var(--color-surface)] p-6 rounded-2xl border transition-all cursor-pointer group flex gap-5'
          ]"
        >
          <div class="flex flex-col items-center justify-center text-center w-14">
            <span class="text-[12px] font-bold text-[var(--color-text-tertiary)] uppercase tracking-wider">{{ new Date(meeting.data_agendada).toLocaleDateString('pt-BR', { weekday: 'short' }) }}</span>
            <span class="text-2xl font-black text-[var(--color-text-primary)] leading-tight">{{ new Date(meeting.data_agendada).getDate() }}</span>
            <span class="text-[14px] font-bold text-[var(--color-primary)]">{{ formatTime(meeting.data_agendada) }}</span>
          </div>

          <div class="flex-1 space-y-3">
            <div>
              <p class="text-[15px] font-bold text-[var(--color-text-primary)] leading-tight">{{ getClientNameByProjectId(meeting.projeto_id) }}</p>
              <p class="text-[12px] font-bold text-[var(--color-text-tertiary)] mt-1 uppercase tracking-widest">Reunião de Acompanhamento</p>
            </div>

            <div class="flex items-center justify-between">
              <span 
                :class="[
                  meeting.status === 'Realizada' ? 'bg-emerald-50 text-emerald-600 border-emerald-100' : meeting.status === 'No_Show' ? 'bg-rose-50 text-rose-600 border-rose-100' : 'bg-blue-50 text-blue-600 border-blue-100',
                  'flex items-center gap-1.5 px-3 py-1 rounded-full text-[11px] font-bold border'
                ]"
              >
                {{ getStatusLabel(meeting.status) }}
              </span>
              <button 
                @click.stop="openEditMeeting(meeting)"
                class="text-[11px] font-bold text-[var(--color-text-tertiary)] hover:text-[var(--color-primary)] opacity-0 group-hover:opacity-100 transition-opacity"
              >
                EDITAR
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Meeting Details -->
      <div v-if="selectedMeeting" class="lg:col-span-3 bg-[var(--color-surface)] rounded-[var(--radius-apple)] shadow-premium border border-[var(--color-border-soft)] sticky top-28 overflow-hidden">
        <div class="p-10 border-b border-[var(--color-border-soft)]">
          <div class="flex justify-between items-start mb-8">
            <div class="space-y-1">
              <h2 class="text-2xl font-bold text-[var(--color-text-primary)] tracking-tight">{{ getClientNameByProjectId(selectedMeeting.projeto_id) }}</h2>
              <p class="text-[var(--color-text-tertiary)] font-bold text-[12px] uppercase tracking-widest">Detalhes do Encontro</p>
            </div>
            <div :class="[selectedMeeting.status === 'Realizada' ? 'bg-emerald-50 text-emerald-600' : 'bg-blue-50 text-blue-600', 'px-4 py-2 rounded-xl text-xs font-bold border border-transparent']">
              {{ getStatusLabel(selectedMeeting.status) }}
            </div>
          </div>
          
          <div class="grid grid-cols-2 gap-8">
            <div class="space-y-1.5">
              <div class="flex items-center gap-2 text-[var(--color-text-tertiary)]">
                <CalendarDays class="w-4 h-4" />
                <span class="text-[11px] font-bold uppercase tracking-widest">Data e Hora</span>
              </div>
              <p class="text-sm font-bold text-[var(--color-text-primary)]">{{ new Date(selectedMeeting.data_agendada).toLocaleString('pt-BR', { dateStyle: 'long', timeStyle: 'short' }) }}</p>
            </div>
            <div class="space-y-1.5">
              <div class="flex items-center gap-2 text-[var(--color-text-tertiary)]">
                <Package class="w-4 h-4" />
                <span class="text-[11px] font-bold uppercase tracking-widest">Projeto Relacionado</span>
              </div>
              <p class="text-sm font-bold text-[var(--color-text-primary)] font-mono">ID: {{ selectedMeeting.projeto_id.substring(0,8) }}...</p>
            </div>
            <div class="space-y-1.5">
              <div class="flex items-center gap-2 text-[var(--color-text-tertiary)]">
                <User class="w-4 h-4" />
                <span class="text-[11px] font-bold uppercase tracking-widest">Analista Responsável</span>
              </div>
              <p class="text-sm font-bold text-[var(--color-text-primary)]">
                {{ analistas[selectedMeeting.analista_id]?.nome || 'Não atribuído' }}
              </p>
            </div>
          </div>
        </div>

        <div class="p-10 space-y-6">
          <div class="flex items-center justify-between">
            <h3 class="text-lg font-bold text-[var(--color-text-primary)]">Notas Técnicas</h3>
            <div class="flex gap-2">
              <button class="p-2 hover:bg-zinc-100 rounded-lg text-[var(--color-text-tertiary)]"><Bold class="w-4 h-4" /></button>
              <button class="p-2 hover:bg-zinc-100 rounded-lg text-[var(--color-text-tertiary)]"><List class="w-4 h-4" /></button>
              <button class="p-2 hover:bg-zinc-100 rounded-lg text-[var(--color-text-tertiary)]"><LinkIcon class="w-4 h-4" /></button>
            </div>
          </div>
          
          <div class="relative">
            <textarea 
              v-model="selectedMeeting.observacoes"
              class="w-full bg-zinc-50 border border-[var(--color-border-soft)] rounded-2xl p-6 min-h-[300px] focus:outline-none focus:bg-white focus:ring-4 focus:ring-blue-500/5 focus:border-[var(--color-primary)] text-[var(--color-text-primary)] leading-relaxed font-medium transition-all placeholder:text-[var(--color-text-tertiary)]"
              placeholder="Descreva os pontos discutidos, decisões tomadas e próximos passos..."
            ></textarea>
            <div class="absolute bottom-4 right-4 flex items-center gap-2 text-[10px] font-bold text-[var(--color-text-tertiary)]">
              <span>SALVAMENTO AUTOMÁTICO ATIVO</span>
              <div class="w-1.5 h-1.5 bg-emerald-500 rounded-full"></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- SlideOver -->
    <SlideOver 
      :show="showSlideOver" 
      :title="editingMeeting ? 'Editar Reunião' : 'Nova Reunião'"
      :description="editingMeeting ? 'Atualize as informações do agendamento.' : 'Reserve um horário para alinhamento com o cliente.'"
      @close="showSlideOver = false"
    >
      <ReuniaoForm 
        ref="reuniaoFormRef"
        :initial-data="editingMeeting"
        @save="handleSave"
      />
      
      <template #footer>
        <button @click="showSlideOver = false" class="px-6 py-3 rounded-full font-bold text-[var(--color-text-secondary)] hover:bg-zinc-100 text-sm transition-all">Cancelar</button>
        <button 
          @click="reuniaoFormRef?.submit()"
          :disabled="isSaving"
          class="bg-[var(--color-primary)] hover:bg-[var(--color-primary-hover)] disabled:opacity-50 text-white px-8 py-3 rounded-full font-bold flex items-center gap-2 transition-all active:scale-95 shadow-lg shadow-blue-500/20 text-sm"
        >
          <Loader2 v-if="isSaving" class="w-4 h-4 animate-spin" />
          <Save v-else class="w-4 h-4" />
          {{ editingMeeting ? 'Salvar Alterações' : 'Agendar Reunião' }}
        </button>
      </template>
    </SlideOver>
  </div>
</template>
