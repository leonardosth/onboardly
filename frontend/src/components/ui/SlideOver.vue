<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'
import { X } from 'lucide-vue-next'

interface Props {
  show: boolean
  title: string
  description?: string
}

const props = defineProps<Props>()
const emit = defineEmits(['close'])

const close = () => {
  emit('close')
}

const handleEsc = (e: KeyboardEvent) => {
  if (e.key === 'Escape' && props.show) {
    close()
  }
}

onMounted(() => {
  window.addEventListener('keydown', handleEsc)
})

onUnmounted(() => {
  window.removeEventListener('keydown', handleEsc)
})
</script>

<template>
  <Teleport to="body">
    <Transition
      enter-active-class="transition-opacity duration-300 ease-out"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition-opacity duration-200 ease-in"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div v-if="show" class="fixed inset-0 bg-slate-900/40 backdrop-blur-sm z-40" @click="close"></div>
    </Transition>

    <Transition
      enter-active-class="transition-transform duration-500 ease-out"
      enter-from-class="translate-x-full"
      enter-to-class="translate-x-0"
      leave-active-class="transition-transform duration-300 ease-in"
      leave-from-class="translate-x-0"
      leave-to-class="translate-x-full"
    >
      <aside v-if="show" class="fixed top-0 right-0 h-full w-full max-w-md bg-white shadow-2xl z-50 flex flex-col border-l border-slate-100">
        <!-- Header -->
        <header class="px-6 py-6 border-b border-slate-50 flex items-center justify-between">
          <div>
            <h2 class="text-xl font-bold text-slate-900 tracking-tight">{{ title }}</h2>
            <p v-if="description" class="text-sm font-medium text-slate-500 mt-1">{{ description }}</p>
          </div>
          <button @click="close" class="p-2 hover:bg-slate-50 rounded-xl transition-colors group">
            <X class="w-5 h-5 text-slate-400 group-hover:text-slate-600" />
          </button>
        </header>

        <!-- Content -->
        <main class="flex-1 overflow-y-auto px-6 py-8">
          <slot></slot>
        </main>

        <!-- Footer -->
        <footer class="px-6 py-6 border-t border-slate-50 bg-slate-50/50 flex items-center justify-end gap-3">
          <slot name="footer"></slot>
        </footer>
      </aside>
    </Transition>
  </Teleport>
</template>

<style scoped>
/* Custom scrollbar for better aesthetic */
main::-webkit-scrollbar {
  width: 4px;
}
main::-webkit-scrollbar-track {
  background: transparent;
}
main::-webkit-scrollbar-thumb {
  background: #E2E8F0;
  border-radius: 10px;
}
main::-webkit-scrollbar-thumb:hover {
  background: #CBD5E1;
}
</style>
