<script setup lang="ts">
import { ref, onMounted, computed, watch } from 'vue'
import { useToastStore } from '../../stores/toastStore'
import { 
  Bold, Italic, Underline, Strikethrough, Code, Link as LinkIcon, 
  List, ListOrdered, Quote, Image as ImageIcon, Table as TableIcon,
  CheckCircle2, Clock, XCircle, Loader2, Plus, Save
} from 'lucide-vue-next'
import { reuniaoService } from '../../services/reuniaoService'
import { projectService } from '../../services/projectService'
import { clientService } from '../../services/clientService'
import type { Reuniao, Projeto, Cliente } from '../../types'

import SlideOver from '../../components/ui/SlideOver.vue'
import ReuniaoForm from '../../components/forms/ReuniaoForm.vue'

const toastStore = useToastStore()
const meetings = ref<Reuniao[]>([])
const projects = ref<Record<string, Projeto>>({})
const clients = ref<Record<string, Cliente>>({})
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
    const [meetingsData, projectsData, clientsData] = await Promise.all([
      reuniaoService.getAll(),
      projectService.getAll(),
      clientService.getAll()
    ])
    
    meetings.value = meetingsData
    
    projects.value = projectsData.reduce((acc, p) => {
      acc[p.id] = p
      return acc
    }, {} as Record<string, Projeto>)
    
    clients.value = clientsData.reduce((acc, c) => {
      acc[c.id] = c
      return acc
    }, {} as Record<string, Cliente>)

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
  if (!project) return 'Unknown Project'
  return clients.value[project.cliente_id]?.nome || 'Unknown Client'
}

const formatTime = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleTimeString('en-US', { hour: 'numeric', minute: '2-digit', hour12: true })
}

const formatDuration = (dateString: string) => {
  const date = new Date(dateString)
  const endDate = new Date(date.getTime() + 60 * 60 * 1000) // Default 1h duration
  return endDate.toLocaleTimeString('en-US', { hour: 'numeric', minute: '2-digit', hour12: true })
}

const getStatusLabel = (status: string) => {
  switch (status) {
    case 'Agendada': return 'Scheduled'
    case 'Realizada': return 'Completed'
    case 'No_Show': return 'No Show'
    default: return status
  }
}

onMounted(fetchData)
</script>

<template>
  <div class="space-y-8">
    <div class="flex justify-between items-center">
      <h1 class="text-[32px] font-bold text-slate-900">Meetings & Technical Log</h1>
      <button 
        @click="openNewMeeting"
        class="bg-brand-blue text-white px-6 py-3 rounded-xl font-bold flex items-center gap-2 hover:bg-blue-600 transition-all shadow-lg shadow-blue-200 active:scale-95"
      >
        <Plus class="w-5 h-5" />
        New Meeting
      </button>
    </div>

    <div v-if="isLoading" class="flex justify-center items-center h-64">
      <div class="flex flex-col items-center gap-4">
        <Loader2 class="w-10 h-10 text-brand-blue animate-spin" />
        <p class="text-slate-500 font-medium">Loading meetings...</p>
      </div>
    </div>

    <div v-else-if="meetings.length === 0" class="bg-white p-12 rounded-[24px] border border-slate-100 text-center">
      <Clock class="w-12 h-12 text-slate-300 mx-auto mb-4" />
      <h3 class="text-lg font-bold text-slate-900">No meetings found</h3>
      <p class="text-slate-500">Scheduled meetings will appear here.</p>
    </div>

    <div v-else class="grid grid-cols-1 lg:grid-cols-2 gap-8 items-start">
      <!-- Upcoming Meetings Column -->
      <div class="bg-white p-8 rounded-[24px] shadow-soft border border-slate-50">
        <h2 class="text-xl font-bold text-slate-900 mb-8">Meetings</h2>
        
        <div class="space-y-6 relative">
          <!-- Continuous Line -->
          <div class="absolute left-[105px] top-4 bottom-4 w-0.5 bg-slate-100"></div>

          <div 
            v-for="meeting in meetings" 
            :key="meeting.id"
            @click="selectedMeetingId = meeting.id"
            class="flex gap-8 group cursor-pointer"
          >
            <!-- Time -->
            <div class="w-20 pt-1 text-right">
              <p class="text-sm font-bold text-slate-900">{{ formatTime(meeting.data_agendada) }}</p>
              <p class="text-xs font-medium text-slate-400">{{ formatDuration(meeting.data_agendada) }}</p>
            </div>

            <!-- Indicator -->
            <div class="relative pt-2">
              <div 
                :class="[
                  selectedMeetingId === meeting.id ? 'ring-4 ring-blue-100' : '',
                  'w-3 h-3 rounded-full border-2 border-white ring-1 ring-slate-200 z-10 relative transition-all',
                  meeting.status === 'Realizada' ? 'bg-brand-emerald' : meeting.status === 'No_Show' ? 'bg-rose-500' : 'bg-brand-blue'
                ]"
              ></div>
            </div>

            <!-- Card -->
            <div 
              :class="[
                selectedMeetingId === meeting.id ? 'bg-blue-50/50 border-brand-blue/20 ring-1 ring-brand-blue/10 shadow-sm' : 'border-slate-100 hover:border-slate-200',
                'flex-1 p-4 rounded-2xl border transition-all flex items-center justify-between'
              ]"
            >
              <div>
                <p class="font-bold text-slate-900 truncate max-w-[200px]">{{ getClientNameByProjectId(meeting.projeto_id) }} - Meeting</p>
                <p class="text-xs font-medium text-slate-500 mt-0.5">Project: <span class="text-slate-700">{{ projects[meeting.projeto_id]?.id.substring(0,8) }}...</span></p>
              </div>

              <div class="flex flex-col items-end gap-2">
                <div 
                  :class="[
                    meeting.status === 'Realizada' ? 'bg-emerald-50 text-emerald-600 border-emerald-100' : meeting.status === 'No_Show' ? 'bg-rose-50 text-rose-600 border-rose-100' : 'bg-blue-50 text-blue-600 border-blue-100',
                    'flex items-center gap-1.5 px-3 py-1 rounded-full text-[11px] font-bold border'
                  ]"
                >
                  <CheckCircle2 v-if="meeting.status === 'Realizada'" class="w-3 h-3" />
                  <Clock v-else-if="meeting.status === 'Agendada'" class="w-3 h-3" />
                  <XCircle v-else class="w-3 h-3" />
                  {{ getStatusLabel(meeting.status) }}
                </div>
                <button 
                  @click.stop="openEditMeeting(meeting)"
                  class="text-[10px] font-bold text-slate-400 hover:text-brand-blue transition-colors uppercase tracking-wider"
                >
                  Edit
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Meeting Details Column -->
      <div v-if="selectedMeeting" class="bg-white p-8 rounded-[24px] shadow-soft border border-slate-50 space-y-8 sticky top-28">
        <div>
          <h2 class="text-xl font-bold text-slate-900 mb-6">Meeting Details</h2>
          
          <div class="space-y-4">
            <div class="flex items-center gap-2">
              <span class="text-sm font-bold text-slate-900">Client:</span>
              <span class="text-sm font-medium text-slate-500">{{ getClientNameByProjectId(selectedMeeting.projeto_id) }}</span>
            </div>
            <div class="flex items-center gap-2">
              <span class="text-sm font-bold text-slate-900">Date/Time:</span>
              <span class="text-sm font-medium text-slate-500">{{ new Date(selectedMeeting.data_agendada).toLocaleString() }}</span>
            </div>
            <div class="flex items-center gap-2">
              <span class="text-sm font-bold text-slate-900">Status:</span>
              <span :class="[selectedMeeting.status === 'Realizada' ? 'text-brand-emerald' : 'text-brand-blue', 'text-sm font-bold']">
                {{ getStatusLabel(selectedMeeting.status) }}
              </span>
            </div>
          </div>
        </div>

        <div class="border-t border-slate-100 pt-8 space-y-4">
          <h3 class="text-lg font-bold text-slate-900">Technical Notes</h3>
          
          <!-- Rich Text Toolbar Placeholder -->
          <div class="border border-slate-200 rounded-2xl overflow-hidden shadow-sm">
            <div class="bg-slate-50/50 p-3 border-b border-slate-200 flex flex-wrap gap-1.5">
              <button class="p-1.5 hover:bg-white rounded transition-colors text-slate-500"><Bold class="w-4 h-4" /></button>
              <button class="p-1.5 hover:bg-white rounded transition-colors text-slate-500"><Italic class="w-4 h-4" /></button>
              <button class="p-1.5 hover:bg-white rounded transition-colors text-slate-500"><Underline class="w-4 h-4" /></button>
              <button class="p-1.5 hover:bg-white rounded transition-colors text-slate-500"><Strikethrough class="w-4 h-4" /></button>
              <div class="w-px h-4 bg-slate-200 mx-1 mt-1.5"></div>
              <button class="p-1.5 hover:bg-white rounded transition-colors text-slate-500"><List class="w-4 h-4" /></button>
              <button class="p-1.5 hover:bg-white rounded transition-colors text-slate-500"><ListOrdered class="w-4 h-4" /></button>
              <div class="w-px h-4 bg-slate-200 mx-1 mt-1.5"></div>
              <button class="p-1.5 hover:bg-white rounded transition-colors text-slate-500"><LinkIcon class="w-4 h-4" /></button>
              <button class="p-1.5 hover:bg-white rounded transition-colors text-slate-500"><ImageIcon class="w-4 h-4" /></button>
              <button class="p-1.5 hover:bg-white rounded transition-colors text-slate-500"><TableIcon class="w-4 h-4" /></button>
              <button class="p-1.5 hover:bg-white rounded transition-colors text-slate-500 ml-auto"><Code class="w-4 h-4" /></button>
            </div>
            <textarea 
              v-model="selectedMeeting.observacoes"
              class="w-full p-6 min-h-[250px] focus:outline-none text-slate-700 leading-relaxed placeholder:text-slate-300 font-medium"
              placeholder="Start typing your technical notes here..."
            ></textarea>
          </div>
        </div>
      </div>
    </div>

    <!-- SlideOver for Creating/Editing -->
    <SlideOver 
      :show="showSlideOver" 
      :title="editingMeeting ? 'Editar Reunião' : 'Nova Reunião'"
      :description="editingMeeting ? 'Atualize os detalhes do agendamento.' : 'Agende uma nova reunião técnica com o cliente.'"
      @close="showSlideOver = false"
    >
      <ReuniaoForm 
        ref="reuniaoFormRef"
        :initial-data="editingMeeting"
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
          @click="reuniaoFormRef?.submit()"
          :disabled="isSaving"
          class="bg-brand-blue hover:bg-brand-blue/90 disabled:opacity-50 text-white px-8 py-2.5 rounded-xl font-bold flex items-center gap-2 transition-all active:scale-95 shadow-lg shadow-brand-blue/20"
        >
          <Loader2 v-if="isSaving" class="w-4 h-4 animate-spin" />
          <Save v-else class="w-4 h-4" />
          {{ editingMeeting ? 'Salvar Alterações' : 'Agendar Reunião' }}
        </button>
      </template>
    </SlideOver>
  </div>
</template>
