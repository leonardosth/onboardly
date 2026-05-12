<script setup lang="ts">
import { ref, watch } from 'vue'
import type { Cliente } from '../../types'

const props = defineProps<{
  initialData?: Partial<Cliente>
}>()

const emit = defineEmits(['save'])

const form = ref({
  nome: '',
  cnpj: ''
})

watch(() => props.initialData, (newData) => {
  if (newData) {
    form.value = {
      nome: newData.nome || '',
      cnpj: newData.cnpj || ''
    }
  } else {
    form.value = { nome: '', cnpj: '' }
  }
}, { immediate: true })

const submit = () => {
  emit('save', { ...form.value })
}

defineExpose({ submit })
</script>

<template>
  <form @submit.prevent="submit" class="space-y-6" id="cliente-form">
    <div class="space-y-2">
      <label for="nome" class="text-sm font-bold text-slate-700 uppercase tracking-wider">Razão Social / Nome</label>
      <input
        id="nome"
        v-model="form.nome"
        type="text"
        required
        placeholder="Ex: Acme Corp"
        class="w-full px-4 py-3 rounded-xl border border-slate-200 focus:outline-none focus:ring-2 focus:ring-brand-blue/20 focus:border-brand-blue transition-all"
      />
    </div>

    <div class="space-y-2">
      <label for="cnpj" class="text-sm font-bold text-slate-700 uppercase tracking-wider">CNPJ</label>
      <input
        id="cnpj"
        v-model="form.cnpj"
        type="text"
        required
        placeholder="00.000.000/0000-00"
        class="w-full px-4 py-3 rounded-xl border border-slate-200 focus:outline-none focus:ring-2 focus:ring-brand-blue/20 focus:border-brand-blue transition-all font-mono"
      />
    </div>
  </form>
</template>
