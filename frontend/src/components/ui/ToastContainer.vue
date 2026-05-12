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

const getTypeClasses = (type: string) => {
  switch (type) {
    case 'success': return 'bg-emerald-50 text-emerald-800 border-emerald-100'
    case 'error': return 'bg-rose-50 text-rose-800 border-rose-100'
    case 'warning': return 'bg-amber-50 text-amber-800 border-amber-100'
    default: return 'bg-blue-50 text-blue-800 border-blue-100'
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
  <div class="fixed bottom-6 right-6 z-[100] flex flex-col gap-3 pointer-events-none">
    <TransitionGroup 
      name="toast"
      enter-active-class="transition duration-300 ease-out"
      enter-from-class="transform translate-y-4 opacity-0 scale-95"
      enter-to-class="transform translate-y-0 opacity-100 scale-100"
      leave-active-class="transition duration-200 ease-in"
      leave-from-class="transform opacity-100 scale-100"
      leave-to-class="transform opacity-0 scale-95"
    >
      <div 
        v-for="toast in toastStore.toasts" 
        :key="toast.id"
        class="pointer-events-auto flex items-center gap-3 px-4 py-3 rounded-xl border shadow-lg min-w-[320px] max-w-md"
        :class="getTypeClasses(toast.type)"
      >
        <component :is="getIcon(toast.type)" class="w-5 h-5 flex-shrink-0" :class="getIconClasses(toast.type)" />
        <p class="text-sm font-medium flex-1">{{ toast.message }}</p>
        <button 
          @click="toastStore.removeToast(toast.id)"
          class="p-1 hover:bg-black/5 rounded-lg transition-colors"
        >
          <X class="w-4 h-4 opacity-50 hover:opacity-100" />
        </button>
      </div>
    </TransitionGroup>
  </div>
</template>

<style scoped>
.toast-move {
  transition: all 0.3s ease;
}
</style>
