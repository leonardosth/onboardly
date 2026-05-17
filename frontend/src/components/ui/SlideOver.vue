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
    <!-- Overlay -->
    <Transition
      enter-active-class="transition-opacity duration-500 ease-apple"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition-opacity duration-300 ease-apple"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div v-if="show" class="fixed inset-0 bg-zinc-950/20 backdrop-blur-[2px] z-50" @click="close"></div>
    </Transition>

    <!-- Slide Panel -->
    <Transition
      enter-active-class="transition-transform duration-600 ease-apple"
      enter-from-class="translate-x-full"
      enter-to-class="translate-x-0"
      leave-active-class="transition-transform duration-400 ease-apple"
      leave-from-class="translate-x-0"
      leave-to-class="translate-x-full"
    >
      <aside v-if="show" class="fixed top-0 right-0 h-full w-full max-w-lg bg-[var(--color-surface)] shadow-premium z-50 flex flex-col border-l border-[var(--color-border-soft)]">
        <!-- Header -->
        <header class="px-8 py-8 flex items-start justify-between">
          <div>
            <h2 class="text-2xl font-bold text-[var(--color-text-primary)] tracking-tight">{{ title }}</h2>
            <p v-if="description" class="text-[15px] font-medium text-[var(--color-text-secondary)] mt-1.5 leading-relaxed">{{ description }}</p>
          </div>
          <button @click="close" class="p-2.5 hover:bg-zinc-100 rounded-full transition-colors group">
            <X class="w-5 h-5 text-[var(--color-text-tertiary)] group-hover:text-[var(--color-text-primary)]" />
          </button>
        </header>

        <!-- Content -->
        <main class="flex-1 overflow-y-auto px-8 py-4 custom-scroll">
          <slot></slot>
        </main>

        <!-- Footer -->
        <footer class="px-8 py-8 border-t border-[var(--color-border-soft)] flex items-center justify-end gap-3">
          <slot name="footer"></slot>
        </footer>
      </aside>
    </Transition>
  </Teleport>
</template>

<style scoped>
.ease-apple {
  transition-timing-function: cubic-bezier(0.25, 0.1, 0.25, 1);
}

.custom-scroll::-webkit-scrollbar {
  width: 4px;
}
.custom-scroll::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scroll::-webkit-scrollbar-thumb {
  background: #D1D1D6;
  border-radius: 10px;
}
</style>
