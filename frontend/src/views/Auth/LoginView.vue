<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { Loader2 } from 'lucide-vue-next';
import { useToastStore } from '../../stores/toastStore';

const router = useRouter();
const toastStore = useToastStore();
const email = ref('');
const password = ref('');
const isLoading = ref(false);

const handleLogin = async () => {
  isLoading.value = true;
  // Simulação de login conforme solicitado
  setTimeout(() => {
    localStorage.setItem('isAuthenticated', 'true');
    toastStore.success('Bem-vindo ao Onboardly!');
    router.push('/');
    isLoading.value = false;
  }, 800);
};
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-[#F8FAFC] py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-10 bg-white p-12 rounded-[32px] shadow-soft border border-slate-50">
      <div class="text-center">
        <!-- Logo -->
        <div class="inline-flex mb-8">
          <div class="w-12 h-12 border-4 border-brand-cyan rounded-full flex items-center justify-center">
            <div class="w-3 h-3 bg-brand-cyan rounded-full"></div>
          </div>
        </div>
        <h2 class="text-3xl font-bold text-slate-900 tracking-tight">
          Bem-vindo de volta
        </h2>
        <p class="mt-3 text-sm font-medium text-slate-500">
          Acesse sua conta para gerenciar implantações
        </p>
      </div>
      
      <form class="space-y-6" @submit.prevent="handleLogin">
        <div class="space-y-4">
          <div class="space-y-1.5">
            <label for="email-address" class="text-sm font-bold text-slate-900 ml-1">E-mail</label>
            <input
              id="email-address"
              name="email"
              type="email"
              required
              v-model="email"
              class="w-full bg-slate-50 border border-slate-100 rounded-xl px-4 py-3 text-slate-600 font-medium focus:outline-none focus:ring-2 focus:ring-brand-blue/20 focus:bg-white transition-all placeholder:text-slate-300"
              placeholder="seu@email.com"
            />
          </div>
          <div class="space-y-1.5">
            <div class="flex items-center justify-between ml-1">
              <label for="password" class="text-sm font-bold text-slate-900">Senha</label>
              <a href="#" class="text-xs font-bold text-brand-blue hover:underline">Esqueceu a senha?</a>
            </div>
            <input
              id="password"
              name="password"
              type="password"
              required
              v-model="password"
              class="w-full bg-slate-50 border border-slate-100 rounded-xl px-4 py-3 text-slate-600 font-medium focus:outline-none focus:ring-2 focus:ring-brand-blue/20 focus:bg-white transition-all placeholder:text-slate-300"
              placeholder="••••••••"
            />
          </div>
        </div>

        <div class="pt-2">
          <button
            type="submit"
            :disabled="isLoading"
            class="w-full flex justify-center py-3.5 px-4 border border-transparent text-sm font-bold rounded-xl text-white bg-brand-blue hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-blue transition-all active:scale-[0.98] shadow-lg shadow-blue-200 disabled:opacity-50"
          >
            <span v-if="isLoading" class="flex items-center gap-2">
              <Loader2 class="w-4 h-4 animate-spin" />
              Entrando...
            </span>
            <span v-else>Acessar Painel</span>
          </button>
        </div>
      </form>

      <div class="text-center pt-4">
        <p class="text-xs font-medium text-slate-400">
          Não tem uma conta? 
          <a href="#" class="text-brand-blue font-bold hover:underline">Solicite acesso</a>
        </p>
      </div>
    </div>
  </div>
</template>


