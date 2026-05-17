<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import type { Reuniao, Projeto, Cliente, Analista } from '../../types'
import { projectService } from '../../services/projectService'
import { clientService } from '../../services/clientService'
import { analistaService } from '../../services/analistaService'
import { useAuthStore } from '../../stores/authStore'

const props = defineProps<{
  initialData?: Partial<Reuniao>
}>()

const emit = defineEmits(['save'])
const authStore = useAuthStore()

const form = ref({
  projeto_id: '',
  analista_id: authStore.user?.id || '',
  data_agendada: new Date().toISOString().substring(0, 16), // datetime-local format
  status: 'Agendada' as Reuniao['status'],
  observacoes: ''
})

const projects = ref<Projeto[]>([])
const clients = ref<Record<string, Cliente>>({})
const analistas = ref<Analista[]>([])

onMounted(async () => {
  const [projectsData, clientsData, analistasData] = await Promise.all([
    projectService.getAll(),
    clientService.getAll(),
    analistaService.getAll()
  ])
  projects.value = projectsData
  clients.value = clientsData.reduce((acc, c) => ({ ...acc, [c.id]: c }), {})
  analistas.value = analistasData
})

watch(() => props.initialData, (newData) => {
  if (newData) {
    form.value = {
      projeto_id: newData.projeto_id || '',
      analista_id: newData.analista_id || authStore.user?.id || '',
      data_agendada: newData.data_agendada ? newData.data_agendada.substring(0, 16) : new Date().toISOString().substring(0, 16),
      status: newData.status || 'Agendada',
      observacoes: newData.observacoes || ''
    }
  } else {
    form.value = {
      projeto_id: '',
      analista_id: authStore.user?.id || '',
      data_agendada: new Date().toISOString().substring(0, 16),
      status: 'Agendada',
      observacoes: ''
    }
  }
}, { immediate: true })

const submit = () => {
  const payload = { ...form.value }
  if (payload.data_agendada) {
    // Converte YYYY-MM-DDTHH:mm para ISO string (RFC3339) para compatibilidade com Go time.Time
    payload.data_agendada = new Date(payload.data_agendada).toISOString()
  }
  emit('save', payload)
}

defineExpose({ submit })
</script>

<template>
  <form @submit.prevent="submit" class="space-y-8">
    <div class="space-y-2.5">
      <label class="text-[14px] font-semibold text-[var(--color-text-primary)] px-0.5">Projeto / Cliente</label>
      <select
        v-model="form.projeto_id"
        required
        class="w-full px-4 py-3.5 rounded-[var(--radius-apple-sm)] bg-zinc-100/50 border border-zinc-200/50 focus:bg-white focus:outline-none focus:ring-4 focus:ring-[var(--color-primary)]/10 focus:border-[var(--color-primary)] transition-all text-[var(--color-text-primary)] font-medium appearance-none"
      >
        <option value="" disabled>Selecione um projeto</option>
        <option v-for="project in projects" :key="project.id" :value="project.id">
          {{ clients[project.cliente_id]?.nome || 'Projeto ' + project.id.substring(0,8) }}
        </option>
      </select>
    </div>

    <div class="space-y-2.5">
      <label class="text-[14px] font-semibold text-[var(--color-text-primary)] px-0.5">Analista Responsável</label>
      <select
        v-model="form.analista_id"
        required
        class="w-full px-4 py-3.5 rounded-[var(--radius-apple-sm)] bg-zinc-100/50 border border-zinc-200/50 focus:bg-white focus:outline-none focus:ring-4 focus:ring-[var(--color-primary)]/10 focus:border-[var(--color-primary)] transition-all text-[var(--color-text-primary)] font-medium appearance-none"
      >
        <option value="" disabled>Selecione um analista</option>
        <option v-for="analista in analistas" :key="analista.id" :value="analista.id">{{ analista.nome }}</option>
      </select>
    </div>

    <div class="space-y-2.5">
      <label class="text-[14px] font-semibold text-[var(--color-text-primary)] px-0.5">Data e Hora</label>
      <input
        v-model="form.data_agendada"
        type="datetime-local"
        required
        class="w-full px-4 py-3.5 rounded-[var(--radius-apple-sm)] bg-zinc-100/50 border border-zinc-200/50 focus:bg-white focus:outline-none focus:ring-4 focus:ring-[var(--color-primary)]/10 focus:border-[var(--color-primary)] transition-all font-mono text-[var(--color-text-primary)]"
      />
    </div>

    <div class="space-y-2.5">
      <label class="text-[14px] font-semibold text-[var(--color-text-primary)] px-0.5">Status</label>
      <select
        v-model="form.status"
        required
        class="w-full px-4 py-3.5 rounded-[var(--radius-apple-sm)] bg-zinc-100/50 border border-zinc-200/50 focus:bg-white focus:outline-none focus:ring-4 focus:ring-[var(--color-primary)]/10 focus:border-[var(--color-primary)] transition-all text-[var(--color-text-primary)] font-medium appearance-none"
      >
        <option value="Agendada">Agendada</option>
        <option value="Realizada">Realizada</option>
        <option value="Remarcada">Remarcada</option>
        <option value="No_Show">No Show</option>
      </select>
    </div>

    <div class="space-y-2.5">
      <label class="text-[14px] font-semibold text-[var(--color-text-primary)] px-0.5">Observações</label>
      <textarea
        v-model="form.observacoes"
        rows="4"
        placeholder="Notas sobre a reunião..."
        class="w-full px-4 py-3.5 rounded-[var(--radius-apple-sm)] bg-zinc-100/50 border border-zinc-200/50 focus:bg-white focus:outline-none focus:ring-4 focus:ring-[var(--color-primary)]/10 focus:border-[var(--color-primary)] transition-all text-[var(--color-text-primary)] font-medium placeholder:text-[var(--color-text-tertiary)] resize-none"
      ></textarea>
    </div>
  </form>
</template>
