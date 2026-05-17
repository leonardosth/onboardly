<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import DashboardLayout from './components/layout/DashboardLayout.vue'
import ToastContainer from './components/ui/ToastContainer.vue'

const route = useRoute()
const layout = computed(() => route.meta.layout || 'dashboard')
</script>

<template>
  <div class="antialiased font-sans">
    <DashboardLayout v-if="layout === 'dashboard'" />
    <router-view v-else />
    
    <ToastContainer />
  </div>
</template>

<style>
/* Reset and base transitions if needed */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s cubic-bezier(0.25, 0.1, 0.25, 1);
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
