<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import type { Reuniao, Projeto, Cliente } from '../../types'
import { projectService } from '../../services/projectService'
import { clientService } from '../../services/clientService'

const props = defineProps<{
  initialData?: Partial<Reuniao>
}>()

const emit = defineEmits(['save'])

const form = ref({
  projeto_id: '',
  data_agendada: new Date().toISOString().substring(0, 16), // datetime-local format
  status: 'Agendada' as Reuniao['status'],
  observacoes: ''
})

const projects = ref<Projeto[]>([])
const clients = ref<Record<string, Cliente>>({})

onMounted(async () => {
  const [projectsData, clientsData] = await Promise.all([
    projectService.getAll(),
    clientService.getAll()
  ])
  projects.value = projectsData
  clients.value = clientsData.reduce((acc, c) => ({ ...acc, [c.id]: c }), {})
})

watch(() => props.initialData, (newData) => {
  if (newData) {
    form.value = {
      projeto_id: newData.projeto_id || '',
      data_agendada: newData.data_agendada ? newData.data_agendada.substring(0, 16) : new Date().toISOString().substring(0, 16),
      status: newData.status || 'Agendada',
      observacoes: newData.observacoes || ''
    }
  } else {
    form.value = {
      projeto_id: '',
      data_agendada: new Date().toISOString().substring(0, 16),
      status: 'Agendada',
      observacoes: ''
    }
  }
}, { immediate: true })

const submit = () => {
  emit('save', { ...form.value })
}

defineExpose({ submit })
</script>

<template>
  <form @submit.prevent="submit" class="space-y-6">
    <div class="space-y-2">
      <label class="text-sm font-bold text-slate-700 uppercase tracking-wider">Projeto / Cliente</label>
      <select
        v-model="form.projeto_id"
        required
        class="w-full px-4 py-3 rounded-xl border border-slate-200 focus:outline-none focus:ring-2 focus:ring-brand-blue/20 focus:border-brand-blue transition-all bg-white"
      >
        <option value="" disabled>Selecione um projeto</option>
        <option v-for="project in projects" :key="project.id" :value="project.id">
          {{ clients[project.cliente_id]?.nome || 'Projeto ' + project.id.substring(0,8) }}
        </option>
      </select>
    </div>

    <div class="space-y-2">
      <label class="text-sm font-bold text-slate-700 uppercase tracking-wider">Data e Hora</label>
      <input
        v-model="form.data_agendada"
        type="datetime-local"
        required
        class="w-full px-4 py-3 rounded-xl border border-slate-200 focus:outline-none focus:ring-2 focus:ring-brand-blue/20 focus:border-brand-blue transition-all font-mono"
      />
    </div>

    <div class="space-y-2">
      <label class="text-sm font-bold text-slate-700 uppercase tracking-wider">Status</label>
      <select
        v-model="form.status"
        required
        class="w-full px-4 py-3 rounded-xl border border-slate-200 focus:outline-none focus:ring-2 focus:ring-brand-blue/20 focus:border-brand-blue transition-all bg-white"
      >
        <option value="Agendada">Agendada</option>
        <option value="Realizada">Realizada</option>
        <option value="Remarcada">Remarcada</option>
        <option value="No_Show">No Show</option>
      </select>
    </div>

    <div class="space-y-2">
      <label class="text-sm font-bold text-slate-700 uppercase tracking-wider">Observações</label>
      <textarea
        v-model="form.observacoes"
        rows="4"
        placeholder="Notas sobre a reunião..."
        class="w-full px-4 py-3 rounded-xl border border-slate-200 focus:outline-none focus:ring-2 focus:ring-brand-blue/20 focus:border-brand-blue transition-all"
      ></textarea>
    </div>
  </form>
</template>
