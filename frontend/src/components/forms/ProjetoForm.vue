<script setup lang="ts">
import { ref, watch, onMounted } from 'vue'
import type { Projeto, Cliente, Analista } from '../../types'
import { clientService } from '../../services/clientService'
import { analistaService } from '../../services/analistaService'

const props = defineProps<{
  initialData?: Partial<Projeto>
}>()

const emit = defineEmits(['save'])

const form = ref({
  cliente_id: '',
  analista_id: '',
  data_contratacao: new Date().toISOString().split('T')[0],
  status_projeto: 'Backlog' as Projeto['status_projeto']
})

const clients = ref<Cliente[]>([])
const analistas = ref<Analista[]>([])

onMounted(async () => {
  const [clientsData, analistasData] = await Promise.all([
    clientService.getAll(),
    analistaService.getAll()
  ])
  clients.value = clientsData
  analistas.value = analistasData
})

watch(() => props.initialData, (newData) => {
  if (newData) {
    form.value = {
      cliente_id: newData.cliente_id || '',
      analista_id: newData.analista_id || '',
      data_contratacao: newData.data_contratacao ? newData.data_contratacao.split('T')[0] : new Date().toISOString().split('T')[0],
      status_projeto: newData.status_projeto || 'Backlog'
    }
  } else {
    form.value = {
      cliente_id: '',
      analista_id: '',
      data_contratacao: new Date().toISOString().split('T')[0],
      status_projeto: 'Backlog'
    }
  }
}, { immediate: true })

const submit = () => {
  const payload = { ...form.value }
  if (payload.data_contratacao) {
    // Converte YYYY-MM-DD para ISO string (RFC3339) para compatibilidade com Go time.Time
    payload.data_contratacao = new Date(`${payload.data_contratacao}T00:00:00Z`).toISOString()
  }
  emit('save', payload)
}

defineExpose({ submit })
</script>

<template>
  <form @submit.prevent="submit" class="space-y-8">
    <div class="space-y-2.5">
      <label class="text-[14px] font-semibold text-[var(--color-text-primary)] px-0.5">Cliente</label>
      <select
        v-model="form.cliente_id"
        required
        class="w-full px-4 py-3.5 rounded-[var(--radius-apple-sm)] bg-zinc-100/50 border border-zinc-200/50 focus:bg-white focus:outline-none focus:ring-4 focus:ring-[var(--color-primary)]/10 focus:border-[var(--color-primary)] transition-all text-[var(--color-text-primary)] font-medium appearance-none"
      >
        <option value="" disabled>Selecione um cliente</option>
        <option v-for="client in clients" :key="client.id" :value="client.id">{{ client.nome }}</option>
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
      <label class="text-[14px] font-semibold text-[var(--color-text-primary)] px-0.5">Data de Contratação</label>
      <input
        v-model="form.data_contratacao"
        type="date"
        required
        class="w-full px-4 py-3.5 rounded-[var(--radius-apple-sm)] bg-zinc-100/50 border border-zinc-200/50 focus:bg-white focus:outline-none focus:ring-4 focus:ring-[var(--color-primary)]/10 focus:border-[var(--color-primary)] transition-all font-mono text-[var(--color-text-primary)]"
      />
    </div>

    <div class="space-y-4">
      <label class="text-[14px] font-semibold text-[var(--color-text-primary)] px-0.5">Status Inicial</label>
      <div class="p-1 bg-zinc-100 rounded-xl flex gap-1">
        <button
          v-for="status in ['Backlog', 'Em_Andamento', 'Concluido']"
          :key="status"
          type="button"
          @click="form.status_projeto = status as any"
          :class="[
            'flex-1 px-3 py-2.5 rounded-lg text-xs font-bold transition-all duration-200',
            form.status_projeto === status 
              ? 'bg-white text-[var(--color-primary)] shadow-sm' 
              : 'text-[var(--color-text-tertiary)] hover:text-[var(--color-text-secondary)]'
          ]"
        >
          {{ status.replace('_', ' ') }}
        </button>
      </div>
    </div>
  </form>
</template>
