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
  emit('save', { ...form.value })
}

defineExpose({ submit })
</script>

<template>
  <form @submit.prevent="submit" class="space-y-6">
    <div class="space-y-2">
      <label class="text-sm font-bold text-slate-700 uppercase tracking-wider">Cliente</label>
      <select
        v-model="form.cliente_id"
        required
        class="w-full px-4 py-3 rounded-xl border border-slate-200 focus:outline-none focus:ring-2 focus:ring-brand-blue/20 focus:border-brand-blue transition-all bg-white"
      >
        <option value="" disabled>Selecione um cliente</option>
        <option v-for="client in clients" :key="client.id" :value="client.id">{{ client.nome }}</option>
      </select>
    </div>

    <div class="space-y-2">
      <label class="text-sm font-bold text-slate-700 uppercase tracking-wider">Analista Responsável</label>
      <select
        v-model="form.analista_id"
        required
        class="w-full px-4 py-3 rounded-xl border border-slate-200 focus:outline-none focus:ring-2 focus:ring-brand-blue/20 focus:border-brand-blue transition-all bg-white"
      >
        <option value="" disabled>Selecione um analista</option>
        <option v-for="analista in analistas" :key="analista.id" :value="analista.id">{{ analista.nome }}</option>
      </select>
    </div>

    <div class="space-y-2">
      <label class="text-sm font-bold text-slate-700 uppercase tracking-wider">Data de Contratação</label>
      <input
        v-model="form.data_contratacao"
        type="date"
        required
        class="w-full px-4 py-3 rounded-xl border border-slate-200 focus:outline-none focus:ring-2 focus:ring-brand-blue/20 focus:border-brand-blue transition-all font-mono"
      />
    </div>

    <div class="space-y-2">
      <label class="text-sm font-bold text-slate-700 uppercase tracking-wider">Status Inicial</label>
      <div class="grid grid-cols-3 gap-3">
        <button
          v-for="status in ['Backlog', 'Em_Andamento', 'Concluido']"
          :key="status"
          type="button"
          @click="form.status_projeto = status as any"
          :class="[
            'px-3 py-2 rounded-lg border text-[10px] font-bold transition-all uppercase tracking-tighter',
            form.status_projeto === status 
              ? 'border-brand-blue bg-brand-blue/5 text-brand-blue shadow-sm' 
              : 'border-slate-100 text-slate-400 hover:bg-slate-50'
          ]"
        >
          {{ status.replace('_', ' ') }}
        </button>
      </div>
    </div>
  </form>
</template>
