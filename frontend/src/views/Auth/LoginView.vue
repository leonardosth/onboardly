<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { Loader2 } from 'lucide-vue-next';
import { useToastStore } from '../../stores/toastStore';
import { useAuthStore } from '../../stores/authStore';
import { authService } from '../../services/authService';

const router = useRouter();
const toastStore = useToastStore();
const authStore = useAuthStore();
const email = ref('');
const password = ref('');
const isLoading = ref(false);

const handleLogin = async () => {
  if (!email.value || !password.value) {
    toastStore.error('Preencha todos os campos.');
    return;
  }

  isLoading.value = true;
  try {
    const response = await authService.login({
      email: email.value,
      senha: password.value
    });

    authStore.setToken(response.token);
    authStore.setUser(response.user);

    toastStore.success(`Bem-vindo de volta, ${response.user.nome}!`);
    router.push('/');
  } catch (error: any) {
    console.error('Login error:', error);
    const message = error.response?.data?.error || 'Erro ao realizar login. Verifique suas credenciais.';
    toastStore.error(message);
  } finally {
    isLoading.value = false;
  }
};
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-[var(--color-background-app)] py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-12 bg-[var(--color-surface)] p-12 rounded-[var(--radius-apple)] shadow-premium border border-[var(--color-border-soft)]">
      <div class="text-center">
        <!-- Logo -->
        <div class="inline-flex mb-10">
          <div class="w-14 h-14 bg-[var(--color-primary)] rounded-2xl flex items-center justify-center shadow-lg shadow-blue-500/20">
            <div class="w-2 h-2 bg-white rounded-full"></div>
          </div>
        </div>
        <h2 class="text-3xl font-bold text-[var(--color-text-primary)] tracking-tight">
          Entrar no Onboardly
        </h2>
        <p class="mt-3 text-[15px] font-medium text-[var(--color-text-secondary)]">
          Gerencie a jornada de seus clientes com excelência.
        </p>
      </div>
      
      <form class="space-y-8" @submit.prevent="handleLogin">
        <div class="space-y-5">
          <div class="space-y-2">
            <label for="email-address" class="text-[14px] font-semibold text-[var(--color-text-primary)] px-1">E-mail</label>
            <input
              id="email-address"
              name="email"
              type="email"
              required
              v-model="email"
              class="w-full bg-zinc-100/50 border border-zinc-200/50 rounded-xl px-4 py-3.5 text-[var(--color-text-primary)] font-medium focus:bg-white focus:outline-none focus:ring-4 focus:ring-blue-500/5 focus:border-[var(--color-primary)] transition-all placeholder:text-[var(--color-text-tertiary)]"
              placeholder="seu@email.com"
            />
          </div>
          <div class="space-y-2">
            <div class="flex items-center justify-between px-1">
              <label for="password" class="text-[14px] font-semibold text-[var(--color-text-primary)]">Senha</label>
              <a href="#" class="text-[13px] font-bold text-[var(--color-primary)] hover:opacity-80 transition-opacity">Esqueceu a senha?</a>
            </div>
            <input
              id="password"
              name="password"
              type="password"
              required
              v-model="password"
              class="w-full bg-zinc-100/50 border border-zinc-200/50 rounded-xl px-4 py-3.5 text-[var(--color-text-primary)] font-medium focus:bg-white focus:outline-none focus:ring-4 focus:ring-blue-500/5 focus:border-[var(--color-primary)] transition-all placeholder:text-[var(--color-text-tertiary)]"
              placeholder="••••••••"
            />
          </div>
        </div>

        <div class="pt-4">
          <button
            type="submit"
            :disabled="isLoading"
            class="w-full flex justify-center py-4 px-4 border border-transparent text-sm font-bold rounded-full text-white bg-[var(--color-primary)] hover:bg-[var(--color-primary-hover)] focus:outline-none focus:ring-4 focus:ring-blue-500/10 transition-all active:scale-[0.98] shadow-lg shadow-blue-500/20 disabled:opacity-50"
          >
            <span v-if="isLoading" class="flex items-center gap-2">
              <Loader2 class="w-4 h-4 animate-spin" />
              Verificando...
            </span>
            <span v-else>Acessar Painel</span>
          </button>
        </div>
      </form>

      <div class="text-center pt-2">
        <p class="text-[13px] font-medium text-[var(--color-text-tertiary)]">
          Novo por aqui? 
          <a href="#" class="text-[var(--color-primary)] font-bold hover:opacity-80 transition-opacity ml-1">Solicite uma conta</a>
        </p>
      </div>
    </div>
  </div>
</template>
