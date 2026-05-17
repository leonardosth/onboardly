<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useToastStore } from '../../stores/toastStore'
import { Plus, Search, ChevronLeft, ChevronRight, Loader2, Save, Mail, Shield, User } from 'lucide-vue-next'
import { usuarioService } from '../../services/usuarioService'
import type { Usuario } from '../../types'

import SlideOver from '../../components/ui/SlideOver.vue'
import AnalistaForm from '../../components/forms/AnalistaForm.vue'

const toastStore = useToastStore()
const analysts = ref<Usuario[]>([])
const isLoading = ref(true)

const showSlideOver = ref(false)
const selectedAnalyst = ref<Usuario | undefined>(undefined)
const isSaving = ref(false)
const analistaFormRef = ref<any>(null)

const searchQuery = ref('')

const fetchData = async () => {
  isLoading.value = true
  try {
    const data = await usuarioService.getAnalistas()
    analysts.value = data
  } catch (error) {
    console.error('Error fetching data:', error)
    toastStore.error('Erro ao carregar analistas.')
  } finally {
    isLoading.value = false
  }
}

const openNewAnalyst = () => {
  selectedAnalyst.value = undefined
  showSlideOver.value = true
}

const openEditAnalyst = (analyst: Usuario) => {
  selectedAnalyst.value = analyst
  showSlideOver.value = true
}

const handleSave = async (formData: any) => {
  isSaving.value = true
  try {
    if (selectedAnalyst.value) {
      await usuarioService.updateAnalista(selectedAnalyst.value.id, formData)
      toastStore.success('Analista atualizado com sucesso!')
    } else {
      await usuarioService.createAnalista(formData)
      toastStore.success('Analista cadastrado com sucesso!')
    }
    await fetchData()
    showSlideOver.value = false
  } catch (error: any) {
    console.error('Error saving analyst:', error)
    toastStore.error(error.message || 'Erro ao salvar analista.')
  } finally {
    isSaving.value = false
  }
}

const filteredAnalysts = computed(() => {
  return analysts.value.filter(item => {
    return item.nome.toLowerCase().includes(searchQuery.value.toLowerCase()) || 
           item.email.toLowerCase().includes(searchQuery.value.toLowerCase())
  })
})

onMounted(fetchData)
</script>

<template>
  <div class="space-y-10">
    <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-6">
      <div class="space-y-1">
        <h1 class="text-3xl font-bold tracking-tight text-[var(--color-text-primary)]">Gestão de Analistas</h1>
        <p class="text-[var(--color-text-secondary)] font-medium">Cadastre e gerencie os profissionais de implantação.</p>
      </div>
      <button 
        @click="openNewAnalyst"
        class="bg-[var(--color-primary)] text-white px-6 py-3 rounded-full font-bold flex items-center gap-2 hover:bg-[var(--color-primary-hover)] transition-all shadow-lg shadow-blue-500/20 active:scale-95 text-sm"
      >
        <Plus class="w-5 h-5" />
        Novo Analista
      </button>
    </div>

    <!-- Search Bar -->
    <div class="bg-[var(--color-surface)] p-2 rounded-2xl shadow-premium border border-[var(--color-border-soft)] flex flex-col md:flex-row items-center gap-2">
      <div class="relative flex-1 w-full md:w-auto">
        <Search class="w-4 h-4 absolute left-4 top-1/2 -translate-y-1/2 text-[var(--color-text-tertiary)]" />
        <input 
          v-model="searchQuery"
          type="text" 
          placeholder="Buscar por nome ou e-mail..." 
          class="w-full bg-transparent border-none focus:ring-0 pl-11 pr-4 py-3 text-sm text-[var(--color-text-primary)] placeholder:text-[var(--color-text-tertiary)]"
        />
      </div>
    </div>

    <!-- Grid Container -->
    <div v-if="isLoading" class="flex flex-col items-center justify-center py-20 bg-[var(--color-surface)] rounded-[var(--radius-apple)] border border-[var(--color-border-soft)] shadow-premium">
      <Loader2 class="w-10 h-10 text-[var(--color-primary)] animate-spin mb-3" />
      <span class="text-[var(--color-text-tertiary)] font-medium">Carregando time...</span>
    </div>

    <div v-else-if="filteredAnalysts.length === 0" class="flex flex-col items-center justify-center py-20 bg-[var(--color-surface)] rounded-[var(--radius-apple)] border border-[var(--color-border-soft)] shadow-premium">
      <User class="w-12 h-12 text-[var(--color-text-tertiary)] mb-3 opacity-20" />
      <span class="text-[var(--color-text-tertiary)] font-medium">Nenhum analista encontrado.</span>
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      <div 
        v-for="analyst in filteredAnalysts" 
        :key="analyst.id"
        @click="openEditAnalyst(analyst)"
        class="bg-[var(--color-surface)] p-6 rounded-[var(--radius-apple)] shadow-premium border border-[var(--color-border-soft)] hover:border-[var(--color-primary)]/30 transition-all cursor-pointer group relative overflow-hidden"
      >
        <div class="flex items-start justify-between mb-4">
          <div class="w-14 h-14 rounded-2xl bg-[var(--color-primary)]/10 flex items-center justify-center text-[var(--color-primary)] font-bold text-xl group-hover:scale-110 transition-transform">
            {{ analyst.nome.charAt(0).toUpperCase() }}
          </div>
          <span 
            class="px-3 py-1 rounded-full text-[10px] font-black uppercase tracking-widest border"
            :class="analyst.cargo === 'Admin' ? 'bg-purple-50 text-purple-600 border-purple-100' : 'bg-blue-50 text-blue-600 border-blue-100'"
          >
            {{ analyst.cargo }}
          </span>
        </div>

        <div class="space-y-1">
          <h3 class="font-bold text-lg text-[var(--color-text-primary)] group-hover:text-[var(--color-primary)] transition-colors line-clamp-1">
            {{ analyst.nome }}
          </h3>
          <div class="flex items-center gap-2 text-sm text-[var(--color-text-secondary)]">
            <Mail class="w-3.5 h-3.5 opacity-50" />
            <span class="line-clamp-1">{{ analyst.email }}</span>
          </div>
        </div>

        <div class="mt-6 pt-6 border-t border-[var(--color-border-divider)] flex items-center justify-between">
          <div class="flex items-center gap-1.5 text-[11px] font-bold text-[var(--color-text-tertiary)] uppercase tracking-wider">
            <Shield class="w-3.5 h-3.5" />
            Ativo
          </div>
          <button class="text-[var(--color-primary)] text-xs font-bold hover:underline opacity-0 group-hover:opacity-100 transition-opacity">
            Editar Perfil
          </button>
        </div>
      </div>
    </div>

    <!-- SlideOver for Creating/Editing -->
    <SlideOver 
      :show="showSlideOver" 
      :title="selectedAnalyst ? 'Editar Analista' : 'Novo Analista'"
      :description="selectedAnalyst ? 'Atualize os dados e o cargo do profissional.' : 'O analista poderá logar no sistema com as credenciais abaixo.'"
      @close="showSlideOver = false"
    >
      <AnalistaForm 
        ref="analistaFormRef"
        :initial-data="selectedAnalyst"
        @save="handleSave"
      />
      
      <template #footer>
        <button 
          @click="showSlideOver = false"
          class="px-6 py-3 rounded-full font-bold text-[var(--color-text-secondary)] hover:bg-zinc-100 transition-all text-sm"
        >
          Cancelar
        </button>
        <button 
          @click="analistaFormRef?.submit()"
          :disabled="isSaving"
          class="bg-[var(--color-primary)] hover:bg-[var(--color-primary-hover)] disabled:opacity-50 text-white px-8 py-3 rounded-full font-bold flex items-center gap-2 transition-all active:scale-95 shadow-lg shadow-blue-500/20 text-sm"
        >
          <Loader2 v-if="isSaving" class="w-4 h-4 animate-spin" />
          <Save v-else class="w-4 h-4" />
          {{ selectedAnalyst ? 'Salvar Alterações' : 'Cadastrar Analista' }}
        </button>
      </template>
    </SlideOver>
  </div>
</template>
