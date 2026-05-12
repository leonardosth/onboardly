<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import DashboardLayout from './components/layout/DashboardLayout.vue'
import ToastContainer from './components/ui/ToastContainer.vue'

const route = useRoute()
const layout = computed(() => route.meta.layout || 'dashboard')
</script>

<template>
  <div class="app-container">
    <DashboardLayout v-if="layout === 'dashboard'" />
    <router-view v-else />
    
    <ToastContainer />
  </div>
</template>

<style>
@reference "tailwindcss";

body {
  @apply antialiased text-slate-900 bg-slate-50;
}

/* Global scrollbar styling */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

::-webkit-scrollbar-track {
  @apply bg-transparent;
}

::-webkit-scrollbar-thumb {
  @apply bg-slate-200 rounded-full hover:bg-slate-300 transition-colors;
}
</style>
