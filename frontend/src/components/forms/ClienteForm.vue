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
  <form @submit.prevent="submit" class="space-y-8" id="cliente-form">
    <div class="space-y-2.5">
      <label for="nome" class="text-[14px] font-semibold text-[var(--color-text-primary)] px-0.5">Razão Social / Nome</label>
      <input
        id="nome"
        v-model="form.nome"
        type="text"
        required
        placeholder="Digite o nome da empresa"
        class="w-full px-4 py-3.5 rounded-[var(--radius-apple-sm)] bg-zinc-100/50 border border-zinc-200/50 focus:bg-white focus:outline-none focus:ring-4 focus:ring-[var(--color-primary)]/10 focus:border-[var(--color-primary)] transition-all placeholder:text-[var(--color-text-tertiary)] text-[var(--color-text-primary)] font-medium"
      />
    </div>

    <div class="space-y-2.5">
      <label for="cnpj" class="text-[14px] font-semibold text-[var(--color-text-primary)] px-0.5">CNPJ</label>
      <input
        id="cnpj"
        v-model="form.cnpj"
        type="text"
        required
        placeholder="00.000.000/0000-00"
        class="w-full px-4 py-3.5 rounded-[var(--radius-apple-sm)] bg-zinc-100/50 border border-zinc-200/50 focus:bg-white focus:outline-none focus:ring-4 focus:ring-[var(--color-primary)]/10 focus:border-[var(--color-primary)] transition-all font-mono placeholder:text-[var(--color-text-tertiary)] text-[var(--color-text-primary)]"
      />
    </div>
  </form>
</template>
