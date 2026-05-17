<script setup lang="ts">
import { useToastStore } from '../../stores/toastStore'
import { CheckCircle, AlertCircle, Info, AlertTriangle, X } from 'lucide-vue-next'

const toastStore = useToastStore()

const getIcon = (type: string) => {
  switch (type) {
    case 'success': return CheckCircle
    case 'error': return AlertCircle
    case 'warning': return AlertTriangle
    default: return Info
  }
}

const getIconClasses = (type: string) => {
  switch (type) {
    case 'success': return 'text-emerald-500'
    case 'error': return 'text-rose-500'
    case 'warning': return 'text-amber-500'
    default: return 'text-blue-500'
  }
}
</script>

<template>
  <div class="fixed top-6 right-6 z-[100] flex flex-col gap-3 pointer-events-none">
    <TransitionGroup 
      name="toast"
      enter-active-class="transition duration-500 ease-apple"
      enter-from-class="transform translate-y-[-20px] opacity-0 scale-95"
      enter-to-class="transform translate-y-0 opacity-100 scale-100"
      leave-active-class="transition duration-300 ease-apple"
      leave-from-class="transform opacity-100 scale-100"
      leave-to-class="transform opacity-0 scale-95"
    >
      <div 
        v-for="toast in toastStore.toasts" 
        :key="toast.id"
        class="pointer-events-auto flex items-center gap-4 px-5 py-4 rounded-[var(--radius-apple-sm)] bg-[var(--color-surface)] border border-[var(--color-border-soft)] shadow-premium min-w-[340px] max-w-md"
      >
        <div class="flex-shrink-0">
          <component :is="getIcon(toast.type)" class="w-5 h-5" :class="getIconClasses(toast.type)" />
        </div>
        <p class="text-[14px] font-semibold text-[var(--color-text-primary)] flex-1 leading-tight">{{ toast.message }}</p>
        <button 
          @click="toastStore.removeToast(toast.id)"
          class="p-1.5 hover:bg-zinc-100 rounded-full transition-colors group"
        >
          <X class="w-4 h-4 text-[var(--color-text-tertiary)] group-hover:text-[var(--color-text-primary)]" />
        </button>
      </div>
    </TransitionGroup>
  </div>
</template>

<style scoped>
.ease-apple {
  transition-timing-function: cubic-bezier(0.25, 0.1, 0.25, 1);
}

.toast-move {
  transition: all 0.4s cubic-bezier(0.25, 0.1, 0.25, 1);
}
</style>
