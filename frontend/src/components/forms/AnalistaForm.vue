<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { Usuario } from '../../types'

const props = defineProps<{
  initialData?: Usuario
}>()

const emit = defineEmits(['save'])

const form = ref({
  nome: '',
  email: '',
  senha: '',
  cargo: 'Analista'
})

onMounted(() => {
  if (props.initialData) {
    form.value = {
      nome: props.initialData.nome,
      email: props.initialData.email,
      senha: '', // Senha não deve ser carregada
      cargo: props.initialData.cargo
    }
  }
})

const submit = () => {
  emit('save', { ...form.value })
}

defineExpose({ submit })
</script>

<template>
  <form @submit.prevent="submit" class="space-y-6">
    <div class="space-y-2">
      <label class="text-sm font-bold text-[var(--color-text-primary)] px-1">Nome Completo</label>
      <input 
        v-model="form.nome"
        type="text" 
        required
        placeholder="Ex: João Silva" 
        class="w-full bg-zinc-50 border border-zinc-200 rounded-2xl px-5 py-4 text-[15px] focus:ring-2 focus:ring-[var(--color-primary)]/20 focus:border-[var(--color-primary)] transition-all placeholder:text-zinc-400"
      />
    </div>

    <div class="space-y-2">
      <label class="text-sm font-bold text-[var(--color-text-primary)] px-1">E-mail Corporativo</label>
      <input 
        v-model="form.email"
        type="email" 
        required
        placeholder="Ex: joao@empresa.com" 
        class="w-full bg-zinc-50 border border-zinc-200 rounded-2xl px-5 py-4 text-[15px] focus:ring-2 focus:ring-[var(--color-primary)]/20 focus:border-[var(--color-primary)] transition-all placeholder:text-zinc-400"
      />
    </div>

    <div v-if="!initialData" class="space-y-2">
      <label class="text-sm font-bold text-[var(--color-text-primary)] px-1">Senha Inicial</label>
      <input 
        v-model="form.senha"
        type="password" 
        :required="!initialData"
        placeholder="Mínimo 6 caracteres" 
        class="w-full bg-zinc-50 border border-zinc-200 rounded-2xl px-5 py-4 text-[15px] focus:ring-2 focus:ring-[var(--color-primary)]/20 focus:border-[var(--color-primary)] transition-all placeholder:text-zinc-400"
      />
    </div>

    <div class="space-y-2">
      <label class="text-sm font-bold text-[var(--color-text-primary)] px-1">Cargo / Função</label>
      <select 
        v-model="form.cargo"
        required
        class="w-full bg-zinc-50 border border-zinc-200 rounded-2xl px-5 py-4 text-[15px] focus:ring-2 focus:ring-[var(--color-primary)]/20 focus:border-[var(--color-primary)] transition-all"
      >
        <option value="Analista">Analista de Implantação</option>
        <option value="Admin">Administrador</option>
      </select>
    </div>
  </form>
</template>
